// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: crypto_get_account_balance.proto

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
// Get the balance of a cryptocurrency account. This returns only the balance, so it is a smaller
// reply than CryptoGetInfo, which returns the balance plus additional information.
type CryptoGetAccountBalanceQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// Standard info sent from client to node, including the signed payment, and what kind of
	// response is requested (cost, state proof, both, or neither).
	Header *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// Types that are assignable to BalanceSource:
	//
	//	*CryptoGetAccountBalanceQuery_AccountID
	//	*CryptoGetAccountBalanceQuery_ContractID
	BalanceSource isCryptoGetAccountBalanceQuery_BalanceSource `protobuf_oneof:"balanceSource"`
}

func (x *CryptoGetAccountBalanceQuery) Reset() {
	*x = CryptoGetAccountBalanceQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_get_account_balance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoGetAccountBalanceQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoGetAccountBalanceQuery) ProtoMessage() {}

func (x *CryptoGetAccountBalanceQuery) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_get_account_balance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoGetAccountBalanceQuery.ProtoReflect.Descriptor instead.
func (*CryptoGetAccountBalanceQuery) Descriptor() ([]byte, []int) {
	return file_crypto_get_account_balance_proto_rawDescGZIP(), []int{0}
}

func (x *CryptoGetAccountBalanceQuery) GetHeader() *QueryHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (m *CryptoGetAccountBalanceQuery) GetBalanceSource() isCryptoGetAccountBalanceQuery_BalanceSource {
	if m != nil {
		return m.BalanceSource
	}
	return nil
}

func (x *CryptoGetAccountBalanceQuery) GetAccountID() *AccountID {
	if x, ok := x.GetBalanceSource().(*CryptoGetAccountBalanceQuery_AccountID); ok {
		return x.AccountID
	}
	return nil
}

func (x *CryptoGetAccountBalanceQuery) GetContractID() *ContractID {
	if x, ok := x.GetBalanceSource().(*CryptoGetAccountBalanceQuery_ContractID); ok {
		return x.ContractID
	}
	return nil
}

type isCryptoGetAccountBalanceQuery_BalanceSource interface {
	isCryptoGetAccountBalanceQuery_BalanceSource()
}

type CryptoGetAccountBalanceQuery_AccountID struct {
	// *
	// The account ID for which information is requested
	AccountID *AccountID `protobuf:"bytes,2,opt,name=accountID,proto3,oneof"`
}

type CryptoGetAccountBalanceQuery_ContractID struct {
	// *
	// The account ID for which information is requested
	ContractID *ContractID `protobuf:"bytes,3,opt,name=contractID,proto3,oneof"`
}

func (*CryptoGetAccountBalanceQuery_AccountID) isCryptoGetAccountBalanceQuery_BalanceSource() {}

func (*CryptoGetAccountBalanceQuery_ContractID) isCryptoGetAccountBalanceQuery_BalanceSource() {}

// *
// Response when the client sends the node CryptoGetAccountBalanceQuery
type CryptoGetAccountBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// Standard response from node to client, including the requested fields: cost, or state proof,
	// or both, or neither.
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// *
	// The account ID that is being described (this is useful with state proofs, for proving to a
	// third party)
	AccountID *AccountID `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
	// *
	// The current balance, in tinybars.
	Balance uint64 `protobuf:"varint,3,opt,name=balance,proto3" json:"balance,omitempty"`
	// *
	// [DEPRECATED] The balances of the tokens associated to the account. This field was
	// deprecated by <a href="https://hips.hedera.com/hip/hip-367">HIP-367</a>, which allowed
	// an account to be associated to an unlimited number of tokens. This scale makes it more
	// efficient for users to consult mirror nodes to review their token balances.
	//
	// Deprecated: Do not use.
	TokenBalances []*TokenBalance `protobuf:"bytes,4,rep,name=tokenBalances,proto3" json:"tokenBalances,omitempty"`
}

func (x *CryptoGetAccountBalanceResponse) Reset() {
	*x = CryptoGetAccountBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_get_account_balance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoGetAccountBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoGetAccountBalanceResponse) ProtoMessage() {}

func (x *CryptoGetAccountBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_get_account_balance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoGetAccountBalanceResponse.ProtoReflect.Descriptor instead.
func (*CryptoGetAccountBalanceResponse) Descriptor() ([]byte, []int) {
	return file_crypto_get_account_balance_proto_rawDescGZIP(), []int{1}
}

func (x *CryptoGetAccountBalanceResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *CryptoGetAccountBalanceResponse) GetAccountID() *AccountID {
	if x != nil {
		return x.AccountID
	}
	return nil
}

func (x *CryptoGetAccountBalanceResponse) GetBalance() uint64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

// Deprecated: Do not use.
func (x *CryptoGetAccountBalanceResponse) GetTokenBalances() []*TokenBalance {
	if x != nil {
		return x.TokenBalances
	}
	return nil
}

var File_crypto_get_account_balance_proto protoreflect.FileDescriptor

var file_crypto_get_account_balance_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x1c, 0x43, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x48, 0x00, 0x52, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x33, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x48, 0x00, 0x52,
	0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x42, 0x0f, 0x0a, 0x0d, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0xd9, 0x01, 0x0a,
	0x1f, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x2e, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x44, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12,
	0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0d, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e,
	0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crypto_get_account_balance_proto_rawDescOnce sync.Once
	file_crypto_get_account_balance_proto_rawDescData = file_crypto_get_account_balance_proto_rawDesc
)

func file_crypto_get_account_balance_proto_rawDescGZIP() []byte {
	file_crypto_get_account_balance_proto_rawDescOnce.Do(func() {
		file_crypto_get_account_balance_proto_rawDescData = protoimpl.X.CompressGZIP(file_crypto_get_account_balance_proto_rawDescData)
	})
	return file_crypto_get_account_balance_proto_rawDescData
}

var file_crypto_get_account_balance_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_crypto_get_account_balance_proto_goTypes = []interface{}{
	(*CryptoGetAccountBalanceQuery)(nil),    // 0: proto.CryptoGetAccountBalanceQuery
	(*CryptoGetAccountBalanceResponse)(nil), // 1: proto.CryptoGetAccountBalanceResponse
	(*QueryHeader)(nil),                     // 2: proto.QueryHeader
	(*AccountID)(nil),                       // 3: proto.AccountID
	(*ContractID)(nil),                      // 4: proto.ContractID
	(*ResponseHeader)(nil),                  // 5: proto.ResponseHeader
	(*TokenBalance)(nil),                    // 6: proto.TokenBalance
}
var file_crypto_get_account_balance_proto_depIdxs = []int32{
	2, // 0: proto.CryptoGetAccountBalanceQuery.header:type_name -> proto.QueryHeader
	3, // 1: proto.CryptoGetAccountBalanceQuery.accountID:type_name -> proto.AccountID
	4, // 2: proto.CryptoGetAccountBalanceQuery.contractID:type_name -> proto.ContractID
	5, // 3: proto.CryptoGetAccountBalanceResponse.header:type_name -> proto.ResponseHeader
	3, // 4: proto.CryptoGetAccountBalanceResponse.accountID:type_name -> proto.AccountID
	6, // 5: proto.CryptoGetAccountBalanceResponse.tokenBalances:type_name -> proto.TokenBalance
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_crypto_get_account_balance_proto_init() }
func file_crypto_get_account_balance_proto_init() {
	if File_crypto_get_account_balance_proto != nil {
		return
	}
	file_basic_types_proto_init()
	file_query_header_proto_init()
	file_response_header_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_crypto_get_account_balance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoGetAccountBalanceQuery); i {
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
		file_crypto_get_account_balance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoGetAccountBalanceResponse); i {
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
	file_crypto_get_account_balance_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CryptoGetAccountBalanceQuery_AccountID)(nil),
		(*CryptoGetAccountBalanceQuery_ContractID)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_crypto_get_account_balance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_crypto_get_account_balance_proto_goTypes,
		DependencyIndexes: file_crypto_get_account_balance_proto_depIdxs,
		MessageInfos:      file_crypto_get_account_balance_proto_msgTypes,
	}.Build()
	File_crypto_get_account_balance_proto = out.File
	file_crypto_get_account_balance_proto_rawDesc = nil
	file_crypto_get_account_balance_proto_goTypes = nil
	file_crypto_get_account_balance_proto_depIdxs = nil
}
