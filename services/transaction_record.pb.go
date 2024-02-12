// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: transaction_record.proto

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
// Response when the client sends the node TransactionGetRecordResponse
type TransactionRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// The status (reach consensus, or failed, or is unknown) and the ID of any new
	// account/file/instance created.
	Receipt *TransactionReceipt `protobuf:"bytes,1,opt,name=receipt,proto3" json:"receipt,omitempty"`
	// *
	// The hash of the Transaction that executed (not the hash of any Transaction that failed for
	// having a duplicate TransactionID)
	TransactionHash []byte `protobuf:"bytes,2,opt,name=transactionHash,proto3" json:"transactionHash,omitempty"`
	// *
	// The consensus timestamp (or null if didn't reach consensus yet)
	ConsensusTimestamp *Timestamp `protobuf:"bytes,3,opt,name=consensusTimestamp,proto3" json:"consensusTimestamp,omitempty"`
	// *
	// The ID of the transaction this record represents
	TransactionID *TransactionID `protobuf:"bytes,4,opt,name=transactionID,proto3" json:"transactionID,omitempty"`
	// *
	// The memo that was submitted as part of the transaction (max 100 bytes)
	Memo string `protobuf:"bytes,5,opt,name=memo,proto3" json:"memo,omitempty"`
	// *
	// The actual transaction fee charged, not the original transactionFee value from
	// TransactionBody
	TransactionFee uint64 `protobuf:"varint,6,opt,name=transactionFee,proto3" json:"transactionFee,omitempty"`
	// Types that are assignable to Body:
	//
	//	*TransactionRecord_ContractCallResult
	//	*TransactionRecord_ContractCreateResult
	Body isTransactionRecord_Body `protobuf_oneof:"body"`
	// *
	// All hbar transfers as a result of this transaction, such as fees, or transfers performed by
	// the transaction, or by a smart contract it calls, or by the creation of threshold records
	// that it triggers.
	TransferList *TransferList `protobuf:"bytes,10,opt,name=transferList,proto3" json:"transferList,omitempty"`
	// *
	// All Token transfers as a result of this transaction
	TokenTransferLists []*TokenTransferList `protobuf:"bytes,11,rep,name=tokenTransferLists,proto3" json:"tokenTransferLists,omitempty"`
	// *
	// Reference to the scheduled transaction ID that this transaction record represent
	ScheduleRef *ScheduleID `protobuf:"bytes,12,opt,name=scheduleRef,proto3" json:"scheduleRef,omitempty"`
	// *
	// All custom fees that were assessed during a CryptoTransfer, and must be paid if the
	// transaction status resolved to SUCCESS
	AssessedCustomFees []*AssessedCustomFee `protobuf:"bytes,13,rep,name=assessed_custom_fees,json=assessedCustomFees,proto3" json:"assessed_custom_fees,omitempty"`
	// *
	// All token associations implicitly created while handling this transaction
	AutomaticTokenAssociations []*TokenAssociation `protobuf:"bytes,14,rep,name=automatic_token_associations,json=automaticTokenAssociations,proto3" json:"automatic_token_associations,omitempty"`
	// *
	// In the record of an internal transaction, the consensus timestamp of the user
	// transaction that spawned it.
	ParentConsensusTimestamp *Timestamp `protobuf:"bytes,15,opt,name=parent_consensus_timestamp,json=parentConsensusTimestamp,proto3" json:"parent_consensus_timestamp,omitempty"`
	// *
	// In the record of a CryptoCreate transaction triggered by a user transaction with a
	// (previously unused) alias, the new account's alias.
	Alias []byte `protobuf:"bytes,16,opt,name=alias,proto3" json:"alias,omitempty"`
	// *
	// The keccak256 hash of the ethereumData. This field will only be populated for
	// EthereumTransaction.
	EthereumHash []byte `protobuf:"bytes,17,opt,name=ethereum_hash,json=ethereumHash,proto3" json:"ethereum_hash,omitempty"`
	// *
	// List of accounts with the corresponding staking rewards paid as a result of a transaction.
	PaidStakingRewards []*AccountAmount `protobuf:"bytes,18,rep,name=paid_staking_rewards,json=paidStakingRewards,proto3" json:"paid_staking_rewards,omitempty"`
	// Types that are assignable to Entropy:
	//
	//	*TransactionRecord_PrngBytes
	//	*TransactionRecord_PrngNumber
	Entropy isTransactionRecord_Entropy `protobuf_oneof:"entropy"`
	// *
	// The new default EVM address of the account created by this transaction.
	// This field is populated only when the EVM address is not specified in the related transaction body.
	EvmAddress []byte `protobuf:"bytes,21,opt,name=evm_address,json=evmAddress,proto3" json:"evm_address,omitempty"`
}

func (x *TransactionRecord) Reset() {
	*x = TransactionRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionRecord) ProtoMessage() {}

func (x *TransactionRecord) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionRecord.ProtoReflect.Descriptor instead.
func (*TransactionRecord) Descriptor() ([]byte, []int) {
	return file_transaction_record_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionRecord) GetReceipt() *TransactionReceipt {
	if x != nil {
		return x.Receipt
	}
	return nil
}

func (x *TransactionRecord) GetTransactionHash() []byte {
	if x != nil {
		return x.TransactionHash
	}
	return nil
}

func (x *TransactionRecord) GetConsensusTimestamp() *Timestamp {
	if x != nil {
		return x.ConsensusTimestamp
	}
	return nil
}

func (x *TransactionRecord) GetTransactionID() *TransactionID {
	if x != nil {
		return x.TransactionID
	}
	return nil
}

func (x *TransactionRecord) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *TransactionRecord) GetTransactionFee() uint64 {
	if x != nil {
		return x.TransactionFee
	}
	return 0
}

func (m *TransactionRecord) GetBody() isTransactionRecord_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *TransactionRecord) GetContractCallResult() *ContractFunctionResult {
	if x, ok := x.GetBody().(*TransactionRecord_ContractCallResult); ok {
		return x.ContractCallResult
	}
	return nil
}

func (x *TransactionRecord) GetContractCreateResult() *ContractFunctionResult {
	if x, ok := x.GetBody().(*TransactionRecord_ContractCreateResult); ok {
		return x.ContractCreateResult
	}
	return nil
}

func (x *TransactionRecord) GetTransferList() *TransferList {
	if x != nil {
		return x.TransferList
	}
	return nil
}

func (x *TransactionRecord) GetTokenTransferLists() []*TokenTransferList {
	if x != nil {
		return x.TokenTransferLists
	}
	return nil
}

func (x *TransactionRecord) GetScheduleRef() *ScheduleID {
	if x != nil {
		return x.ScheduleRef
	}
	return nil
}

func (x *TransactionRecord) GetAssessedCustomFees() []*AssessedCustomFee {
	if x != nil {
		return x.AssessedCustomFees
	}
	return nil
}

func (x *TransactionRecord) GetAutomaticTokenAssociations() []*TokenAssociation {
	if x != nil {
		return x.AutomaticTokenAssociations
	}
	return nil
}

func (x *TransactionRecord) GetParentConsensusTimestamp() *Timestamp {
	if x != nil {
		return x.ParentConsensusTimestamp
	}
	return nil
}

func (x *TransactionRecord) GetAlias() []byte {
	if x != nil {
		return x.Alias
	}
	return nil
}

func (x *TransactionRecord) GetEthereumHash() []byte {
	if x != nil {
		return x.EthereumHash
	}
	return nil
}

func (x *TransactionRecord) GetPaidStakingRewards() []*AccountAmount {
	if x != nil {
		return x.PaidStakingRewards
	}
	return nil
}

func (m *TransactionRecord) GetEntropy() isTransactionRecord_Entropy {
	if m != nil {
		return m.Entropy
	}
	return nil
}

func (x *TransactionRecord) GetPrngBytes() []byte {
	if x, ok := x.GetEntropy().(*TransactionRecord_PrngBytes); ok {
		return x.PrngBytes
	}
	return nil
}

func (x *TransactionRecord) GetPrngNumber() int32 {
	if x, ok := x.GetEntropy().(*TransactionRecord_PrngNumber); ok {
		return x.PrngNumber
	}
	return 0
}

func (x *TransactionRecord) GetEvmAddress() []byte {
	if x != nil {
		return x.EvmAddress
	}
	return nil
}

type isTransactionRecord_Body interface {
	isTransactionRecord_Body()
}

type TransactionRecord_ContractCallResult struct {
	// *
	// Record of the value returned by the smart contract function (if it completed and didn't
	// fail) from ContractCallTransaction
	ContractCallResult *ContractFunctionResult `protobuf:"bytes,7,opt,name=contractCallResult,proto3,oneof"`
}

type TransactionRecord_ContractCreateResult struct {
	// *
	// Record of the value returned by the smart contract constructor (if it completed and
	// didn't fail) from ContractCreateTransaction
	ContractCreateResult *ContractFunctionResult `protobuf:"bytes,8,opt,name=contractCreateResult,proto3,oneof"`
}

func (*TransactionRecord_ContractCallResult) isTransactionRecord_Body() {}

func (*TransactionRecord_ContractCreateResult) isTransactionRecord_Body() {}

type isTransactionRecord_Entropy interface {
	isTransactionRecord_Entropy()
}

type TransactionRecord_PrngBytes struct {
	// *
	// In the record of a UtilPrng transaction with no output range, a pseudorandom 384-bit string.
	PrngBytes []byte `protobuf:"bytes,19,opt,name=prng_bytes,json=prngBytes,proto3,oneof"`
}

type TransactionRecord_PrngNumber struct {
	// *
	// In the record of a PRNG transaction with an output range, the output of a PRNG whose input was a 384-bit string.
	PrngNumber int32 `protobuf:"varint,20,opt,name=prng_number,json=prngNumber,proto3,oneof"`
}

func (*TransactionRecord_PrngBytes) isTransactionRecord_Entropy() {}

func (*TransactionRecord_PrngNumber) isTransactionRecord_Entropy() {}

var File_transaction_record_proto protoreflect.FileDescriptor

var file_transaction_record_proto_rawDesc = []byte{
	0x0a, 0x18, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x66, 0x65,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x08, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x33, 0x0a, 0x07, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74,
	0x12, 0x28, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x40, 0x0a, 0x12, 0x63, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x73, 0x75, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3a, 0x0a, 0x0d,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x26, 0x0a, 0x0e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x46, 0x65, 0x65, 0x12, 0x4f, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48,
	0x00, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x53, 0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x48, 0x00, 0x52, 0x14, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x37, 0x0a, 0x0c, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x12, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x12, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x33, 0x0a,
	0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x66, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52,
	0x65, 0x66, 0x12, 0x4a, 0x0a, 0x14, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x66, 0x65, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x65,
	0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x46, 0x65, 0x65, 0x52, 0x12, 0x61, 0x73, 0x73, 0x65,
	0x73, 0x73, 0x65, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x46, 0x65, 0x65, 0x73, 0x12, 0x59,
	0x0a, 0x1c, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0e,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x1a, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x73, 0x73,
	0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4e, 0x0a, 0x1a, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x18, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69,
	0x61, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12,
	0x23, 0x0a, 0x0d, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x46, 0x0a, 0x14, 0x70, 0x61, 0x69, 0x64, 0x5f, 0x73, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x12, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x12, 0x70, 0x61, 0x69, 0x64, 0x53, 0x74,
	0x61, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x1f, 0x0a, 0x0a,
	0x70, 0x72, 0x6e, 0x67, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0c,
	0x48, 0x01, 0x52, 0x09, 0x70, 0x72, 0x6e, 0x67, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x21, 0x0a,
	0x0b, 0x70, 0x72, 0x6e, 0x67, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x01, 0x52, 0x0a, 0x70, 0x72, 0x6e, 0x67, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x76, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x65, 0x76, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x09, 0x0a, 0x07, 0x65, 0x6e, 0x74,
	0x72, 0x6f, 0x70, 0x79, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transaction_record_proto_rawDescOnce sync.Once
	file_transaction_record_proto_rawDescData = file_transaction_record_proto_rawDesc
)

func file_transaction_record_proto_rawDescGZIP() []byte {
	file_transaction_record_proto_rawDescOnce.Do(func() {
		file_transaction_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_transaction_record_proto_rawDescData)
	})
	return file_transaction_record_proto_rawDescData
}

var file_transaction_record_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_transaction_record_proto_goTypes = []interface{}{
	(*TransactionRecord)(nil),      // 0: proto.TransactionRecord
	(*TransactionReceipt)(nil),     // 1: proto.TransactionReceipt
	(*Timestamp)(nil),              // 2: proto.Timestamp
	(*TransactionID)(nil),          // 3: proto.TransactionID
	(*ContractFunctionResult)(nil), // 4: proto.ContractFunctionResult
	(*TransferList)(nil),           // 5: proto.TransferList
	(*TokenTransferList)(nil),      // 6: proto.TokenTransferList
	(*ScheduleID)(nil),             // 7: proto.ScheduleID
	(*AssessedCustomFee)(nil),      // 8: proto.AssessedCustomFee
	(*TokenAssociation)(nil),       // 9: proto.TokenAssociation
	(*AccountAmount)(nil),          // 10: proto.AccountAmount
}
var file_transaction_record_proto_depIdxs = []int32{
	1,  // 0: proto.TransactionRecord.receipt:type_name -> proto.TransactionReceipt
	2,  // 1: proto.TransactionRecord.consensusTimestamp:type_name -> proto.Timestamp
	3,  // 2: proto.TransactionRecord.transactionID:type_name -> proto.TransactionID
	4,  // 3: proto.TransactionRecord.contractCallResult:type_name -> proto.ContractFunctionResult
	4,  // 4: proto.TransactionRecord.contractCreateResult:type_name -> proto.ContractFunctionResult
	5,  // 5: proto.TransactionRecord.transferList:type_name -> proto.TransferList
	6,  // 6: proto.TransactionRecord.tokenTransferLists:type_name -> proto.TokenTransferList
	7,  // 7: proto.TransactionRecord.scheduleRef:type_name -> proto.ScheduleID
	8,  // 8: proto.TransactionRecord.assessed_custom_fees:type_name -> proto.AssessedCustomFee
	9,  // 9: proto.TransactionRecord.automatic_token_associations:type_name -> proto.TokenAssociation
	2,  // 10: proto.TransactionRecord.parent_consensus_timestamp:type_name -> proto.Timestamp
	10, // 11: proto.TransactionRecord.paid_staking_rewards:type_name -> proto.AccountAmount
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_transaction_record_proto_init() }
func file_transaction_record_proto_init() {
	if File_transaction_record_proto != nil {
		return
	}
	file_timestamp_proto_init()
	file_basic_types_proto_init()
	file_custom_fees_proto_init()
	file_transaction_receipt_proto_init()
	file_crypto_transfer_proto_init()
	file_contract_call_local_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_transaction_record_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionRecord); i {
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
	file_transaction_record_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TransactionRecord_ContractCallResult)(nil),
		(*TransactionRecord_ContractCreateResult)(nil),
		(*TransactionRecord_PrngBytes)(nil),
		(*TransactionRecord_PrngNumber)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transaction_record_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transaction_record_proto_goTypes,
		DependencyIndexes: file_transaction_record_proto_depIdxs,
		MessageInfos:      file_transaction_record_proto_msgTypes,
	}.Build()
	File_transaction_record_proto = out.File
	file_transaction_record_proto_rawDesc = nil
	file_transaction_record_proto_goTypes = nil
	file_transaction_record_proto_depIdxs = nil
}
