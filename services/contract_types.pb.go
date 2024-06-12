// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: contract_types.proto

package services

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// *
// Info about a contract account's nonce value.
// A nonce of a contract is only incremented when that contract creates another contract.
type ContractNonceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// Id of the contract
	ContractId *ContractID `protobuf:"bytes,1,opt,name=contract_id,json=contractId,proto3" json:"contract_id,omitempty"`
	// *
	// The current value of the contract account's nonce property
	Nonce int64 `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (x *ContractNonceInfo) Reset() {
	*x = ContractNonceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractNonceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractNonceInfo) ProtoMessage() {}

func (x *ContractNonceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_contract_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractNonceInfo.ProtoReflect.Descriptor instead.
func (*ContractNonceInfo) Descriptor() ([]byte, []int) {
	return file_contract_types_proto_rawDescGZIP(), []int{0}
}

func (x *ContractNonceInfo) GetContractId() *ContractID {
	if x != nil {
		return x.ContractId
	}
	return nil
}

func (x *ContractNonceInfo) GetNonce() int64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

var File_contract_types_proto protoreflect.FileDescriptor

var file_contract_types_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5d, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4e, 0x6f, 0x6e, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x52, 0x0a, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x42,
	0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73,
	0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contract_types_proto_rawDescOnce sync.Once
	file_contract_types_proto_rawDescData = file_contract_types_proto_rawDesc
)

func file_contract_types_proto_rawDescGZIP() []byte {
	file_contract_types_proto_rawDescOnce.Do(func() {
		file_contract_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_types_proto_rawDescData)
	})
	return file_contract_types_proto_rawDescData
}

var file_contract_types_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_contract_types_proto_goTypes = []interface{}{
	(*ContractNonceInfo)(nil), // 0: proto.ContractNonceInfo
	(*ContractID)(nil),        // 1: proto.ContractID
}
var file_contract_types_proto_depIdxs = []int32{
	1, // 0: proto.ContractNonceInfo.contract_id:type_name -> proto.ContractID
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_contract_types_proto_init() }
func file_contract_types_proto_init() {
	if File_contract_types_proto != nil {
		return
	}
	file_basic_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_contract_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractNonceInfo); i {
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
			RawDescriptor: file_contract_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contract_types_proto_goTypes,
		DependencyIndexes: file_contract_types_proto_depIdxs,
		MessageInfos:      file_contract_types_proto_msgTypes,
	}.Build()
	File_contract_types_proto = out.File
	file_contract_types_proto_rawDesc = nil
	file_contract_types_proto_goTypes = nil
	file_contract_types_proto_depIdxs = nil
}
