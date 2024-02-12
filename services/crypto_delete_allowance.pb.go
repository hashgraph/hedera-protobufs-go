// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: crypto_delete_allowance.proto

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
// Deletes one or more non-fungible approved allowances from an owner's account. This operation
// will remove the allowances granted to one or more specific non-fungible token serial numbers. Each owner account
// listed as wiping an allowance must sign the transaction. Hbar and fungible token allowances
// can be removed by setting the amount to zero in CryptoApproveAllowance.
type CryptoDeleteAllowanceTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// List of non-fungible token allowances to remove.
	NftAllowances []*NftRemoveAllowance `protobuf:"bytes,2,rep,name=nftAllowances,proto3" json:"nftAllowances,omitempty"`
}

func (x *CryptoDeleteAllowanceTransactionBody) Reset() {
	*x = CryptoDeleteAllowanceTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_delete_allowance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoDeleteAllowanceTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoDeleteAllowanceTransactionBody) ProtoMessage() {}

func (x *CryptoDeleteAllowanceTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_delete_allowance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoDeleteAllowanceTransactionBody.ProtoReflect.Descriptor instead.
func (*CryptoDeleteAllowanceTransactionBody) Descriptor() ([]byte, []int) {
	return file_crypto_delete_allowance_proto_rawDescGZIP(), []int{0}
}

func (x *CryptoDeleteAllowanceTransactionBody) GetNftAllowances() []*NftRemoveAllowance {
	if x != nil {
		return x.NftAllowances
	}
	return nil
}

// *
// Nft allowances to be removed on an owner account
type NftRemoveAllowance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// The token that the allowance pertains to.
	TokenId *TokenID `protobuf:"bytes,1,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	// *
	// The account ID of the token owner (ie. the grantor of the allowance).
	Owner *AccountID `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// *
	// The list of serial numbers to remove allowances from.
	SerialNumbers []int64 `protobuf:"varint,3,rep,packed,name=serial_numbers,json=serialNumbers,proto3" json:"serial_numbers,omitempty"`
}

func (x *NftRemoveAllowance) Reset() {
	*x = NftRemoveAllowance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_delete_allowance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NftRemoveAllowance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftRemoveAllowance) ProtoMessage() {}

func (x *NftRemoveAllowance) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_delete_allowance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftRemoveAllowance.ProtoReflect.Descriptor instead.
func (*NftRemoveAllowance) Descriptor() ([]byte, []int) {
	return file_crypto_delete_allowance_proto_rawDescGZIP(), []int{1}
}

func (x *NftRemoveAllowance) GetTokenId() *TokenID {
	if x != nil {
		return x.TokenId
	}
	return nil
}

func (x *NftRemoveAllowance) GetOwner() *AccountID {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *NftRemoveAllowance) GetSerialNumbers() []int64 {
	if x != nil {
		return x.SerialNumbers
	}
	return nil
}

var File_crypto_delete_allowance_proto protoreflect.FileDescriptor

var file_crypto_delete_allowance_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x24, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e,
	0x63, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64,
	0x79, 0x12, 0x3f, 0x0a, 0x0d, 0x6e, 0x66, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4e, 0x66, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x0d, 0x6e, 0x66, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x22, 0x8e, 0x01, 0x0a, 0x12, 0x4e, 0x66, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x52, 0x07, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_crypto_delete_allowance_proto_rawDescOnce sync.Once
	file_crypto_delete_allowance_proto_rawDescData = file_crypto_delete_allowance_proto_rawDesc
)

func file_crypto_delete_allowance_proto_rawDescGZIP() []byte {
	file_crypto_delete_allowance_proto_rawDescOnce.Do(func() {
		file_crypto_delete_allowance_proto_rawDescData = protoimpl.X.CompressGZIP(file_crypto_delete_allowance_proto_rawDescData)
	})
	return file_crypto_delete_allowance_proto_rawDescData
}

var file_crypto_delete_allowance_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_crypto_delete_allowance_proto_goTypes = []interface{}{
	(*CryptoDeleteAllowanceTransactionBody)(nil), // 0: proto.CryptoDeleteAllowanceTransactionBody
	(*NftRemoveAllowance)(nil),                   // 1: proto.NftRemoveAllowance
	(*TokenID)(nil),                              // 2: proto.TokenID
	(*AccountID)(nil),                            // 3: proto.AccountID
}
var file_crypto_delete_allowance_proto_depIdxs = []int32{
	1, // 0: proto.CryptoDeleteAllowanceTransactionBody.nftAllowances:type_name -> proto.NftRemoveAllowance
	2, // 1: proto.NftRemoveAllowance.token_id:type_name -> proto.TokenID
	3, // 2: proto.NftRemoveAllowance.owner:type_name -> proto.AccountID
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_crypto_delete_allowance_proto_init() }
func file_crypto_delete_allowance_proto_init() {
	if File_crypto_delete_allowance_proto != nil {
		return
	}
	file_basic_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_crypto_delete_allowance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoDeleteAllowanceTransactionBody); i {
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
		file_crypto_delete_allowance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NftRemoveAllowance); i {
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
			RawDescriptor: file_crypto_delete_allowance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_crypto_delete_allowance_proto_goTypes,
		DependencyIndexes: file_crypto_delete_allowance_proto_depIdxs,
		MessageInfos:      file_crypto_delete_allowance_proto_msgTypes,
	}.Build()
	File_crypto_delete_allowance_proto = out.File
	file_crypto_delete_allowance_proto_rawDesc = nil
	file_crypto_delete_allowance_proto_goTypes = nil
	file_crypto_delete_allowance_proto_depIdxs = nil
}
