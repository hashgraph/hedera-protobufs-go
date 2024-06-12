// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: block_info.proto

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
// Information about ongoing, most recently completed, and last 256 blocks.
type BlockInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// The last block number, this is the last completed immutable block.
	LastBlockNumber int64 `protobuf:"varint,1,opt,name=last_block_number,json=lastBlockNumber,proto3" json:"last_block_number,omitempty"`
	// *
	// The consensus time of the first transaction of the last block, this is the last completed immutable block.
	FirstConsTimeOfLastBlock *Timestamp `protobuf:"bytes,2,opt,name=first_cons_time_of_last_block,json=firstConsTimeOfLastBlock,proto3" json:"first_cons_time_of_last_block,omitempty"`
	// *
	// SHA384 48 byte hashes of the last 256 blocks in single byte array.
	// First 48 bytes is the oldest block.
	// Last 48 bytes is the newest block, which is the last fully completed immutable block.
	// If we are shortly after genesis and there are less than 256 blocks then this could contain less than 256 hashes.
	BlockHashes []byte `protobuf:"bytes,3,opt,name=block_hashes,json=blockHashes,proto3" json:"block_hashes,omitempty"`
	// *
	// The consensus time of the last transaction that was handled by the node. This property is how we 'advance the
	// consensus clock', i.e. continually setting this property to the latest consensus timestamp (and thus transaction)
	// handled by the node.
	ConsTimeOfLastHandledTxn *Timestamp `protobuf:"bytes,4,opt,name=cons_time_of_last_handled_txn,json=consTimeOfLastHandledTxn,proto3" json:"cons_time_of_last_handled_txn,omitempty"`
	// *
	// A flag indicating whether migration records have been published. This property should be marked 'false'
	// immediately following a node upgrade, and marked 'true' once the migration records (if any) are published, which
	// should happen during the first transaction handled by the node.
	MigrationRecordsStreamed bool `protobuf:"varint,5,opt,name=migration_records_streamed,json=migrationRecordsStreamed,proto3" json:"migration_records_streamed,omitempty"`
	// *
	// The consensus time of the first transaction in the current block; necessary for reconnecting nodes to detect
	// when the current block is finished.
	FirstConsTimeOfCurrentBlock *Timestamp `protobuf:"bytes,6,opt,name=first_cons_time_of_current_block,json=firstConsTimeOfCurrentBlock,proto3" json:"first_cons_time_of_current_block,omitempty"`
}

func (x *BlockInfo) Reset() {
	*x = BlockInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockInfo) ProtoMessage() {}

func (x *BlockInfo) ProtoReflect() protoreflect.Message {
	mi := &file_block_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockInfo.ProtoReflect.Descriptor instead.
func (*BlockInfo) Descriptor() ([]byte, []int) {
	return file_block_info_proto_rawDescGZIP(), []int{0}
}

func (x *BlockInfo) GetLastBlockNumber() int64 {
	if x != nil {
		return x.LastBlockNumber
	}
	return 0
}

func (x *BlockInfo) GetFirstConsTimeOfLastBlock() *Timestamp {
	if x != nil {
		return x.FirstConsTimeOfLastBlock
	}
	return nil
}

func (x *BlockInfo) GetBlockHashes() []byte {
	if x != nil {
		return x.BlockHashes
	}
	return nil
}

func (x *BlockInfo) GetConsTimeOfLastHandledTxn() *Timestamp {
	if x != nil {
		return x.ConsTimeOfLastHandledTxn
	}
	return nil
}

func (x *BlockInfo) GetMigrationRecordsStreamed() bool {
	if x != nil {
		return x.MigrationRecordsStreamed
	}
	return false
}

func (x *BlockInfo) GetFirstConsTimeOfCurrentBlock() *Timestamp {
	if x != nil {
		return x.FirstConsTimeOfCurrentBlock
	}
	return nil
}

var File_block_info_proto protoreflect.FileDescriptor

var file_block_info_proto_rawDesc = []byte{
	0x0a, 0x10, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x03, 0x0a, 0x09, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x51, 0x0a, 0x1d, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x18, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x4c, 0x61,
	0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0x51, 0x0a, 0x1d, 0x63, 0x6f,
	0x6e, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x5f, 0x74, 0x78, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x18, 0x63, 0x6f, 0x6e, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x4c,
	0x61, 0x73, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x54, 0x78, 0x6e, 0x12, 0x3c, 0x0a,
	0x1a, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x18, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x64, 0x12, 0x57, 0x0a, 0x20, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6f,
	0x66, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x1b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x43, 0x6f,
	0x6e, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_block_info_proto_rawDescOnce sync.Once
	file_block_info_proto_rawDescData = file_block_info_proto_rawDesc
)

func file_block_info_proto_rawDescGZIP() []byte {
	file_block_info_proto_rawDescOnce.Do(func() {
		file_block_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_block_info_proto_rawDescData)
	})
	return file_block_info_proto_rawDescData
}

var file_block_info_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_block_info_proto_goTypes = []interface{}{
	(*BlockInfo)(nil), // 0: proto.BlockInfo
	(*Timestamp)(nil), // 1: proto.Timestamp
}
var file_block_info_proto_depIdxs = []int32{
	1, // 0: proto.BlockInfo.first_cons_time_of_last_block:type_name -> proto.Timestamp
	1, // 1: proto.BlockInfo.cons_time_of_last_handled_txn:type_name -> proto.Timestamp
	1, // 2: proto.BlockInfo.first_cons_time_of_current_block:type_name -> proto.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_block_info_proto_init() }
func file_block_info_proto_init() {
	if File_block_info_proto != nil {
		return
	}
	file_timestamp_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_block_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockInfo); i {
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
			RawDescriptor: file_block_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_block_info_proto_goTypes,
		DependencyIndexes: file_block_info_proto_depIdxs,
		MessageInfos:      file_block_info_proto_msgTypes,
	}.Build()
	File_block_info_proto = out.File
	file_block_info_proto_rawDesc = nil
	file_block_info_proto_goTypes = nil
	file_block_info_proto_depIdxs = nil
}
