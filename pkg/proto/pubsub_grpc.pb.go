// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: pubsub.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PubSubClient is the client API for PubSub service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PubSubClient interface {
	// Subscribe subscribes the topics and receives messages.
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (PubSub_SubscribeClient, error)
	// Unsubscribe unsubscribes the topics.
	Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*UnsubscribeResponse, error)
	// PublishAck is used to acknowledge the message received.
	// When subscriber re-subscribe the topic, it will receive the offline messages.
	PublishAck(ctx context.Context, in *PublishAckRequest, opts ...grpc.CallOption) (*PublishAckResponse, error)
}

type pubSubClient struct {
	cc grpc.ClientConnInterface
}

func NewPubSubClient(cc grpc.ClientConnInterface) PubSubClient {
	return &pubSubClient{cc}
}

func (c *pubSubClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (PubSub_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &PubSub_ServiceDesc.Streams[0], "/pubsub.PubSub/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &pubSubSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PubSub_SubscribeClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type pubSubSubscribeClient struct {
	grpc.ClientStream
}

func (x *pubSubSubscribeClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pubSubClient) Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*UnsubscribeResponse, error) {
	out := new(UnsubscribeResponse)
	err := c.cc.Invoke(ctx, "/pubsub.PubSub/Unsubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSubClient) PublishAck(ctx context.Context, in *PublishAckRequest, opts ...grpc.CallOption) (*PublishAckResponse, error) {
	out := new(PublishAckResponse)
	err := c.cc.Invoke(ctx, "/pubsub.PubSub/PublishAck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PubSubServer is the server API for PubSub service.
// All implementations must embed UnimplementedPubSubServer
// for forward compatibility
type PubSubServer interface {
	// Subscribe subscribes the topics and receives messages.
	Subscribe(*SubscribeRequest, PubSub_SubscribeServer) error
	// Unsubscribe unsubscribes the topics.
	Unsubscribe(context.Context, *UnsubscribeRequest) (*UnsubscribeResponse, error)
	// PublishAck is used to acknowledge the message received.
	// When subscriber re-subscribe the topic, it will receive the offline messages.
	PublishAck(context.Context, *PublishAckRequest) (*PublishAckResponse, error)
	mustEmbedUnimplementedPubSubServer()
}

// UnimplementedPubSubServer must be embedded to have forward compatible implementations.
type UnimplementedPubSubServer struct {
}

func (UnimplementedPubSubServer) Subscribe(*SubscribeRequest, PubSub_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedPubSubServer) Unsubscribe(context.Context, *UnsubscribeRequest) (*UnsubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsubscribe not implemented")
}
func (UnimplementedPubSubServer) PublishAck(context.Context, *PublishAckRequest) (*PublishAckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishAck not implemented")
}
func (UnimplementedPubSubServer) mustEmbedUnimplementedPubSubServer() {}

// UnsafePubSubServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PubSubServer will
// result in compilation errors.
type UnsafePubSubServer interface {
	mustEmbedUnimplementedPubSubServer()
}

func RegisterPubSubServer(s grpc.ServiceRegistrar, srv PubSubServer) {
	s.RegisterService(&PubSub_ServiceDesc, srv)
}

func _PubSub_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PubSubServer).Subscribe(m, &pubSubSubscribeServer{stream})
}

type PubSub_SubscribeServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type pubSubSubscribeServer struct {
	grpc.ServerStream
}

func (x *pubSubSubscribeServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _PubSub_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSubServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pubsub.PubSub/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSubServer).Unsubscribe(ctx, req.(*UnsubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubSub_PublishAck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishAckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSubServer).PublishAck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pubsub.PubSub/PublishAck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSubServer).PublishAck(ctx, req.(*PublishAckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PubSub_ServiceDesc is the grpc.ServiceDesc for PubSub service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PubSub_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pubsub.PubSub",
	HandlerType: (*PubSubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unsubscribe",
			Handler:    _PubSub_Unsubscribe_Handler,
		},
		{
			MethodName: "PublishAck",
			Handler:    _PubSub_PublishAck_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _PubSub_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pubsub.proto",
}
