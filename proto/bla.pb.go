// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: proto/bla.proto

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

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bla_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bla_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_proto_bla_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Node) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RequestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User      *Node `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Timestamp int32 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *RequestMessage) Reset() {
	*x = RequestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bla_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMessage) ProtoMessage() {}

func (x *RequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bla_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMessage.ProtoReflect.Descriptor instead.
func (*RequestMessage) Descriptor() ([]byte, []int) {
	return file_proto_bla_proto_rawDescGZIP(), []int{1}
}

func (x *RequestMessage) GetUser() *Node {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *RequestMessage) GetTimestamp() int32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type ReplyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User      *Node `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Timestamp int32 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ReplyMessage) Reset() {
	*x = ReplyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bla_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyMessage) ProtoMessage() {}

func (x *ReplyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bla_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyMessage.ProtoReflect.Descriptor instead.
func (*ReplyMessage) Descriptor() ([]byte, []int) {
	return file_proto_bla_proto_rawDescGZIP(), []int{2}
}

func (x *ReplyMessage) GetUser() *Node {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *ReplyMessage) GetTimestamp() int32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bla_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bla_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_proto_bla_proto_rawDescGZIP(), []int{3}
}

type OkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OkResponse) Reset() {
	*x = OkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bla_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OkResponse) ProtoMessage() {}

func (x *OkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bla_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OkResponse.ProtoReflect.Descriptor instead.
func (*OkResponse) Descriptor() ([]byte, []int) {
	return file_proto_bla_proto_rawDescGZIP(), []int{4}
}

var File_proto_bla_proto protoreflect.FileDescriptor

var file_proto_bla_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6c, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4f, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x4d, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x22, 0x06, 0x0a, 0x04, 0x76, 0x6f, 0x69, 0x64, 0x22, 0x0c, 0x0a, 0x0a,
	0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xbe, 0x01, 0x0a, 0x10, 0x45,
	0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3e, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61,
	0x6c, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12,
	0x36, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x76, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0c, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x42, 0x3d, 0x5a, 0x3b, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x61, 0x72, 0x6c, 0x39, 0x2f, 0x4f, 0x6e, 0x65, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x2f, 0x53, 0x6b, 0x72, 0x69, 0x76, 0x65, 0x62, 0x6f, 0x72, 0x64, 0x2f, 0x62,
	0x65, 0x74, 0x74, 0x65, 0x72, 0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x45, 0x78, 0x63, 0x6c, 0x75,
	0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_bla_proto_rawDescOnce sync.Once
	file_proto_bla_proto_rawDescData = file_proto_bla_proto_rawDesc
)

func file_proto_bla_proto_rawDescGZIP() []byte {
	file_proto_bla_proto_rawDescOnce.Do(func() {
		file_proto_bla_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_bla_proto_rawDescData)
	})
	return file_proto_bla_proto_rawDescData
}

var file_proto_bla_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_bla_proto_goTypes = []interface{}{
	(*Node)(nil),           // 0: proto.Node
	(*RequestMessage)(nil), // 1: proto.RequestMessage
	(*ReplyMessage)(nil),   // 2: proto.ReplyMessage
	(*Void)(nil),           // 3: proto.void
	(*OkResponse)(nil),     // 4: proto.okResponse
}
var file_proto_bla_proto_depIdxs = []int32{
	0, // 0: proto.RequestMessage.user:type_name -> proto.Node
	0, // 1: proto.ReplyMessage.user:type_name -> proto.Node
	1, // 2: proto.ExclusionService.AccessCritical:input_type -> proto.RequestMessage
	1, // 3: proto.ExclusionService.ReceiveRequest:input_type -> proto.RequestMessage
	2, // 4: proto.ExclusionService.ReceiveReply:input_type -> proto.ReplyMessage
	2, // 5: proto.ExclusionService.AccessCritical:output_type -> proto.ReplyMessage
	3, // 6: proto.ExclusionService.ReceiveRequest:output_type -> proto.void
	3, // 7: proto.ExclusionService.ReceiveReply:output_type -> proto.void
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_bla_proto_init() }
func file_proto_bla_proto_init() {
	if File_proto_bla_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_bla_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_proto_bla_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMessage); i {
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
		file_proto_bla_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyMessage); i {
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
		file_proto_bla_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
		file_proto_bla_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OkResponse); i {
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
			RawDescriptor: file_proto_bla_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_bla_proto_goTypes,
		DependencyIndexes: file_proto_bla_proto_depIdxs,
		MessageInfos:      file_proto_bla_proto_msgTypes,
	}.Build()
	File_proto_bla_proto = out.File
	file_proto_bla_proto_rawDesc = nil
	file_proto_bla_proto_goTypes = nil
	file_proto_bla_proto_depIdxs = nil
}
