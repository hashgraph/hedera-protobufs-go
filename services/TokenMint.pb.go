// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: TokenMint.proto

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

//
//Mints tokens to the Token's treasury Account. If no Supply Key is defined, the transaction will resolve to TOKEN_HAS_NO_SUPPLY_KEY.
//The operation increases the Total Supply of the Token. The maximum total supply a token can have is 2^63-1.
//The amount provided must be in the lowest denomination possible. Example:
//Token A has 2 decimals. In order to mint 100 tokens, one must provide amount of 10000. In order to mint 100.55 tokens, one must provide amount of 10055.
type TokenMintTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    *TokenID `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`       // The token for which to mint tokens. If token does not exist, transaction results in INVALID_TOKEN_ID
	Amount   uint64   `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`    // Applicable to tokens of type FUNGIBLE_COMMON. The amount to mint to the Treasury Account. Amount must be a positive non-zero number represented in the lowest denomination of the token. The new supply must be lower than 2^63.
	Metadata [][]byte `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty"` // Applicable to tokens of type NON_FUNGIBLE_UNIQUE. A list of metadata that are being created. Maximum allowed size of each metadata is 100 bytes
}

func (x *TokenMintTransactionBody) Reset() {
	*x = TokenMintTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TokenMint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenMintTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenMintTransactionBody) ProtoMessage() {}

func (x *TokenMintTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_TokenMint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenMintTransactionBody.ProtoReflect.Descriptor instead.
func (*TokenMintTransactionBody) Descriptor() ([]byte, []int) {
	return file_TokenMint_proto_rawDescGZIP(), []int{0}
}

func (x *TokenMintTransactionBody) GetToken() *TokenID {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *TokenMintTransactionBody) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TokenMintTransactionBody) GetMetadata() [][]byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_TokenMint_proto protoreflect.FileDescriptor

var file_TokenMint_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x18, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x4d, 0x69, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x24, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x49, 0x44, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61,
	0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TokenMint_proto_rawDescOnce sync.Once
	file_TokenMint_proto_rawDescData = file_TokenMint_proto_rawDesc
)

func file_TokenMint_proto_rawDescGZIP() []byte {
	file_TokenMint_proto_rawDescOnce.Do(func() {
		file_TokenMint_proto_rawDescData = protoimpl.X.CompressGZIP(file_TokenMint_proto_rawDescData)
	})
	return file_TokenMint_proto_rawDescData
}

var file_TokenMint_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TokenMint_proto_goTypes = []interface{}{
	(*TokenMintTransactionBody)(nil), // 0: proto.TokenMintTransactionBody
	(*TokenID)(nil),                  // 1: proto.TokenID
}
var file_TokenMint_proto_depIdxs = []int32{
	1, // 0: proto.TokenMintTransactionBody.token:type_name -> proto.TokenID
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_TokenMint_proto_init() }
func file_TokenMint_proto_init() {
	if File_TokenMint_proto != nil {
		return
	}
	file_BasicTypes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TokenMint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenMintTransactionBody); i {
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
			RawDescriptor: file_TokenMint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TokenMint_proto_goTypes,
		DependencyIndexes: file_TokenMint_proto_depIdxs,
		MessageInfos:      file_TokenMint_proto_msgTypes,
	}.Build()
	File_TokenMint_proto = out.File
	file_TokenMint_proto_rawDesc = nil
	file_TokenMint_proto_goTypes = nil
	file_TokenMint_proto_depIdxs = nil
}
