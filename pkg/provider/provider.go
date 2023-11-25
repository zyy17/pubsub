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

package provider

import (
	"context"

	pb "github.com/zyy17/pubsub/pkg/proto"
)

// Provider is the interface for the storage provider of pubsub.
type Provider interface {
	// StoreMessage stores the message and offline message list in the provider.
	StoreMessage(ctx context.Context, message *pb.Message, subscribers []string) error

	// GetSubscribers returns the subscribers for the topic.
	GetSubscribers(ctx context.Context, topic string) ([]string, error)

	// Subscribe stores the subscription in the provider and return the offline messages.
	Subscribe(ctx context.Context, subscriberId string, topics []string) ([]*pb.Message, error)

	// Unsubscribe removes the subscription from the provider.
	Unsubscribe(ctx context.Context, subscriberId string, topics []string) error

	// RemoveOfflineMessage removes the offline message from the provider.
	// When every subscriber has received the message, the message will be removed.
	RemoveOfflineMessage(ctx context.Context, subscriberId, topic, messageId string) error
}

type StoreType string

const (
	// StoreTypeEtcd uses etcd as the provider.
	StoreTypeEtcd = "etcd"
)

type Options struct {
	StorageEndpoints []string
	StoreType        StoreType
}
