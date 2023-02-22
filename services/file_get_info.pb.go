// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: file_get_info.proto

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
// Get all of the information about a file, except for its contents. When a file expires, it no
// longer exists, and there will be no info about it, and the fileInfo field will be blank. If a
// transaction or smart contract deletes the file, but it has not yet expired, then the fileInfo
// field will be non-empty, the deleted field will be true, its size will be 0, and its contents
// will be empty.
type FileGetInfoQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// Standard info sent from client to node, including the signed payment, and what kind of
	// response is requested (cost, state proof, both, or neither).
	Header *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// *
	// The file ID of the file for which information is requested
	FileID *FileID `protobuf:"bytes,2,opt,name=fileID,proto3" json:"fileID,omitempty"`
}

func (x *FileGetInfoQuery) Reset() {
	*x = FileGetInfoQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_get_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileGetInfoQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileGetInfoQuery) ProtoMessage() {}

func (x *FileGetInfoQuery) ProtoReflect() protoreflect.Message {
	mi := &file_file_get_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileGetInfoQuery.ProtoReflect.Descriptor instead.
func (*FileGetInfoQuery) Descriptor() ([]byte, []int) {
	return file_file_get_info_proto_rawDescGZIP(), []int{0}
}

func (x *FileGetInfoQuery) GetHeader() *QueryHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *FileGetInfoQuery) GetFileID() *FileID {
	if x != nil {
		return x.FileID
	}
	return nil
}

// *
// Response when the client sends the node FileGetInfoQuery
type FileGetInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// Standard response from node to client, including the requested fields: cost, or state proof,
	// or both, or neither
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// *
	// The information about the file
	FileInfo *FileGetInfoResponse_FileInfo `protobuf:"bytes,2,opt,name=fileInfo,proto3" json:"fileInfo,omitempty"`
}

func (x *FileGetInfoResponse) Reset() {
	*x = FileGetInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_get_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileGetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileGetInfoResponse) ProtoMessage() {}

func (x *FileGetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_get_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileGetInfoResponse.ProtoReflect.Descriptor instead.
func (*FileGetInfoResponse) Descriptor() ([]byte, []int) {
	return file_file_get_info_proto_rawDescGZIP(), []int{1}
}

func (x *FileGetInfoResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *FileGetInfoResponse) GetFileInfo() *FileGetInfoResponse_FileInfo {
	if x != nil {
		return x.FileInfo
	}
	return nil
}

type FileGetInfoResponse_FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// The file ID of the file for which information is requested
	FileID *FileID `protobuf:"bytes,1,opt,name=fileID,proto3" json:"fileID,omitempty"`
	// *
	// Number of bytes in contents
	Size int64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	// *
	// The current time at which this account is set to expire
	ExpirationTime *Timestamp `protobuf:"bytes,3,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`
	// *
	// True if deleted but not yet expired
	Deleted bool `protobuf:"varint,4,opt,name=deleted,proto3" json:"deleted,omitempty"`
	// *
	// One of these keys must sign in order to modify or delete the file
	Keys *KeyList `protobuf:"bytes,5,opt,name=keys,proto3" json:"keys,omitempty"`
	// *
	// The memo associated with the file
	Memo string `protobuf:"bytes,6,opt,name=memo,proto3" json:"memo,omitempty"`
	// *
	// The ledger ID the response was returned from; please see <a href="https://github.com/hashgraph/hedera-improvement-proposal/blob/master/HIP/hip-198.md">HIP-198</a> for the network-specific IDs.
	LedgerId []byte `protobuf:"bytes,7,opt,name=ledger_id,json=ledgerId,proto3" json:"ledger_id,omitempty"`
}

func (x *FileGetInfoResponse_FileInfo) Reset() {
	*x = FileGetInfoResponse_FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_get_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileGetInfoResponse_FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileGetInfoResponse_FileInfo) ProtoMessage() {}

func (x *FileGetInfoResponse_FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_file_get_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileGetInfoResponse_FileInfo.ProtoReflect.Descriptor instead.
func (*FileGetInfoResponse_FileInfo) Descriptor() ([]byte, []int) {
	return file_file_get_info_proto_rawDescGZIP(), []int{1, 0}
}

func (x *FileGetInfoResponse_FileInfo) GetFileID() *FileID {
	if x != nil {
		return x.FileID
	}
	return nil
}

func (x *FileGetInfoResponse_FileInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileGetInfoResponse_FileInfo) GetExpirationTime() *Timestamp {
	if x != nil {
		return x.ExpirationTime
	}
	return nil
}

func (x *FileGetInfoResponse_FileInfo) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

func (x *FileGetInfoResponse_FileInfo) GetKeys() *KeyList {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *FileGetInfoResponse_FileInfo) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *FileGetInfoResponse_FileInfo) GetLedgerId() []byte {
	if x != nil {
		return x.LedgerId
	}
	return nil
}

var File_file_get_info_proto protoreflect.FileDescriptor

var file_file_get_info_proto_rawDesc = []byte{
	0x0a, 0x13, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x12, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x10, 0x46,
	0x69, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x06, 0x66,
	0x69, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x44, 0x22, 0xf6, 0x02, 0x0a, 0x13, 0x46, 0x69, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0xee, 0x01, 0x0a, 0x08, 0x46,
	0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x25, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x12, 0x38, 0x0a, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65,
	0x6d, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x1b,
	0x0a, 0x09, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x49, 0x64, 0x42, 0x26, 0x0a, 0x22, 0x63,
	0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76,
	0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_file_get_info_proto_rawDescOnce sync.Once
	file_file_get_info_proto_rawDescData = file_file_get_info_proto_rawDesc
)

func file_file_get_info_proto_rawDescGZIP() []byte {
	file_file_get_info_proto_rawDescOnce.Do(func() {
		file_file_get_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_get_info_proto_rawDescData)
	})
	return file_file_get_info_proto_rawDescData
}

var file_file_get_info_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_file_get_info_proto_goTypes = []interface{}{
	(*FileGetInfoQuery)(nil),             // 0: proto.FileGetInfoQuery
	(*FileGetInfoResponse)(nil),          // 1: proto.FileGetInfoResponse
	(*FileGetInfoResponse_FileInfo)(nil), // 2: proto.FileGetInfoResponse.FileInfo
	(*QueryHeader)(nil),                  // 3: proto.QueryHeader
	(*FileID)(nil),                       // 4: proto.FileID
	(*ResponseHeader)(nil),               // 5: proto.ResponseHeader
	(*Timestamp)(nil),                    // 6: proto.Timestamp
	(*KeyList)(nil),                      // 7: proto.KeyList
}
var file_file_get_info_proto_depIdxs = []int32{
	3, // 0: proto.FileGetInfoQuery.header:type_name -> proto.QueryHeader
	4, // 1: proto.FileGetInfoQuery.fileID:type_name -> proto.FileID
	5, // 2: proto.FileGetInfoResponse.header:type_name -> proto.ResponseHeader
	2, // 3: proto.FileGetInfoResponse.fileInfo:type_name -> proto.FileGetInfoResponse.FileInfo
	4, // 4: proto.FileGetInfoResponse.FileInfo.fileID:type_name -> proto.FileID
	6, // 5: proto.FileGetInfoResponse.FileInfo.expirationTime:type_name -> proto.Timestamp
	7, // 6: proto.FileGetInfoResponse.FileInfo.keys:type_name -> proto.KeyList
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_file_get_info_proto_init() }
func file_file_get_info_proto_init() {
	if File_file_get_info_proto != nil {
		return
	}
	file_timestamp_proto_init()
	file_basic_types_proto_init()
	file_query_header_proto_init()
	file_response_header_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_file_get_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileGetInfoQuery); i {
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
		file_file_get_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileGetInfoResponse); i {
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
		file_file_get_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileGetInfoResponse_FileInfo); i {
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
			RawDescriptor: file_file_get_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_file_get_info_proto_goTypes,
		DependencyIndexes: file_file_get_info_proto_depIdxs,
		MessageInfos:      file_file_get_info_proto_msgTypes,
	}.Build()
	File_file_get_info_proto = out.File
	file_file_get_info_proto_rawDesc = nil
	file_file_get_info_proto_goTypes = nil
	file_file_get_info_proto_depIdxs = nil
}
