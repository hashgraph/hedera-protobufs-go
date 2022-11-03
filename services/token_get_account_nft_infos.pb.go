// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.9
// source: token_get_account_nft_infos.proto

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
// Applicable only to tokens of type NON_FUNGIBLE_UNIQUE. Gets info on NFTs N through M owned by the
// specified accountId.
// Example: If Account A owns 5 NFTs (might be of different Token Entity), having start=0 and end=5
// will return all of the NFTs
//
// INVALID_QUERY_RANGE response code will be returned if:
//  1. Start > End
//  2. Start and End indices are non-positive
//  3. Start and End indices are out of boundaries for the retrieved nft list
//  4. The range between Start and End is bigger than the global dynamic property for maximum query
//     range
//
// NOT_SUPPORTED response code will be returned if the queried token is of type FUNGIBLE_COMMON
//
// INVALID_ACCOUNT_ID response code will be returned if the queried account does not exist
//
// ACCOUNT_DELETED response code will be returned if the queried account has been deleted
type TokenGetAccountNftInfosQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// Standard info sent from client to node, including the signed payment, and what kind of
	// response is requested (cost, state proof, both, or neither).
	Header *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// *
	// The Account for which information is requested
	AccountID *AccountID `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
	// *
	// Specifies the start index (inclusive) of the range of NFTs to query for. Value must be in the
	// range [0; ownedNFTs-1]
	Start int64 `protobuf:"varint,3,opt,name=start,proto3" json:"start,omitempty"`
	// *
	// Specifies the end index (exclusive) of the range of NFTs to query for. Value must be in the
	// range (start; ownedNFTs]
	End int64 `protobuf:"varint,4,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *TokenGetAccountNftInfosQuery) Reset() {
	*x = TokenGetAccountNftInfosQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_get_account_nft_infos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenGetAccountNftInfosQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenGetAccountNftInfosQuery) ProtoMessage() {}

func (x *TokenGetAccountNftInfosQuery) ProtoReflect() protoreflect.Message {
	mi := &file_token_get_account_nft_infos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenGetAccountNftInfosQuery.ProtoReflect.Descriptor instead.
func (*TokenGetAccountNftInfosQuery) Descriptor() ([]byte, []int) {
	return file_token_get_account_nft_infos_proto_rawDescGZIP(), []int{0}
}

func (x *TokenGetAccountNftInfosQuery) GetHeader() *QueryHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *TokenGetAccountNftInfosQuery) GetAccountID() *AccountID {
	if x != nil {
		return x.AccountID
	}
	return nil
}

func (x *TokenGetAccountNftInfosQuery) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *TokenGetAccountNftInfosQuery) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

// *
// UNDOCUMENTED
type TokenGetAccountNftInfosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// Standard response from node to client, including the requested fields: cost, or state proof,
	// or both, or neither
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// *
	// List of NFTs associated to the account
	Nfts []*TokenNftInfo `protobuf:"bytes,2,rep,name=nfts,proto3" json:"nfts,omitempty"`
}

func (x *TokenGetAccountNftInfosResponse) Reset() {
	*x = TokenGetAccountNftInfosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_get_account_nft_infos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenGetAccountNftInfosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenGetAccountNftInfosResponse) ProtoMessage() {}

func (x *TokenGetAccountNftInfosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_token_get_account_nft_infos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenGetAccountNftInfosResponse.ProtoReflect.Descriptor instead.
func (*TokenGetAccountNftInfosResponse) Descriptor() ([]byte, []int) {
	return file_token_get_account_nft_infos_proto_rawDescGZIP(), []int{1}
}

func (x *TokenGetAccountNftInfosResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *TokenGetAccountNftInfosResponse) GetNfts() []*TokenNftInfo {
	if x != nil {
		return x.Nfts
	}
	return nil
}

var File_token_get_account_nft_infos_proto protoreflect.FileDescriptor

var file_token_get_account_nft_infos_proto_rawDesc = []byte{
	0x0a, 0x21, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x66, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x6e, 0x66, 0x74, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x1c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x65, 0x74, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x2e, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x44, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x79, 0x0a, 0x1f, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x04, 0x6e, 0x66, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6e, 0x66,
	0x74, 0x73, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_token_get_account_nft_infos_proto_rawDescOnce sync.Once
	file_token_get_account_nft_infos_proto_rawDescData = file_token_get_account_nft_infos_proto_rawDesc
)

func file_token_get_account_nft_infos_proto_rawDescGZIP() []byte {
	file_token_get_account_nft_infos_proto_rawDescOnce.Do(func() {
		file_token_get_account_nft_infos_proto_rawDescData = protoimpl.X.CompressGZIP(file_token_get_account_nft_infos_proto_rawDescData)
	})
	return file_token_get_account_nft_infos_proto_rawDescData
}

var file_token_get_account_nft_infos_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_token_get_account_nft_infos_proto_goTypes = []interface{}{
	(*TokenGetAccountNftInfosQuery)(nil),    // 0: proto.TokenGetAccountNftInfosQuery
	(*TokenGetAccountNftInfosResponse)(nil), // 1: proto.TokenGetAccountNftInfosResponse
	(*QueryHeader)(nil),                     // 2: proto.QueryHeader
	(*AccountID)(nil),                       // 3: proto.AccountID
	(*ResponseHeader)(nil),                  // 4: proto.ResponseHeader
	(*TokenNftInfo)(nil),                    // 5: proto.TokenNftInfo
}
var file_token_get_account_nft_infos_proto_depIdxs = []int32{
	2, // 0: proto.TokenGetAccountNftInfosQuery.header:type_name -> proto.QueryHeader
	3, // 1: proto.TokenGetAccountNftInfosQuery.accountID:type_name -> proto.AccountID
	4, // 2: proto.TokenGetAccountNftInfosResponse.header:type_name -> proto.ResponseHeader
	5, // 3: proto.TokenGetAccountNftInfosResponse.nfts:type_name -> proto.TokenNftInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_token_get_account_nft_infos_proto_init() }
func file_token_get_account_nft_infos_proto_init() {
	if File_token_get_account_nft_infos_proto != nil {
		return
	}
	file_basic_types_proto_init()
	file_token_get_nft_info_proto_init()
	file_query_header_proto_init()
	file_response_header_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_token_get_account_nft_infos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenGetAccountNftInfosQuery); i {
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
		file_token_get_account_nft_infos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenGetAccountNftInfosResponse); i {
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
			RawDescriptor: file_token_get_account_nft_infos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_token_get_account_nft_infos_proto_goTypes,
		DependencyIndexes: file_token_get_account_nft_infos_proto_depIdxs,
		MessageInfos:      file_token_get_account_nft_infos_proto_msgTypes,
	}.Build()
	File_token_get_account_nft_infos_proto = out.File
	file_token_get_account_nft_infos_proto_rawDesc = nil
	file_token_get_account_nft_infos_proto_goTypes = nil
	file_token_get_account_nft_infos_proto_depIdxs = nil
}
