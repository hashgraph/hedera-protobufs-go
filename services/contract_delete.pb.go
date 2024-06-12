// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: contract_delete.proto

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
// At consensus, marks a contract as deleted and transfers its remaining hBars, if any, to a
// designated receiver. After a contract is deleted, it can no longer be called.
//
// If the target contract is immutable (that is, was created without an admin key), then this
// transaction resolves to MODIFYING_IMMUTABLE_CONTRACT.
//
// --- Signing Requirements ---
// 1. The admin key of the target contract must sign.
// 2. If the transfer account or contract has receiverSigRequired, its associated key must also sign
type ContractDeleteTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// The id of the contract to be deleted
	ContractID *ContractID `protobuf:"bytes,1,opt,name=contractID,proto3" json:"contractID,omitempty"`
	// Types that are assignable to Obtainers:
	//
	//	*ContractDeleteTransactionBody_TransferAccountID
	//	*ContractDeleteTransactionBody_TransferContractID
	Obtainers isContractDeleteTransactionBody_Obtainers `protobuf_oneof:"obtainers"`
	// *
	// If set to true, means this is a "synthetic" system transaction being used to
	// alert mirror nodes that the contract is being permanently removed from the ledger.
	// <b>IMPORTANT:</b> User transactions cannot set this field to true, as permanent
	// removal is always managed by the ledger itself. Any ContractDeleteTransactionBody
	// submitted to HAPI with permanent_removal=true will be rejected with precheck status
	// PERMANENT_REMOVAL_REQUIRES_SYSTEM_INITIATION.
	PermanentRemoval bool `protobuf:"varint,4,opt,name=permanent_removal,json=permanentRemoval,proto3" json:"permanent_removal,omitempty"`
}

func (x *ContractDeleteTransactionBody) Reset() {
	*x = ContractDeleteTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_delete_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractDeleteTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractDeleteTransactionBody) ProtoMessage() {}

func (x *ContractDeleteTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_contract_delete_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractDeleteTransactionBody.ProtoReflect.Descriptor instead.
func (*ContractDeleteTransactionBody) Descriptor() ([]byte, []int) {
	return file_contract_delete_proto_rawDescGZIP(), []int{0}
}

func (x *ContractDeleteTransactionBody) GetContractID() *ContractID {
	if x != nil {
		return x.ContractID
	}
	return nil
}

func (m *ContractDeleteTransactionBody) GetObtainers() isContractDeleteTransactionBody_Obtainers {
	if m != nil {
		return m.Obtainers
	}
	return nil
}

func (x *ContractDeleteTransactionBody) GetTransferAccountID() *AccountID {
	if x, ok := x.GetObtainers().(*ContractDeleteTransactionBody_TransferAccountID); ok {
		return x.TransferAccountID
	}
	return nil
}

func (x *ContractDeleteTransactionBody) GetTransferContractID() *ContractID {
	if x, ok := x.GetObtainers().(*ContractDeleteTransactionBody_TransferContractID); ok {
		return x.TransferContractID
	}
	return nil
}

func (x *ContractDeleteTransactionBody) GetPermanentRemoval() bool {
	if x != nil {
		return x.PermanentRemoval
	}
	return false
}

type isContractDeleteTransactionBody_Obtainers interface {
	isContractDeleteTransactionBody_Obtainers()
}

type ContractDeleteTransactionBody_TransferAccountID struct {
	// *
	// The id of an account to receive any remaining hBars from the deleted contract
	TransferAccountID *AccountID `protobuf:"bytes,2,opt,name=transferAccountID,proto3,oneof"`
}

type ContractDeleteTransactionBody_TransferContractID struct {
	// *
	// The id of a contract to receive any remaining hBars from the deleted contract
	TransferContractID *ContractID `protobuf:"bytes,3,opt,name=transferContractID,proto3,oneof"`
}

func (*ContractDeleteTransactionBody_TransferAccountID) isContractDeleteTransactionBody_Obtainers() {}

func (*ContractDeleteTransactionBody_TransferContractID) isContractDeleteTransactionBody_Obtainers() {
}

var File_contract_delete_proto protoreflect.FileDescriptor

var file_contract_delete_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x93, 0x02, 0x0a, 0x1d, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x6f, 0x64, 0x79, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x12, 0x40, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x44, 0x48, 0x00, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x43, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x48, 0x00, 0x52, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x12, 0x2b, 0x0a,
	0x11, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76,
	0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6e,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x61, 0x6c, 0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x62,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contract_delete_proto_rawDescOnce sync.Once
	file_contract_delete_proto_rawDescData = file_contract_delete_proto_rawDesc
)

func file_contract_delete_proto_rawDescGZIP() []byte {
	file_contract_delete_proto_rawDescOnce.Do(func() {
		file_contract_delete_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_delete_proto_rawDescData)
	})
	return file_contract_delete_proto_rawDescData
}

var file_contract_delete_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_contract_delete_proto_goTypes = []interface{}{
	(*ContractDeleteTransactionBody)(nil), // 0: proto.ContractDeleteTransactionBody
	(*ContractID)(nil),                    // 1: proto.ContractID
	(*AccountID)(nil),                     // 2: proto.AccountID
}
var file_contract_delete_proto_depIdxs = []int32{
	1, // 0: proto.ContractDeleteTransactionBody.contractID:type_name -> proto.ContractID
	2, // 1: proto.ContractDeleteTransactionBody.transferAccountID:type_name -> proto.AccountID
	1, // 2: proto.ContractDeleteTransactionBody.transferContractID:type_name -> proto.ContractID
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_contract_delete_proto_init() }
func file_contract_delete_proto_init() {
	if File_contract_delete_proto != nil {
		return
	}
	file_basic_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_contract_delete_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractDeleteTransactionBody); i {
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
	file_contract_delete_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ContractDeleteTransactionBody_TransferAccountID)(nil),
		(*ContractDeleteTransactionBody_TransferContractID)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_contract_delete_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contract_delete_proto_goTypes,
		DependencyIndexes: file_contract_delete_proto_depIdxs,
		MessageInfos:      file_contract_delete_proto_msgTypes,
	}.Build()
	File_contract_delete_proto = out.File
	file_contract_delete_proto_rawDesc = nil
	file_contract_delete_proto_goTypes = nil
	file_contract_delete_proto_depIdxs = nil
}
