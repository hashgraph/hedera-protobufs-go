// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: crypto_delete.proto

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

// *
// Mark an account as deleted, moving all its current hbars to another account. It will remain in
// the ledger, marked as deleted, until it expires. Transfers into it a deleted account fail. But a
// deleted account can still have its expiration extended in the normal way.
type CryptoDeleteTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// The account ID which will receive all remaining hbars
	TransferAccountID *AccountID `protobuf:"bytes,1,opt,name=transferAccountID,proto3" json:"transferAccountID,omitempty"`
	// *
	// The account ID which should be deleted
	DeleteAccountID *AccountID `protobuf:"bytes,2,opt,name=deleteAccountID,proto3" json:"deleteAccountID,omitempty"`
}

func (x *CryptoDeleteTransactionBody) Reset() {
	*x = CryptoDeleteTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_delete_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoDeleteTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoDeleteTransactionBody) ProtoMessage() {}

func (x *CryptoDeleteTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_delete_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoDeleteTransactionBody.ProtoReflect.Descriptor instead.
func (*CryptoDeleteTransactionBody) Descriptor() ([]byte, []int) {
	return file_crypto_delete_proto_rawDescGZIP(), []int{0}
}

func (x *CryptoDeleteTransactionBody) GetTransferAccountID() *AccountID {
	if x != nil {
		return x.TransferAccountID
	}
	return nil
}

func (x *CryptoDeleteTransactionBody) GetDeleteAccountID() *AccountID {
	if x != nil {
		return x.DeleteAccountID
	}
	return nil
}

var File_crypto_delete_proto protoreflect.FileDescriptor

var file_crypto_delete_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x99, 0x01, 0x0a, 0x1b, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12,
	0x3e, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x11, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12,
	0x3a, 0x0a, 0x0f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x0f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x26, 0x0a, 0x22, 0x63,
	0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76,
	0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crypto_delete_proto_rawDescOnce sync.Once
	file_crypto_delete_proto_rawDescData = file_crypto_delete_proto_rawDesc
)

func file_crypto_delete_proto_rawDescGZIP() []byte {
	file_crypto_delete_proto_rawDescOnce.Do(func() {
		file_crypto_delete_proto_rawDescData = protoimpl.X.CompressGZIP(file_crypto_delete_proto_rawDescData)
	})
	return file_crypto_delete_proto_rawDescData
}

var file_crypto_delete_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_crypto_delete_proto_goTypes = []interface{}{
	(*CryptoDeleteTransactionBody)(nil), // 0: proto.CryptoDeleteTransactionBody
	(*AccountID)(nil),                   // 1: proto.AccountID
}
var file_crypto_delete_proto_depIdxs = []int32{
	1, // 0: proto.CryptoDeleteTransactionBody.transferAccountID:type_name -> proto.AccountID
	1, // 1: proto.CryptoDeleteTransactionBody.deleteAccountID:type_name -> proto.AccountID
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_crypto_delete_proto_init() }
func file_crypto_delete_proto_init() {
	if File_crypto_delete_proto != nil {
		return
	}
	file_basic_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_crypto_delete_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoDeleteTransactionBody); i {
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
			RawDescriptor: file_crypto_delete_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_crypto_delete_proto_goTypes,
		DependencyIndexes: file_crypto_delete_proto_depIdxs,
		MessageInfos:      file_crypto_delete_proto_msgTypes,
	}.Build()
	File_crypto_delete_proto = out.File
	file_crypto_delete_proto_rawDesc = nil
	file_crypto_delete_proto_goTypes = nil
	file_crypto_delete_proto_depIdxs = nil
}
