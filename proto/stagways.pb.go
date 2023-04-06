// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: stagways.proto

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PushTask  string `protobuf:"bytes,1,opt,name=push_task,json=pushTask,proto3" json:"push_task,omitempty"`
	PushType  uint32 `protobuf:"varint,2,opt,name=push_type,json=pushType,proto3" json:"push_type,omitempty"`
	MsgBody   string `protobuf:"bytes,3,opt,name=msg_body,json=msgBody,proto3" json:"msg_body,omitempty"`
	Timestamp int64  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stagways_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_stagways_proto_msgTypes[0]
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
	return file_stagways_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetPushTask() string {
	if x != nil {
		return x.PushTask
	}
	return ""
}

func (x *Message) GetPushType() uint32 {
	if x != nil {
		return x.PushType
	}
	return 0
}

func (x *Message) GetMsgBody() string {
	if x != nil {
		return x.MsgBody
	}
	return ""
}

func (x *Message) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type PushMessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Channel uint32   `protobuf:"varint,2,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *PushMessageReq) Reset() {
	*x = PushMessageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stagways_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMessageReq) ProtoMessage() {}

func (x *PushMessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_stagways_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMessageReq.ProtoReflect.Descriptor instead.
func (*PushMessageReq) Descriptor() ([]byte, []int) {
	return file_stagways_proto_rawDescGZIP(), []int{1}
}

func (x *PushMessageReq) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *PushMessageReq) GetChannel() uint32 {
	if x != nil {
		return x.Channel
	}
	return 0
}

type PushMessageResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Receipt string `protobuf:"bytes,1,opt,name=receipt,proto3" json:"receipt,omitempty"`
}

func (x *PushMessageResp) Reset() {
	*x = PushMessageResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stagways_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMessageResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMessageResp) ProtoMessage() {}

func (x *PushMessageResp) ProtoReflect() protoreflect.Message {
	mi := &file_stagways_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMessageResp.ProtoReflect.Descriptor instead.
func (*PushMessageResp) Descriptor() ([]byte, []int) {
	return file_stagways_proto_rawDescGZIP(), []int{2}
}

func (x *PushMessageResp) GetReceipt() string {
	if x != nil {
		return x.Receipt
	}
	return ""
}

type AddChannelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel uint32 `protobuf:"varint,1,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *AddChannelReq) Reset() {
	*x = AddChannelReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stagways_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddChannelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddChannelReq) ProtoMessage() {}

func (x *AddChannelReq) ProtoReflect() protoreflect.Message {
	mi := &file_stagways_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddChannelReq.ProtoReflect.Descriptor instead.
func (*AddChannelReq) Descriptor() ([]byte, []int) {
	return file_stagways_proto_rawDescGZIP(), []int{3}
}

func (x *AddChannelReq) GetChannel() uint32 {
	if x != nil {
		return x.Channel
	}
	return 0
}

type AddChannelResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddChannelResp) Reset() {
	*x = AddChannelResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stagways_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddChannelResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddChannelResp) ProtoMessage() {}

func (x *AddChannelResp) ProtoReflect() protoreflect.Message {
	mi := &file_stagways_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddChannelResp.ProtoReflect.Descriptor instead.
func (*AddChannelResp) Descriptor() ([]byte, []int) {
	return file_stagways_proto_rawDescGZIP(), []int{4}
}

type DelChannelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel uint32 `protobuf:"varint,1,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *DelChannelReq) Reset() {
	*x = DelChannelReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stagways_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelChannelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelChannelReq) ProtoMessage() {}

func (x *DelChannelReq) ProtoReflect() protoreflect.Message {
	mi := &file_stagways_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelChannelReq.ProtoReflect.Descriptor instead.
func (*DelChannelReq) Descriptor() ([]byte, []int) {
	return file_stagways_proto_rawDescGZIP(), []int{5}
}

func (x *DelChannelReq) GetChannel() uint32 {
	if x != nil {
		return x.Channel
	}
	return 0
}

type DelChannelResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DelChannelResp) Reset() {
	*x = DelChannelResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stagways_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelChannelResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelChannelResp) ProtoMessage() {}

func (x *DelChannelResp) ProtoReflect() protoreflect.Message {
	mi := &file_stagways_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelChannelResp.ProtoReflect.Descriptor instead.
func (*DelChannelResp) Descriptor() ([]byte, []int) {
	return file_stagways_proto_rawDescGZIP(), []int{6}
}

var File_stagways_proto protoreflect.FileDescriptor

var file_stagways_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x74, 0x61, 0x67, 0x77, 0x61, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x53, 0x74, 0x61, 0x67, 0x77, 0x61, 0x79, 0x73, 0x22, 0x7c, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x74, 0x61,
	0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x73, 0x68, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x57, 0x0a, 0x0e, 0x50, 0x75, 0x73, 0x68,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x2b, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x53, 0x74,
	0x61, 0x67, 0x77, 0x61, 0x79, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x22, 0x2b, 0x0a, 0x0f, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x22, 0x29,
	0x0a, 0x0d, 0x41, 0x64, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x10, 0x0a, 0x0e, 0x41, 0x64, 0x64,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x29, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x32, 0x4d, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x12, 0x44, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x18, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x77, 0x61, 0x79, 0x73, 0x2e, 0x50, 0x75, 0x73, 0x68,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x53, 0x74, 0x61,
	0x67, 0x77, 0x61, 0x79, 0x73, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x32, 0x8d, 0x01, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x12, 0x41, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12,
	0x17, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x77, 0x61, 0x79, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x77,
	0x61, 0x79, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x12, 0x17, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x77, 0x61, 0x79, 0x73, 0x2e, 0x44, 0x65,
	0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x53, 0x74,
	0x61, 0x67, 0x77, 0x61, 0x79, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stagways_proto_rawDescOnce sync.Once
	file_stagways_proto_rawDescData = file_stagways_proto_rawDesc
)

func file_stagways_proto_rawDescGZIP() []byte {
	file_stagways_proto_rawDescOnce.Do(func() {
		file_stagways_proto_rawDescData = protoimpl.X.CompressGZIP(file_stagways_proto_rawDescData)
	})
	return file_stagways_proto_rawDescData
}

var file_stagways_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_stagways_proto_goTypes = []interface{}{
	(*Message)(nil),         // 0: Stagways.Message
	(*PushMessageReq)(nil),  // 1: Stagways.PushMessageReq
	(*PushMessageResp)(nil), // 2: Stagways.PushMessageResp
	(*AddChannelReq)(nil),   // 3: Stagways.AddChannelReq
	(*AddChannelResp)(nil),  // 4: Stagways.AddChannelResp
	(*DelChannelReq)(nil),   // 5: Stagways.DelChannelReq
	(*DelChannelResp)(nil),  // 6: Stagways.DelChannelResp
}
var file_stagways_proto_depIdxs = []int32{
	0, // 0: Stagways.PushMessageReq.message:type_name -> Stagways.Message
	1, // 1: Stagways.Proxy.PushMessage:input_type -> Stagways.PushMessageReq
	3, // 2: Stagways.Agent.AddChannel:input_type -> Stagways.AddChannelReq
	5, // 3: Stagways.Agent.DelChannel:input_type -> Stagways.DelChannelReq
	2, // 4: Stagways.Proxy.PushMessage:output_type -> Stagways.PushMessageResp
	4, // 5: Stagways.Agent.AddChannel:output_type -> Stagways.AddChannelResp
	6, // 6: Stagways.Agent.DelChannel:output_type -> Stagways.DelChannelResp
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_stagways_proto_init() }
func file_stagways_proto_init() {
	if File_stagways_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stagways_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_stagways_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMessageReq); i {
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
		file_stagways_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMessageResp); i {
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
		file_stagways_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddChannelReq); i {
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
		file_stagways_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddChannelResp); i {
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
		file_stagways_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelChannelReq); i {
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
		file_stagways_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelChannelResp); i {
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
			RawDescriptor: file_stagways_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_stagways_proto_goTypes,
		DependencyIndexes: file_stagways_proto_depIdxs,
		MessageInfos:      file_stagways_proto_msgTypes,
	}.Build()
	File_stagways_proto = out.File
	file_stagways_proto_rawDesc = nil
	file_stagways_proto_goTypes = nil
	file_stagways_proto_depIdxs = nil
}
