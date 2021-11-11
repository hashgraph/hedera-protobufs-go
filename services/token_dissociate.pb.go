// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: token_dissociate.proto

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

//*
// Dissociates the provided account with the provided tokens. Must be signed by the provided
// Account's key.
// If the provided account is not found, the transaction will resolve to INVALID_ACCOUNT_ID.
// If the provided account has been deleted, the transaction will resolve to ACCOUNT_DELETED.
// If any of the provided tokens is not found, the transaction will resolve to INVALID_TOKEN_REF.
// If any of the provided tokens has been deleted, the transaction will resolve to TOKEN_WAS_DELETED.
// If an association between the provided account and any of the tokens does not exist, the
// transaction will resolve to TOKEN_NOT_ASSOCIATED_TO_ACCOUNT.
// If a token has not been deleted and has not expired, and the user has a nonzero balance, the
// transaction will resolve to TRANSACTION_REQUIRES_ZERO_TOKEN_BALANCES.
// If a <b>fungible token</b> has expired, the user can disassociate even if their token balance is
// not zero.
// If a <b>non fungible token</b> has expired, the user can <b>not</b> disassociate if their token
// balance is not zero. The transaction will resolve to TRANSACTION_REQUIRED_ZERO_TOKEN_BALANCES.
// On success, associations between the provided account and tokens are removed.
type TokenDissociateTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// The account to be dissociated with the provided tokens
	Account *AccountID `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	//*
	// The tokens to be dissociated with the provided account
	Tokens []*TokenID `protobuf:"bytes,2,rep,name=tokens,proto3" json:"tokens,omitempty"`
}

func (x *TokenDissociateTransactionBody) Reset() {
	*x = TokenDissociateTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_dissociate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenDissociateTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenDissociateTransactionBody) ProtoMessage() {}

func (x *TokenDissociateTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_token_dissociate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenDissociateTransactionBody.ProtoReflect.Descriptor instead.
func (*TokenDissociateTransactionBody) Descriptor() ([]byte, []int) {
	return file_token_dissociate_proto_rawDescGZIP(), []int{0}
}

func (x *TokenDissociateTransactionBody) GetAccount() *AccountID {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *TokenDissociateTransactionBody) GetTokens() []*TokenID {
	if x != nil {
		return x.Tokens
	}
	return nil
}

var File_token_dissociate_proto protoreflect.FileDescriptor

var file_token_dissociate_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x64, 0x69, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x74, 0x0a, 0x1e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x69, 0x73, 0x73, 0x6f,
	0x63, 0x69, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x6f, 0x64, 0x79, 0x12, 0x2a, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x26, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44,
	0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e,
	0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_token_dissociate_proto_rawDescOnce sync.Once
	file_token_dissociate_proto_rawDescData = file_token_dissociate_proto_rawDesc
)

func file_token_dissociate_proto_rawDescGZIP() []byte {
	file_token_dissociate_proto_rawDescOnce.Do(func() {
		file_token_dissociate_proto_rawDescData = protoimpl.X.CompressGZIP(file_token_dissociate_proto_rawDescData)
	})
	return file_token_dissociate_proto_rawDescData
}

var file_token_dissociate_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_token_dissociate_proto_goTypes = []interface{}{
	(*TokenDissociateTransactionBody)(nil), // 0: proto.TokenDissociateTransactionBody
	(*AccountID)(nil),                      // 1: proto.AccountID
	(*TokenID)(nil),                        // 2: proto.TokenID
}
var file_token_dissociate_proto_depIdxs = []int32{
	1, // 0: proto.TokenDissociateTransactionBody.account:type_name -> proto.AccountID
	2, // 1: proto.TokenDissociateTransactionBody.tokens:type_name -> proto.TokenID
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_token_dissociate_proto_init() }
func file_token_dissociate_proto_init() {
	if File_token_dissociate_proto != nil {
		return
	}
	file_basic_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_token_dissociate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenDissociateTransactionBody); i {
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
			RawDescriptor: file_token_dissociate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_token_dissociate_proto_goTypes,
		DependencyIndexes: file_token_dissociate_proto_depIdxs,
		MessageInfos:      file_token_dissociate_proto_msgTypes,
	}.Build()
	File_token_dissociate_proto = out.File
	file_token_dissociate_proto_rawDesc = nil
	file_token_dissociate_proto_goTypes = nil
	file_token_dissociate_proto_depIdxs = nil
}
