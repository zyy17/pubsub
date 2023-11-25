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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: pubsub.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Publisher publishes Message of a Topic to Subscribers.
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// message_id is a unique identifier for this message.
	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	// topic is the topic this message belongs to. Each message belongs to exactly one topic.
	Topic string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	// timestamp is the unix timestamp when this message is published.
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// payload is the content of this message.
	Payload []byte `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pubsub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_pubsub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_pubsub_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *Message) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Message) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Message) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

// SubscribeRequest is the request to subscribe topics.
type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topics       []string `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
	SubscriberId string   `protobuf:"bytes,2,opt,name=subscriber_id,json=subscriberId,proto3" json:"subscriber_id,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pubsub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pubsub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_pubsub_proto_rawDescGZIP(), []int{1}
}

func (x *SubscribeRequest) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

func (x *SubscribeRequest) GetSubscriberId() string {
	if x != nil {
		return x.SubscriberId
	}
	return ""
}

type UnsubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topics              []string `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
	SubscriberId        string   `protobuf:"bytes,2,opt,name=subscriber_id,json=subscriberId,proto3" json:"subscriber_id,omitempty"`
	DropOfflineMessages bool     `protobuf:"varint,3,opt,name=drop_offline_messages,json=dropOfflineMessages,proto3" json:"drop_offline_messages,omitempty"`
}

func (x *UnsubscribeRequest) Reset() {
	*x = UnsubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pubsub_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeRequest) ProtoMessage() {}

func (x *UnsubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pubsub_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeRequest.ProtoReflect.Descriptor instead.
func (*UnsubscribeRequest) Descriptor() ([]byte, []int) {
	return file_pubsub_proto_rawDescGZIP(), []int{2}
}

func (x *UnsubscribeRequest) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

func (x *UnsubscribeRequest) GetSubscriberId() string {
	if x != nil {
		return x.SubscriberId
	}
	return ""
}

func (x *UnsubscribeRequest) GetDropOfflineMessages() bool {
	if x != nil {
		return x.DropOfflineMessages
	}
	return false
}

type UnsubscribeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnsubscribeResponse) Reset() {
	*x = UnsubscribeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pubsub_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeResponse) ProtoMessage() {}

func (x *UnsubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pubsub_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeResponse.ProtoReflect.Descriptor instead.
func (*UnsubscribeResponse) Descriptor() ([]byte, []int) {
	return file_pubsub_proto_rawDescGZIP(), []int{3}
}

// If the subscriber receives the message, it should send a PublishAckRequest back to the publisher.
type PublishAckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId    string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	SubscriberId string `protobuf:"bytes,2,opt,name=subscriber_id,json=subscriberId,proto3" json:"subscriber_id,omitempty"`
	Topic        string `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
}

func (x *PublishAckRequest) Reset() {
	*x = PublishAckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pubsub_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishAckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishAckRequest) ProtoMessage() {}

func (x *PublishAckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pubsub_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishAckRequest.ProtoReflect.Descriptor instead.
func (*PublishAckRequest) Descriptor() ([]byte, []int) {
	return file_pubsub_proto_rawDescGZIP(), []int{4}
}

func (x *PublishAckRequest) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *PublishAckRequest) GetSubscriberId() string {
	if x != nil {
		return x.SubscriberId
	}
	return ""
}

func (x *PublishAckRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

type PublishAckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublishAckResponse) Reset() {
	*x = PublishAckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pubsub_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishAckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishAckResponse) ProtoMessage() {}

func (x *PublishAckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pubsub_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishAckResponse.ProtoReflect.Descriptor instead.
func (*PublishAckResponse) Descriptor() ([]byte, []int) {
	return file_pubsub_proto_rawDescGZIP(), []int{5}
}

var File_pubsub_proto protoreflect.FileDescriptor

var file_pubsub_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x22, 0x76, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x4f,
	0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x85, 0x01, 0x0a, 0x12, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x23,
	0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x6f, 0x66, 0x66, 0x6c,
	0x69, 0x6e, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x13, 0x64, 0x72, 0x6f, 0x70, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x6e, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6d,
	0x0a, 0x11, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x22, 0x14, 0x0a,
	0x12, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xcf, 0x01, 0x0a, 0x06, 0x50, 0x75, 0x62, 0x53, 0x75, 0x62, 0x12, 0x38,
	0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x75,
	0x62, 0x73, 0x75, 0x62, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x30, 0x01, 0x12, 0x46, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62,
	0x2e, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x55, 0x6e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x43, 0x0a, 0x0a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x41, 0x63, 0x6b, 0x12, 0x19,
	0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x41,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x75, 0x62, 0x73,
	0x75, 0x62, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x79, 0x79, 0x31, 0x37, 0x2f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pubsub_proto_rawDescOnce sync.Once
	file_pubsub_proto_rawDescData = file_pubsub_proto_rawDesc
)

func file_pubsub_proto_rawDescGZIP() []byte {
	file_pubsub_proto_rawDescOnce.Do(func() {
		file_pubsub_proto_rawDescData = protoimpl.X.CompressGZIP(file_pubsub_proto_rawDescData)
	})
	return file_pubsub_proto_rawDescData
}

var file_pubsub_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pubsub_proto_goTypes = []interface{}{
	(*Message)(nil),             // 0: pubsub.Message
	(*SubscribeRequest)(nil),    // 1: pubsub.SubscribeRequest
	(*UnsubscribeRequest)(nil),  // 2: pubsub.UnsubscribeRequest
	(*UnsubscribeResponse)(nil), // 3: pubsub.UnsubscribeResponse
	(*PublishAckRequest)(nil),   // 4: pubsub.PublishAckRequest
	(*PublishAckResponse)(nil),  // 5: pubsub.PublishAckResponse
}
var file_pubsub_proto_depIdxs = []int32{
	1, // 0: pubsub.PubSub.Subscribe:input_type -> pubsub.SubscribeRequest
	2, // 1: pubsub.PubSub.Unsubscribe:input_type -> pubsub.UnsubscribeRequest
	4, // 2: pubsub.PubSub.PublishAck:input_type -> pubsub.PublishAckRequest
	0, // 3: pubsub.PubSub.Subscribe:output_type -> pubsub.Message
	3, // 4: pubsub.PubSub.Unsubscribe:output_type -> pubsub.UnsubscribeResponse
	5, // 5: pubsub.PubSub.PublishAck:output_type -> pubsub.PublishAckResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pubsub_proto_init() }
func file_pubsub_proto_init() {
	if File_pubsub_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pubsub_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pubsub_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pubsub_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsubscribeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pubsub_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsubscribeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pubsub_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishAckRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pubsub_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishAckResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pubsub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pubsub_proto_goTypes,
		DependencyIndexes: file_pubsub_proto_depIdxs,
		MessageInfos:      file_pubsub_proto_msgTypes,
	}.Build()
	File_pubsub_proto = out.File
	file_pubsub_proto_rawDesc = nil
	file_pubsub_proto_goTypes = nil
	file_pubsub_proto_depIdxs = nil
}