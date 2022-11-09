// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: transaction_contents.proto

package services

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

type SignedTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// TransactionBody serialized into bytes, which needs to be signed
	BodyBytes []byte `protobuf:"bytes,1,opt,name=bodyBytes,proto3" json:"bodyBytes,omitempty"`
	//*
	// The signatures on the body with the new format, to authorize the transaction
	SigMap *SignatureMap `protobuf:"bytes,2,opt,name=sigMap,proto3" json:"sigMap,omitempty"`
}

func (x *SignedTransaction) Reset() {
	*x = SignedTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_contents_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedTransaction) ProtoMessage() {}

func (x *SignedTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_contents_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedTransaction.ProtoReflect.Descriptor instead.
func (*SignedTransaction) Descriptor() ([]byte, []int) {
	return file_transaction_contents_proto_rawDescGZIP(), []int{0}
}

func (x *SignedTransaction) GetBodyBytes() []byte {
	if x != nil {
		return x.BodyBytes
	}
	return nil
}

func (x *SignedTransaction) GetSigMap() *SignatureMap {
	if x != nil {
		return x.SigMap
	}
	return nil
}

var File_transaction_contents_proto protoreflect.FileDescriptor

var file_transaction_contents_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x11, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x62,
	0x6f, 0x64, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x62, 0x6f, 0x64, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x69, 0x67,
	0x4d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x52, 0x06,
	0x73, 0x69, 0x67, 0x4d, 0x61, 0x70, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transaction_contents_proto_rawDescOnce sync.Once
	file_transaction_contents_proto_rawDescData = file_transaction_contents_proto_rawDesc
)

func file_transaction_contents_proto_rawDescGZIP() []byte {
	file_transaction_contents_proto_rawDescOnce.Do(func() {
		file_transaction_contents_proto_rawDescData = protoimpl.X.CompressGZIP(file_transaction_contents_proto_rawDescData)
	})
	return file_transaction_contents_proto_rawDescData
}

var file_transaction_contents_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_transaction_contents_proto_goTypes = []interface{}{
	(*SignedTransaction)(nil), // 0: proto.SignedTransaction
	(*SignatureMap)(nil),      // 1: proto.SignatureMap
}
var file_transaction_contents_proto_depIdxs = []int32{
	1, // 0: proto.SignedTransaction.sigMap:type_name -> proto.SignatureMap
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_transaction_contents_proto_init() }
func file_transaction_contents_proto_init() {
	if File_transaction_contents_proto != nil {
		return
	}
	file_basic_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_transaction_contents_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedTransaction); i {
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
			RawDescriptor: file_transaction_contents_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transaction_contents_proto_goTypes,
		DependencyIndexes: file_transaction_contents_proto_depIdxs,
		MessageInfos:      file_transaction_contents_proto_msgTypes,
	}.Build()
	File_transaction_contents_proto = out.File
	file_transaction_contents_proto_rawDesc = nil
	file_transaction_contents_proto_goTypes = nil
	file_transaction_contents_proto_depIdxs = nil
}
