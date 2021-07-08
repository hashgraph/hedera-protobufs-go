// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: GetByKey.proto

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

// Get all accounts, claims, files, and smart contract instances whose associated keys include the given Key. The given Key must not be a contractID or a ThresholdKey. This is not yet implemented in the API, but will be in the future.
type GetByKeyQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"` // Standard info sent from client to node, including the signed payment, and what kind of response is requested (cost, state proof, both, or neither).
	Key    *Key         `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`       // The key to search for. It must not contain a contractID nor a ThresholdSignature.
}

func (x *GetByKeyQuery) Reset() {
	*x = GetByKeyQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetByKey_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByKeyQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByKeyQuery) ProtoMessage() {}

func (x *GetByKeyQuery) ProtoReflect() protoreflect.Message {
	mi := &file_GetByKey_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByKeyQuery.ProtoReflect.Descriptor instead.
func (*GetByKeyQuery) Descriptor() ([]byte, []int) {
	return file_GetByKey_proto_rawDescGZIP(), []int{0}
}

func (x *GetByKeyQuery) GetHeader() *QueryHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *GetByKeyQuery) GetKey() *Key {
	if x != nil {
		return x.Key
	}
	return nil
}

// the ID for a single entity (account, livehash, file, or smart contract instance)
type EntityID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Entity:
	//	*EntityID_AccountID
	//	*EntityID_LiveHash
	//	*EntityID_FileID
	//	*EntityID_ContractID
	Entity isEntityID_Entity `protobuf_oneof:"entity"`
}

func (x *EntityID) Reset() {
	*x = EntityID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetByKey_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityID) ProtoMessage() {}

func (x *EntityID) ProtoReflect() protoreflect.Message {
	mi := &file_GetByKey_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityID.ProtoReflect.Descriptor instead.
func (*EntityID) Descriptor() ([]byte, []int) {
	return file_GetByKey_proto_rawDescGZIP(), []int{1}
}

func (m *EntityID) GetEntity() isEntityID_Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (x *EntityID) GetAccountID() *AccountID {
	if x, ok := x.GetEntity().(*EntityID_AccountID); ok {
		return x.AccountID
	}
	return nil
}

func (x *EntityID) GetLiveHash() *LiveHash {
	if x, ok := x.GetEntity().(*EntityID_LiveHash); ok {
		return x.LiveHash
	}
	return nil
}

func (x *EntityID) GetFileID() *FileID {
	if x, ok := x.GetEntity().(*EntityID_FileID); ok {
		return x.FileID
	}
	return nil
}

func (x *EntityID) GetContractID() *ContractID {
	if x, ok := x.GetEntity().(*EntityID_ContractID); ok {
		return x.ContractID
	}
	return nil
}

type isEntityID_Entity interface {
	isEntityID_Entity()
}

type EntityID_AccountID struct {
	AccountID *AccountID `protobuf:"bytes,1,opt,name=accountID,proto3,oneof"` // The Account ID for the cryptocurrency account
}

type EntityID_LiveHash struct {
	LiveHash *LiveHash `protobuf:"bytes,2,opt,name=liveHash,proto3,oneof"` // A uniquely identifying livehash of an acount
}

type EntityID_FileID struct {
	FileID *FileID `protobuf:"bytes,3,opt,name=fileID,proto3,oneof"` // The file ID of the file
}

type EntityID_ContractID struct {
	ContractID *ContractID `protobuf:"bytes,4,opt,name=contractID,proto3,oneof"` // The smart contract ID that identifies instance
}

func (*EntityID_AccountID) isEntityID_Entity() {}

func (*EntityID_LiveHash) isEntityID_Entity() {}

func (*EntityID_FileID) isEntityID_Entity() {}

func (*EntityID_ContractID) isEntityID_Entity() {}

// Response when the client sends the node GetByKeyQuery
type GetByKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header   *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`     //Standard response from node to client, including the requested fields: cost, or state proof, or both, or neither
	Entities []*EntityID     `protobuf:"bytes,2,rep,name=entities,proto3" json:"entities,omitempty"` // The list of entities that include this public key in their associated Key list
}

func (x *GetByKeyResponse) Reset() {
	*x = GetByKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetByKey_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByKeyResponse) ProtoMessage() {}

func (x *GetByKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_GetByKey_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByKeyResponse.ProtoReflect.Descriptor instead.
func (*GetByKeyResponse) Descriptor() ([]byte, []int) {
	return file_GetByKey_proto_rawDescGZIP(), []int{2}
}

func (x *GetByKeyResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *GetByKeyResponse) GetEntities() []*EntityID {
	if x != nil {
		return x.Entities
	}
	return nil
}

var File_GetByKey_proto protoreflect.FileDescriptor

var file_GetByKey_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x76,
	0x65, 0x48, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65,
	0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0xd3, 0x01, 0x0a, 0x08, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x48, 0x00, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x2d, 0x0a, 0x08, 0x6c, 0x69, 0x76, 0x65, 0x48, 0x61, 0x73,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4c, 0x69, 0x76, 0x65, 0x48, 0x61, 0x73, 0x68, 0x48, 0x00, 0x52, 0x08, 0x6c, 0x69, 0x76, 0x65,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x27, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x49, 0x44, 0x48, 0x00, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x33, 0x0a,
	0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x49, 0x44, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x49, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x6e, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x2b, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x44, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x26, 0x0a, 0x22,
	0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61,
	0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetByKey_proto_rawDescOnce sync.Once
	file_GetByKey_proto_rawDescData = file_GetByKey_proto_rawDesc
)

func file_GetByKey_proto_rawDescGZIP() []byte {
	file_GetByKey_proto_rawDescOnce.Do(func() {
		file_GetByKey_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetByKey_proto_rawDescData)
	})
	return file_GetByKey_proto_rawDescData
}

var file_GetByKey_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_GetByKey_proto_goTypes = []interface{}{
	(*GetByKeyQuery)(nil),    // 0: proto.GetByKeyQuery
	(*EntityID)(nil),         // 1: proto.EntityID
	(*GetByKeyResponse)(nil), // 2: proto.GetByKeyResponse
	(*QueryHeader)(nil),      // 3: proto.QueryHeader
	(*Key)(nil),              // 4: proto.Key
	(*AccountID)(nil),        // 5: proto.AccountID
	(*LiveHash)(nil),         // 6: proto.LiveHash
	(*FileID)(nil),           // 7: proto.FileID
	(*ContractID)(nil),       // 8: proto.ContractID
	(*ResponseHeader)(nil),   // 9: proto.ResponseHeader
}
var file_GetByKey_proto_depIdxs = []int32{
	3, // 0: proto.GetByKeyQuery.header:type_name -> proto.QueryHeader
	4, // 1: proto.GetByKeyQuery.key:type_name -> proto.Key
	5, // 2: proto.EntityID.accountID:type_name -> proto.AccountID
	6, // 3: proto.EntityID.liveHash:type_name -> proto.LiveHash
	7, // 4: proto.EntityID.fileID:type_name -> proto.FileID
	8, // 5: proto.EntityID.contractID:type_name -> proto.ContractID
	9, // 6: proto.GetByKeyResponse.header:type_name -> proto.ResponseHeader
	1, // 7: proto.GetByKeyResponse.entities:type_name -> proto.EntityID
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_GetByKey_proto_init() }
func file_GetByKey_proto_init() {
	if File_GetByKey_proto != nil {
		return
	}
	file_BasicTypes_proto_init()
	file_QueryHeader_proto_init()
	file_ResponseHeader_proto_init()
	file_CryptoAddLiveHash_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetByKey_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByKeyQuery); i {
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
		file_GetByKey_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityID); i {
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
		file_GetByKey_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByKeyResponse); i {
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
	file_GetByKey_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*EntityID_AccountID)(nil),
		(*EntityID_LiveHash)(nil),
		(*EntityID_FileID)(nil),
		(*EntityID_ContractID)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_GetByKey_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetByKey_proto_goTypes,
		DependencyIndexes: file_GetByKey_proto_depIdxs,
		MessageInfos:      file_GetByKey_proto_msgTypes,
	}.Build()
	File_GetByKey_proto = out.File
	file_GetByKey_proto_rawDesc = nil
	file_GetByKey_proto_goTypes = nil
	file_GetByKey_proto_depIdxs = nil
}
