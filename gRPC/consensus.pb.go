// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: gRPC/consensus.proto

package gRPC

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SendToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId int32 `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

func (x *SendToken) Reset() {
	*x = SendToken{}
	mi := &file_gRPC_consensus_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendToken) ProtoMessage() {}

func (x *SendToken) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_consensus_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendToken.ProtoReflect.Descriptor instead.
func (*SendToken) Descriptor() ([]byte, []int) {
	return file_gRPC_consensus_proto_rawDescGZIP(), []int{0}
}

func (x *SendToken) GetNodeId() int32 {
	if x != nil {
		return x.NodeId
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
	mi := &file_gRPC_consensus_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_consensus_proto_msgTypes[1]
	if x != nil {
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
	return file_gRPC_consensus_proto_rawDescGZIP(), []int{1}
}

var File_gRPC_consensus_proto protoreflect.FileDescriptor

var file_gRPC_consensus_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x52, 0x50, 0x43, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x52, 0x50, 0x43, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a,
	0x09, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x6f, 0x64,
	0x65, 0x49, 0x64, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x32, 0x31, 0x0a, 0x09, 0x43,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x0f, 0x2e, 0x67, 0x52, 0x50, 0x43, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x1a, 0x0a, 0x2e, 0x67, 0x52, 0x50, 0x43, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x07,
	0x5a, 0x05, 0x2f, 0x67, 0x52, 0x50, 0x43, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gRPC_consensus_proto_rawDescOnce sync.Once
	file_gRPC_consensus_proto_rawDescData = file_gRPC_consensus_proto_rawDesc
)

func file_gRPC_consensus_proto_rawDescGZIP() []byte {
	file_gRPC_consensus_proto_rawDescOnce.Do(func() {
		file_gRPC_consensus_proto_rawDescData = protoimpl.X.CompressGZIP(file_gRPC_consensus_proto_rawDescData)
	})
	return file_gRPC_consensus_proto_rawDescData
}

var file_gRPC_consensus_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gRPC_consensus_proto_goTypes = []any{
	(*SendToken)(nil), // 0: gRPC.SendToken
	(*Void)(nil),      // 1: gRPC.Void
}
var file_gRPC_consensus_proto_depIdxs = []int32{
	0, // 0: gRPC.Consensus.Token:input_type -> gRPC.SendToken
	1, // 1: gRPC.Consensus.Token:output_type -> gRPC.Void
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gRPC_consensus_proto_init() }
func file_gRPC_consensus_proto_init() {
	if File_gRPC_consensus_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gRPC_consensus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gRPC_consensus_proto_goTypes,
		DependencyIndexes: file_gRPC_consensus_proto_depIdxs,
		MessageInfos:      file_gRPC_consensus_proto_msgTypes,
	}.Build()
	File_gRPC_consensus_proto = out.File
	file_gRPC_consensus_proto_rawDesc = nil
	file_gRPC_consensus_proto_goTypes = nil
	file_gRPC_consensus_proto_depIdxs = nil
}
