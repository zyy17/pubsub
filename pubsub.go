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

package pubsub

import (
	"context"
	"io"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/zyy17/pubsub/pkg/proto"
	"github.com/zyy17/pubsub/pkg/provider"
	"github.com/zyy17/pubsub/pkg/provider/factory"
)

var (
	// ErrSubscribeFailed is the error returned when subscribe failed.
	ErrSubscribeFailed = status.Errorf(codes.Internal, "subscribe failed")

	// ErrGetSubscribersFailed is the error returned when get subscribers failed.
	ErrGetSubscribersFailed = status.Errorf(codes.Internal, "get subscribers failed")

	// ErrSubscribersNotFound is the error returned when subscribers not found.
	ErrSubscribersNotFound = status.Errorf(codes.Internal, "subscribers not found")

	// ErrUnsubscribeFailed is the error returned when unsubscribe failed.
	ErrUnsubscribeFailed = status.Errorf(codes.Internal, "unsubscribe failed")

	// ErrPublishAckFailed is the error returned when puback failed.
	ErrPublishAckFailed = status.Errorf(codes.Internal, "publish ack failed")

	// ErrPublishFailed is the error returned when publish failed.
	ErrPublishFailed = status.Errorf(codes.Internal, "publish failed")
)

// MessageIdGenerator is the function that generates the message id.
// The default message id generator generates a random u64 and encodes it into a string of base 10.
type MessageIdGenerator func() string

// Publisher is the gRPC server that provides the pub/sub service.
type Publisher struct {
	// pubsubServer is the real gRPC server that provides the pub/sub service.
	pubsubServer *pubsubServer

	// messageIdGenerator is the function that generates the unique message id.
	messageIdGenerator MessageIdGenerator

	// logger is the zap logger of the publisher.
	logger *zap.Logger
}

// PublisherOptions is the options for creating the publisher.
type PublisherOptions struct {
	// ServicePort is the port that the publisher listens on.
	ServicePort string

	// ProviderOptions is the options for creating the provider.
	ProviderOptions *provider.Options

	// QueueLength is the length of the message queue of the subscribers.
	QueueLength int

	// EnableDebug enables debug logging.
	EnableDebug bool

	// MessageIdGenerator is the function that generates the message id.
	MessageIdGenerator MessageIdGenerator
}

// StartPublisher creates a new publisher with the given options and starts the gRPC server.
func StartPublisher(opts *PublisherOptions) (*Publisher, error) {
	p, err := factory.NewProvider(opts.ProviderOptions)
	if err != nil {
		return nil, err
	}

	logger, err := newLogger(opts.EnableDebug)
	if err != nil {
		return nil, err
	}

	pubsServer := &pubsubServer{
		provider:    p,
		servicePort: opts.ServicePort,

		queue:       make(map[string]chan *pb.Message),
		queueLength: opts.QueueLength,

		logger: logger,
	}
	go func() {
		if err := pubsServer.start(); err != nil {
			panic(err)
		}
	}()

	return &Publisher{
		pubsubServer:       pubsServer,
		messageIdGenerator: opts.MessageIdGenerator,
		logger:             logger,
	}, nil
}

// DefaultPublisherOptions returns the default options for creating the publisher.
func DefaultPublisherOptions() *PublisherOptions {
	return &PublisherOptions{
		ServicePort: ":14201",
		EnableDebug: true,
		QueueLength: 100,
		ProviderOptions: &provider.Options{
			StorageEndpoints: []string{"localhost:2379"},
			StoreType:        provider.StoreTypeEtcd,
		},
		MessageIdGenerator: messageId,
	}
}

// NewMessage creates a new message of the given topic and payload.
func (p *Publisher) NewMessage(topic string, payload []byte) *pb.Message {
	return &pb.Message{
		Topic:     topic,
		Payload:   payload,
		MessageId: p.messageIdGenerator(),
		Timestamp: time.Now().UnixNano(),
	}
}

// Publish publishes the given message to the subscribers.
func (p *Publisher) Publish(ctx context.Context, message *pb.Message) error {
	return p.pubsubServer.Publish(ctx, message)
}

// Stop stops the publisher.
func (p *Publisher) Stop() {
	p.pubsubServer.stop()
}

// Subscriber is the gRPC client that consumes the pub/sub service.
type Subscriber struct {
	pubsubClient pb.PubSubClient
	conn         *grpc.ClientConn
	subscriberId string
	queue        chan *pb.Message
	errChan      chan error
}

// SubscriberOptions is the options for creating the subscriber.
type SubscriberOptions struct {
	// PubSubEndpoint is the endpoint of the pub/sub service.
	PubSubEndpoint string

	// SubscriberId is the id of the subscriber.
	// The subscriber id should be unique.
	SubscriberId string

	// QueueLength is the client side queue length.
	QueueLength int
}

// NewSubscriber creates a new subscriber with the given options.
func NewSubscriber(opts *SubscriberOptions) (*Subscriber, error) {
	conn, err := grpc.Dial(opts.PubSubEndpoint, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	return &Subscriber{
		pubsubClient: pb.NewPubSubClient(conn),
		subscriberId: opts.SubscriberId,
		queue:        make(chan *pb.Message, opts.QueueLength),
		errChan:      make(chan error, 1),
		conn:         conn,
	}, nil
}

// DefaultSubscriberOptions returns the default options for creating the subscriber.
func DefaultSubscriberOptions() *SubscriberOptions {
	return &SubscriberOptions{
		PubSubEndpoint: DefaultPublisherOptions().ServicePort,
		QueueLength:    100,
		SubscriberId:   "pubsub-client",
	}
}

// Subscribe subscribes the given topics and receives messages.
// After call Subscribe, you should call Recv to receive messages.
func (s *Subscriber) Subscribe(ctx context.Context, topics []string) error {
	stream, err := s.pubsubClient.Subscribe(ctx, &pb.SubscribeRequest{
		SubscriberId: s.subscriberId,
		Topics:       topics,
	})
	if err != nil {
		return err
	}

	go func() {
		defer func() {
			close(s.queue)
			close(s.errChan)
		}()
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				return // The stream is closed by the server.
			}
			if err != nil {
				s.errChan <- err
				return
			}

			// Send the puback to the publisher when receiving the message.
			_, err = s.pubsubClient.PublishAck(ctx, &pb.PublishAckRequest{
				MessageId:    message.GetMessageId(),
				Topic:        message.GetTopic(),
				SubscriberId: s.subscriberId,
			})

			s.queue <- message
		}
	}()

	return nil
}

// Recv receives a message from the publisher.
func (s *Subscriber) Recv() (*pb.Message, error) {
	select {
	case message := <-s.queue:
		return message, nil
	case err := <-s.errChan:
		return nil, err
	}
}

// Close closes the subscriber.
func (s *Subscriber) Close() {
	s.conn.Close()
}

type pubsubServer struct {
	// servicePort is the port that the publisher listens on.
	servicePort string

	// provider is the backend storage provider.
	provider provider.Provider

	// mu protects queue.
	mu sync.RWMutex

	// queue stores the message queue of the subscribers in memory.
	queue map[string]chan *pb.Message

	// queuesLength stores the length of the message queue of the subscribers.
	queueLength int

	grpcServer *grpc.Server
	logger     *zap.Logger
	pb.UnimplementedPubSubServer
}

func (p *pubsubServer) Subscribe(req *pb.SubscribeRequest, stream pb.PubSub_SubscribeServer) error {
	p.logger.Debug("Received subscribe request", zap.Any("request", req))

	if err := p.validateSubscribeRequest(req); err != nil {
		return err
	}

	offlineMessages, err := p.provider.Subscribe(context.Background(), req.GetSubscriberId(), req.GetTopics())
	if err != nil {
		p.logger.Error("Subscribe failed", zap.Error(err))
		return ErrSubscribeFailed
	}

	// Create a fixed length message stream for the subscriber.
	p.mu.Lock()
	if _, ok := p.queue[req.GetSubscriberId()]; !ok {
		p.queue[req.GetSubscriberId()] = make(chan *pb.Message, p.queueLength)
	}
	p.mu.Unlock()

	// Send offline messages to the subscriber.
	for _, message := range offlineMessages {
		if err := stream.Send(message); err != nil {
			p.logger.Error("Subscribe failed", zap.Error(err))
			return ErrSubscribeFailed
		}
	}

	p.mu.RLock()
	queue := p.queue[req.GetSubscriberId()]
	p.mu.RUnlock()

	// Remove the message stream when the subscriber is disconnected.
	ctx := stream.Context()
	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			p.mu.Lock()
			delete(p.queue, req.GetSubscriberId())
			p.mu.Unlock()
		}
	}()

	// Stream messages to the subscriber.
	for {
		select {
		case message := <-queue:
			if err := stream.Send(message); err != nil {
				p.logger.Error("Subscribe failed", zap.Error(err))
				return ErrSubscribeFailed
			}
		}
	}
}

func (p *pubsubServer) Unsubscribe(ctx context.Context, req *pb.UnsubscribeRequest) (*pb.UnsubscribeResponse, error) {
	p.logger.Debug("Received unsubscribe request", zap.Any("request", req))

	if err := p.validateUnsubscribeRequest(req); err != nil {
		return nil, err
	}

	if err := p.provider.Unsubscribe(ctx, req.GetSubscriberId(), req.GetTopics()); err != nil {
		p.logger.Error("Unsubscribe failed", zap.Error(err))
		return nil, ErrUnsubscribeFailed
	}

	return &pb.UnsubscribeResponse{}, nil
}

func (p *pubsubServer) PublishAck(ctx context.Context, req *pb.PublishAckRequest) (*pb.PublishAckResponse, error) {
	p.logger.Debug("Received publish ack request", zap.Any("request", req))

	if err := p.validatePublishAckRequest(req); err != nil {
		return nil, err
	}

	// When receiving the puback, remove the offline message of the subscriber.
	if err := p.provider.RemoveOfflineMessage(ctx, req.GetSubscriberId(), req.GetTopic(), req.GetMessageId()); err != nil {
		p.logger.Error("Publish ack failed", zap.Error(err))
		return nil, ErrPublishAckFailed
	}

	return &pb.PublishAckResponse{}, nil
}

func (p *pubsubServer) Publish(ctx context.Context, message *pb.Message) error {
	p.logger.Debug("Received publish request", zap.Any("request", message))

	if err := p.validateMessage(message); err != nil {
		return err
	}

	subscribers, err := p.provider.GetSubscribers(ctx, message.GetTopic())
	if err != nil {
		p.logger.Error("Get subscribers failed", zap.Error(err))
		return ErrGetSubscribersFailed
	}

	if len(subscribers) == 0 {
		return ErrSubscribersNotFound
	}

	// Store the message at first.
	if err := p.provider.StoreMessage(ctx, message, subscribers); err != nil {
		p.logger.Error("Publish failed", zap.Error(err))
		return ErrPublishFailed
	}

	// Send the message to the subscribers from the internal message queue.
	for _, subscriber := range subscribers {
		p.mu.RLock()
		if queue, ok := p.queue[subscriber]; ok {
			queue <- message
		}
		p.mu.RUnlock()
	}

	return nil
}

// start the gRPC server of the publisher.
func (p *pubsubServer) start() error {
	listener, err := net.Listen("tcp", p.servicePort)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterPubSubServer(s, p)
	p.grpcServer = s

	if err := s.Serve(listener); err != nil {
		return err
	}

	return nil
}

// stop the gRPC server of the publisher.
func (p *pubsubServer) stop() {
	p.grpcServer.Stop()
}

func (p *pubsubServer) validateSubscribeRequest(req *pb.SubscribeRequest) error {
	if len(req.GetTopics()) == 0 {
		return status.Errorf(codes.InvalidArgument, "no topics specified")
	}
	if req.GetSubscriberId() == "" {
		return status.Errorf(codes.InvalidArgument, "no subscriber id specified")
	}
	return nil
}

func (p *pubsubServer) validateUnsubscribeRequest(req *pb.UnsubscribeRequest) error {
	if len(req.GetTopics()) == 0 {
		return status.Errorf(codes.InvalidArgument, "no topics specified")
	}
	if req.GetSubscriberId() == "" {
		return status.Errorf(codes.InvalidArgument, "no subscriber id specified")
	}
	return nil
}

func (p *pubsubServer) validatePublishAckRequest(req *pb.PublishAckRequest) error {
	if req.GetSubscriberId() == "" {
		return status.Errorf(codes.InvalidArgument, "no subscriber id specified")
	}
	if req.GetTopic() == "" {
		return status.Errorf(codes.InvalidArgument, "no topic specified")
	}
	if req.GetMessageId() == "" {
		return status.Errorf(codes.InvalidArgument, "no message id specified")
	}
	return nil
}

func (p *pubsubServer) validateMessage(msg *pb.Message) error {
	if msg.GetTopic() == "" {
		return status.Errorf(codes.InvalidArgument, "no topic specified")
	}
	if msg.GetMessageId() == "" {
		return status.Errorf(codes.InvalidArgument, "no message id specified")
	}
	if len(msg.GetPayload()) == 0 {
		return status.Errorf(codes.InvalidArgument, "empty payload")
	}
	return nil
}

func newLogger(enableDebug bool) (*zap.Logger, error) {
	config := zap.NewProductionConfig()

	// Set a human-readable time format.
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Set DebugLevel if enableDebug is true.
	if enableDebug {
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}

func messageId() string {
	// Generate a random u64 and encode it into a string.
	rand.NewSource(time.Now().UnixNano())
	return strconv.FormatUint(rand.Uint64(), 10)
}
