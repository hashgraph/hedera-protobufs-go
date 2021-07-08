// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: TokenGrantKyc.proto

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

// Grants KYC to the account for the given token. Must be signed by the Token's kycKey.
//If the provided account is not found, the transaction will resolve to INVALID_ACCOUNT_ID.
//If the provided account has been deleted, the transaction will resolve to ACCOUNT_DELETED.
//If the provided token is not found, the transaction will resolve to INVALID_TOKEN_ID.
//If the provided token has been deleted, the transaction will resolve to TOKEN_WAS_DELETED.
//If an Association between the provided token and account is not found, the transaction will resolve to TOKEN_NOT_ASSOCIATED_TO_ACCOUNT.
//If no KYC Key is defined, the transaction will resolve to TOKEN_HAS_NO_KYC_KEY.
//Once executed the Account is marked as KYC Granted.
type TokenGrantKycTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token   *TokenID   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`     // The token for which this account will be granted KYC. If token does not exist, transaction results in INVALID_TOKEN_ID
	Account *AccountID `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"` // The account to be KYCed
}

func (x *TokenGrantKycTransactionBody) Reset() {
	*x = TokenGrantKycTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TokenGrantKyc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenGrantKycTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenGrantKycTransactionBody) ProtoMessage() {}

func (x *TokenGrantKycTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_TokenGrantKyc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenGrantKycTransactionBody.ProtoReflect.Descriptor instead.
func (*TokenGrantKycTransactionBody) Descriptor() ([]byte, []int) {
	return file_TokenGrantKyc_proto_rawDescGZIP(), []int{0}
}

func (x *TokenGrantKycTransactionBody) GetToken() *TokenID {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *TokenGrantKycTransactionBody) GetAccount() *AccountID {
	if x != nil {
		return x.Account
	}
	return nil
}

var File_TokenGrantKyc_proto protoreflect.FileDescriptor

var file_TokenGrantKyc_proto_rawDesc = []byte{
	0x0a, 0x13, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x4b, 0x79, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x42, 0x61,
	0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70,
	0x0a, 0x1c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x4b, 0x79, 0x63, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x24,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2a, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61,
	0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TokenGrantKyc_proto_rawDescOnce sync.Once
	file_TokenGrantKyc_proto_rawDescData = file_TokenGrantKyc_proto_rawDesc
)

func file_TokenGrantKyc_proto_rawDescGZIP() []byte {
	file_TokenGrantKyc_proto_rawDescOnce.Do(func() {
		file_TokenGrantKyc_proto_rawDescData = protoimpl.X.CompressGZIP(file_TokenGrantKyc_proto_rawDescData)
	})
	return file_TokenGrantKyc_proto_rawDescData
}

var file_TokenGrantKyc_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TokenGrantKyc_proto_goTypes = []interface{}{
	(*TokenGrantKycTransactionBody)(nil), // 0: proto.TokenGrantKycTransactionBody
	(*TokenID)(nil),                      // 1: proto.TokenID
	(*AccountID)(nil),                    // 2: proto.AccountID
}
var file_TokenGrantKyc_proto_depIdxs = []int32{
	1, // 0: proto.TokenGrantKycTransactionBody.token:type_name -> proto.TokenID
	2, // 1: proto.TokenGrantKycTransactionBody.account:type_name -> proto.AccountID
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_TokenGrantKyc_proto_init() }
func file_TokenGrantKyc_proto_init() {
	if File_TokenGrantKyc_proto != nil {
		return
	}
	file_BasicTypes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TokenGrantKyc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenGrantKycTransactionBody); i {
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
			RawDescriptor: file_TokenGrantKyc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TokenGrantKyc_proto_goTypes,
		DependencyIndexes: file_TokenGrantKyc_proto_depIdxs,
		MessageInfos:      file_TokenGrantKyc_proto_msgTypes,
	}.Build()
	File_TokenGrantKyc_proto = out.File
	file_TokenGrantKyc_proto_rawDesc = nil
	file_TokenGrantKyc_proto_goTypes = nil
	file_TokenGrantKyc_proto_depIdxs = nil
}
