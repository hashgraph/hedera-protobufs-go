// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.9
// source: transaction_get_receipt.proto

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
// Get the receipt of a transaction, given its transaction ID. Once a transaction reaches consensus,
// then information about whether it succeeded or failed will be available until the end of the
// receipt period.  Before and after the receipt period, and for a transaction that was never
// submitted, the receipt is unknown.  This query is free (the payment field is left empty). No
// State proof is available for this response
type TransactionGetReceiptQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// Standard info sent from client to node, including the signed payment, and what kind of
	// response is requested (cost, state proof, both, or neither).
	Header *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// *
	// The ID of the transaction for which the receipt is requested.
	TransactionID *TransactionID `protobuf:"bytes,2,opt,name=transactionID,proto3" json:"transactionID,omitempty"`
	// *
	// Whether receipts of processing duplicate transactions should be returned along with the
	// receipt of processing the first consensus transaction with the given id whose status was
	// neither <tt>INVALID_NODE_ACCOUNT</tt> nor <tt>INVALID_PAYER_SIGNATURE</tt>; <b>or</b>, if no
	// such receipt exists, the receipt of processing the first transaction to reach consensus with
	// the given transaction id.
	IncludeDuplicates bool `protobuf:"varint,3,opt,name=includeDuplicates,proto3" json:"includeDuplicates,omitempty"`
	// *
	// Whether the response should include the receipts of any child transactions spawned by the
	// top-level transaction with the given transactionID.
	IncludeChildReceipts bool `protobuf:"varint,4,opt,name=include_child_receipts,json=includeChildReceipts,proto3" json:"include_child_receipts,omitempty"`
}

func (x *TransactionGetReceiptQuery) Reset() {
	*x = TransactionGetReceiptQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_get_receipt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionGetReceiptQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionGetReceiptQuery) ProtoMessage() {}

func (x *TransactionGetReceiptQuery) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_get_receipt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionGetReceiptQuery.ProtoReflect.Descriptor instead.
func (*TransactionGetReceiptQuery) Descriptor() ([]byte, []int) {
	return file_transaction_get_receipt_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionGetReceiptQuery) GetHeader() *QueryHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *TransactionGetReceiptQuery) GetTransactionID() *TransactionID {
	if x != nil {
		return x.TransactionID
	}
	return nil
}

func (x *TransactionGetReceiptQuery) GetIncludeDuplicates() bool {
	if x != nil {
		return x.IncludeDuplicates
	}
	return false
}

func (x *TransactionGetReceiptQuery) GetIncludeChildReceipts() bool {
	if x != nil {
		return x.IncludeChildReceipts
	}
	return false
}

// *
// Response when the client sends the node TransactionGetReceiptQuery. If it created a new entity
// (account, file, or smart contract instance) then one of the three ID fields will be filled in
// with the ID of the new entity. Sometimes a single transaction will create more than one new
// entity, such as when a new contract instance is created, and this also creates the new account
// that it owned by that instance. No State proof is available for this response
type TransactionGetReceiptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// Standard response from node to client, including the requested fields: cost, or state proof,
	// or both, or neither
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// *
	// Either the receipt of processing the first consensus transaction with the given id whose
	// status was neither <tt>INVALID_NODE_ACCOUNT</tt> nor <tt>INVALID_PAYER_SIGNATURE</tt>;
	// <b>or</b>, if no such receipt exists, the receipt of processing the first transaction to
	// reach consensus with the given transaction id.
	Receipt *TransactionReceipt `protobuf:"bytes,2,opt,name=receipt,proto3" json:"receipt,omitempty"`
	// *
	// The receipts of processing all transactions with the given id, in consensus time order.
	DuplicateTransactionReceipts []*TransactionReceipt `protobuf:"bytes,4,rep,name=duplicateTransactionReceipts,proto3" json:"duplicateTransactionReceipts,omitempty"`
	// *
	// The receipts (if any) of all child transactions spawned by the transaction with the
	// given top-level id, in consensus order. Always empty if the top-level status is UNKNOWN.
	ChildTransactionReceipts []*TransactionReceipt `protobuf:"bytes,5,rep,name=child_transaction_receipts,json=childTransactionReceipts,proto3" json:"child_transaction_receipts,omitempty"`
}

func (x *TransactionGetReceiptResponse) Reset() {
	*x = TransactionGetReceiptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_get_receipt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionGetReceiptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionGetReceiptResponse) ProtoMessage() {}

func (x *TransactionGetReceiptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_get_receipt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionGetReceiptResponse.ProtoReflect.Descriptor instead.
func (*TransactionGetReceiptResponse) Descriptor() ([]byte, []int) {
	return file_transaction_get_receipt_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionGetReceiptResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *TransactionGetReceiptResponse) GetReceipt() *TransactionReceipt {
	if x != nil {
		return x.Receipt
	}
	return nil
}

func (x *TransactionGetReceiptResponse) GetDuplicateTransactionReceipts() []*TransactionReceipt {
	if x != nil {
		return x.DuplicateTransactionReceipts
	}
	return nil
}

func (x *TransactionGetReceiptResponse) GetChildTransactionReceipts() []*TransactionReceipt {
	if x != nil {
		return x.ChildTransactionReceipts
	}
	return nil
}

var File_transaction_get_receipt_proto protoreflect.FileDescriptor

var file_transaction_get_receipt_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x65,
	0x74, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe8, 0x01, 0x0a, 0x1a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2a,
	0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x0d, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x11, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x11, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f,
	0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x43, 0x68, 0x69,
	0x6c, 0x64, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x22, 0xbb, 0x02, 0x0a, 0x1d, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74,
	0x12, 0x5d, 0x0a, 0x1c, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70,
	0x74, 0x52, 0x1c, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x12,
	0x57, 0x0a, 0x1a, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x18,
	0x63, 0x68, 0x69, 0x6c, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e,
	0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transaction_get_receipt_proto_rawDescOnce sync.Once
	file_transaction_get_receipt_proto_rawDescData = file_transaction_get_receipt_proto_rawDesc
)

func file_transaction_get_receipt_proto_rawDescGZIP() []byte {
	file_transaction_get_receipt_proto_rawDescOnce.Do(func() {
		file_transaction_get_receipt_proto_rawDescData = protoimpl.X.CompressGZIP(file_transaction_get_receipt_proto_rawDescData)
	})
	return file_transaction_get_receipt_proto_rawDescData
}

var file_transaction_get_receipt_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_transaction_get_receipt_proto_goTypes = []interface{}{
	(*TransactionGetReceiptQuery)(nil),    // 0: proto.TransactionGetReceiptQuery
	(*TransactionGetReceiptResponse)(nil), // 1: proto.TransactionGetReceiptResponse
	(*QueryHeader)(nil),                   // 2: proto.QueryHeader
	(*TransactionID)(nil),                 // 3: proto.TransactionID
	(*ResponseHeader)(nil),                // 4: proto.ResponseHeader
	(*TransactionReceipt)(nil),            // 5: proto.TransactionReceipt
}
var file_transaction_get_receipt_proto_depIdxs = []int32{
	2, // 0: proto.TransactionGetReceiptQuery.header:type_name -> proto.QueryHeader
	3, // 1: proto.TransactionGetReceiptQuery.transactionID:type_name -> proto.TransactionID
	4, // 2: proto.TransactionGetReceiptResponse.header:type_name -> proto.ResponseHeader
	5, // 3: proto.TransactionGetReceiptResponse.receipt:type_name -> proto.TransactionReceipt
	5, // 4: proto.TransactionGetReceiptResponse.duplicateTransactionReceipts:type_name -> proto.TransactionReceipt
	5, // 5: proto.TransactionGetReceiptResponse.child_transaction_receipts:type_name -> proto.TransactionReceipt
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_transaction_get_receipt_proto_init() }
func file_transaction_get_receipt_proto_init() {
	if File_transaction_get_receipt_proto != nil {
		return
	}
	file_transaction_receipt_proto_init()
	file_basic_types_proto_init()
	file_query_header_proto_init()
	file_response_header_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_transaction_get_receipt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionGetReceiptQuery); i {
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
		file_transaction_get_receipt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionGetReceiptResponse); i {
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
			RawDescriptor: file_transaction_get_receipt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transaction_get_receipt_proto_goTypes,
		DependencyIndexes: file_transaction_get_receipt_proto_depIdxs,
		MessageInfos:      file_transaction_get_receipt_proto_msgTypes,
	}.Build()
	File_transaction_get_receipt_proto = out.File
	file_transaction_get_receipt_proto_rawDesc = nil
	file_transaction_get_receipt_proto_goTypes = nil
	file_transaction_get_receipt_proto_depIdxs = nil
}
