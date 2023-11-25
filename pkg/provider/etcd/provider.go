// Copyright 2023 zyy17 <zyylsxm@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package etcd

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/protobuf/encoding/protojson"

	pb "github.com/zyy17/pubsub/pkg/proto"
	"github.com/zyy17/pubsub/pkg/provider"
)

const (
	DefaultTimeout = 5 * time.Second

	// Placeholder is the empty string(etcd supports empty value) and used for the keys that don't need the value.
	Placeholder = ""
)

// NewEtcdProvider creates the provider that uses etcd backend storage.
func NewEtcdProvider(opts *provider.Options) (provider.Provider, error) {
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   opts.StorageEndpoints,
		DialTimeout: DefaultTimeout,
	})
	if err != nil {
		return nil, err
	}

	return &etcdProvider{
		etcdClient: etcdClient,
	}, nil
}

type etcdProvider struct {
	etcdClient *clientv3.Client
}

var _ provider.Provider = (*etcdProvider)(nil)

// StoreMessage will store the message in etcd.
// It creates the following key-value pairs:
//   - Message body: /messages/${message_id} -> ${json_message_string}.
//   - Message subscriber list: /messages/${message_id}/${subscriber_id} -> ${placeholder}.
//   - Offline message list: /offline/${subscriber_id}/${topic}/${message_id} -> ${placeholder}.
func (p *etcdProvider) StoreMessage(ctx context.Context, message *pb.Message, subscribers []string) error {
	data, err := p.marshal(message)
	if err != nil {
		return err
	}

	// Create a etcd transaction.
	txn := p.etcdClient.Txn(ctx)

	var ops []clientv3.Op

	// Store the offline message list for each subscriber.
	for _, subscriberId := range subscribers {
		offlineMessageKey := fmt.Sprintf("/offline/%s/%s/%s", subscriberId, message.Topic, message.MessageId)
		subscriberKey := fmt.Sprintf("/messages/%s/%s", message.MessageId, subscriberId)
		ops = append(ops, clientv3.OpPut(offlineMessageKey, Placeholder), clientv3.OpPut(subscriberKey, Placeholder))
	}

	// Store the message body.
	messageKey := fmt.Sprintf("/messages/%s", message.MessageId)
	ops = append(ops, clientv3.OpPut(messageKey, data))

	// Commit the transaction.
	txnResp, err := txn.Then(ops...).Commit()
	if err != nil {
		return err
	}

	if !txnResp.Succeeded {
		return fmt.Errorf("failed to commit transaction")
	}

	return nil
}

// GetSubscribers will get the subscribers of the topic by using the key prefix '/topics/${topic}/$'.
func (p *etcdProvider) GetSubscribers(ctx context.Context, topic string) ([]string, error) {
	subscriberPrefixKey := fmt.Sprintf("/topics/%s", topic)
	resp, err := p.etcdClient.Get(ctx, subscriberPrefixKey, clientv3.WithPrefix(), clientv3.WithKeysOnly())
	if err != nil {
		return nil, err
	}

	var subscribers []string
	for _, kv := range resp.Kvs {
		subscriberId := string(kv.Key[len(subscriberPrefixKey)+1:])
		subscribers = append(subscribers, subscriberId)
	}

	return subscribers, nil
}

// Subscribe will store the subscription in etcd.
// It creates the following key-value pairs:
//   - Topic -> Subscribers: /topics/${topic}/${subscriber_id} -> ${placeholder}.
//   - Subscriber -> Topics: /subscribers/${subscriber_id}/${topic} -> ${placeholder}.
func (p *etcdProvider) Subscribe(ctx context.Context, subscriberId string, topics []string) ([]*pb.Message, error) {
	// Create a etcd transaction.
	txn := p.etcdClient.Txn(ctx)

	// Store the subscription for each topic.
	var ops []clientv3.Op
	for _, topic := range topics {
		// Store the subscription for the topic.
		topicKey := fmt.Sprintf("/topics/%s/%s", topic, subscriberId)
		ops = append(ops, clientv3.OpPut(topicKey, Placeholder))

		// Store the subscription for the subscriber.
		subscriberKey := fmt.Sprintf("/subscribers/%s/%s", subscriberId, topic)
		ops = append(ops, clientv3.OpPut(subscriberKey, Placeholder))
	}

	// Commit the transaction.
	txnResp, err := txn.Then(ops...).Commit()
	if err != nil {
		return nil, err
	}

	if !txnResp.Succeeded {
		return nil, fmt.Errorf("failed to commit transaction")
	}

	// Get the offline messages for the subscriber.
	var offlineMessages []*pb.Message
	for _, topic := range topics {
		offlineMessageKeyPrefix := fmt.Sprintf("/offline/%s/%s", subscriberId, topic)
		resp, err := p.etcdClient.Get(ctx, offlineMessageKeyPrefix, clientv3.WithPrefix(), clientv3.WithKeysOnly())
		if err != nil {
			return nil, err
		}

		for _, kv := range resp.Kvs {
			messageId := string(kv.Key[len(offlineMessageKeyPrefix)+1:])
			messageKey := fmt.Sprintf("/messages/%s", messageId)
			resp, err := p.etcdClient.Get(ctx, messageKey)
			if err != nil {
				return nil, err
			}

			var message pb.Message
			if err = protojson.Unmarshal(resp.Kvs[0].Value, &message); err != nil {
				return nil, err
			}
			offlineMessages = append(offlineMessages, &message)
		}
	}

	return offlineMessages, nil
}

// Unsubscribe will remove the subscription from etcd.
func (p *etcdProvider) Unsubscribe(ctx context.Context, subscriberId string, topics []string) error {
	// Create a etcd transaction.
	txn := p.etcdClient.Txn(ctx)

	// Remove the subscription for each topic.
	var ops []clientv3.Op
	for _, topic := range topics {
		// Remove the subscription for the topic.
		topicKey := fmt.Sprintf("/topics/%s/%s", topic, subscriberId)
		ops = append(ops, clientv3.OpDelete(topicKey))

		// Remove the subscription for the subscriber.
		subscriberKey := fmt.Sprintf("/subscribers/%s/%s", subscriberId, topic)
		ops = append(ops, clientv3.OpDelete(subscriberKey))
	}

	// Commit the transaction.
	txnResp, err := txn.Then(ops...).Commit()
	if err != nil {
		return err
	}

	if !txnResp.Succeeded {
		return fmt.Errorf("failed to commit transaction")
	}

	return nil
}

// RemoveOfflineMessage removes the offline message from etcd. It happens when publisher receives the puback.
// When every subscriber has received the message, the message will be removed.
func (p *etcdProvider) RemoveOfflineMessage(ctx context.Context, subscriberId, topic, messageId string) error {
	// Create a etcd transaction.
	txn := p.etcdClient.Txn(ctx)

	// Remove the offline message for the subscriber.
	txn = txn.Then(
		clientv3.OpDelete(fmt.Sprintf("/offline/%s/%s/%s", subscriberId, topic, messageId)),
		clientv3.OpDelete(fmt.Sprintf("/messages/%s/%s", messageId, subscriberId)),
	)

	// Commit the transaction.
	txnResp, err := txn.Commit()
	if err != nil {
		return err
	}

	if !txnResp.Succeeded {
		return fmt.Errorf("failed to commit transaction")
	}

	// Get the message subscriber list.
	messageSubscriberKeyPrefix := fmt.Sprintf("/messages/%s", messageId)
	resp, err := p.etcdClient.Get(ctx, messageSubscriberKeyPrefix, clientv3.WithPrefix(), clientv3.WithKeysOnly())
	if err != nil {
		return err
	}
	// That means all subscribers have received the message and the rest key is the message body.
	if len(resp.Kvs) == 1 {
		// Remove the message body.
		_, err := p.etcdClient.Delete(ctx, fmt.Sprintf("/messages/%s", messageId))
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *etcdProvider) marshal(msg *pb.Message) (string, error) {
	marshaler := protojson.MarshalOptions{
		EmitUnpopulated: true,
	}
	jsonStr, err := marshaler.Marshal(msg)
	if err != nil {
		return "", err
	}

	return string(jsonStr), nil
}
