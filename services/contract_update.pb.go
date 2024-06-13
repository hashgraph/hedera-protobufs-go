// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: contract_update.proto

package services

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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
// At consensus, updates the fields of a smart contract to the given values.
//
// If no value is given for a field, that field is left unchanged on the contract. For an immutable
// smart contract (that is, a contract created without an adminKey), only the expirationTime may be
// updated; setting any other field in this case will cause the transaction status to resolve to
// MODIFYING_IMMUTABLE_CONTRACT.
//
// --- Signing Requirements ---
//  1. Whether or not a contract has an admin key, its expiry can be extended with only the
//     transaction payer's signature.
//  2. Updating any other field of a mutable contract requires the admin key's signature.
//  3. If the update transaction includes a new admin key, this new key must also sign <b>unless</b>
//     it is exactly an empty <tt>KeyList</tt>. This special sentinel key removes the existing admin
//     key and causes the contract to become immutable. (Other <tt>Key</tt> structures without a
//     constituent <tt>Ed25519</tt> key will be rejected with <tt>INVALID_ADMIN_KEY</tt>.)
//  4. If the update transaction sets the AccountID auto_renew_account_id wrapper field to anything
//     other than the sentinel <tt>0.0.0</tt> value, then the key of the referenced account must sign.
type ContractUpdateTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// The id of the contract to be updated
	ContractID *ContractID `protobuf:"bytes,1,opt,name=contractID,proto3" json:"contractID,omitempty"`
	// *
	// The new expiry of the contract, no earlier than the current expiry (resolves to
	// EXPIRATION_REDUCTION_NOT_ALLOWED otherwise)
	ExpirationTime *Timestamp `protobuf:"bytes,2,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`
	// *
	// The new key to control updates to the contract
	AdminKey *Key `protobuf:"bytes,3,opt,name=adminKey,proto3" json:"adminKey,omitempty"`
	// *
	// [Deprecated] The new id of the account to which the contract is proxy staked
	//
	// Deprecated: Marked as deprecated in contract_update.proto.
	ProxyAccountID *AccountID `protobuf:"bytes,6,opt,name=proxyAccountID,proto3" json:"proxyAccountID,omitempty"`
	// *
	// If an auto-renew account is in use, the lifetime to be added by each auto-renewal.
	AutoRenewPeriod *Duration `protobuf:"bytes,7,opt,name=autoRenewPeriod,proto3" json:"autoRenewPeriod,omitempty"`
	// *
	// This field is unused and will have no impact on the specified smart contract.
	//
	// Deprecated: Marked as deprecated in contract_update.proto.
	FileID *FileID `protobuf:"bytes,8,opt,name=fileID,proto3" json:"fileID,omitempty"`
	// *
	// The new contract memo, assumed to be Unicode encoded with UTF-8 (at most 100 bytes)
	//
	// Types that are assignable to MemoField:
	//
	//	*ContractUpdateTransactionBody_Memo
	//	*ContractUpdateTransactionBody_MemoWrapper
	MemoField isContractUpdateTransactionBody_MemoField `protobuf_oneof:"memoField"`
	// *
	// If set, the new maximum number of tokens that this contract can be
	// automatically associated with (i.e., receive air-drops from).
	MaxAutomaticTokenAssociations *wrapperspb.Int32Value `protobuf:"bytes,11,opt,name=max_automatic_token_associations,json=maxAutomaticTokenAssociations,proto3" json:"max_automatic_token_associations,omitempty"`
	// *
	// If set to the sentinel <tt>0.0.0</tt> AccountID, this field removes the contract's auto-renew
	// account. Otherwise it updates the contract's auto-renew account to the referenced account.
	AutoRenewAccountId *AccountID `protobuf:"bytes,12,opt,name=auto_renew_account_id,json=autoRenewAccountId,proto3" json:"auto_renew_account_id,omitempty"`
	// *
	// ID of the new account or node to which this contract is staking.
	//
	// Types that are assignable to StakedId:
	//
	//	*ContractUpdateTransactionBody_StakedAccountId
	//	*ContractUpdateTransactionBody_StakedNodeId
	StakedId isContractUpdateTransactionBody_StakedId `protobuf_oneof:"staked_id"`
	// *
	// If true, the contract declines receiving a staking reward.
	DeclineReward *wrapperspb.BoolValue `protobuf:"bytes,15,opt,name=decline_reward,json=declineReward,proto3" json:"decline_reward,omitempty"`
}

func (x *ContractUpdateTransactionBody) Reset() {
	*x = ContractUpdateTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_update_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractUpdateTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractUpdateTransactionBody) ProtoMessage() {}

func (x *ContractUpdateTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_contract_update_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractUpdateTransactionBody.ProtoReflect.Descriptor instead.
func (*ContractUpdateTransactionBody) Descriptor() ([]byte, []int) {
	return file_contract_update_proto_rawDescGZIP(), []int{0}
}

func (x *ContractUpdateTransactionBody) GetContractID() *ContractID {
	if x != nil {
		return x.ContractID
	}
	return nil
}

func (x *ContractUpdateTransactionBody) GetExpirationTime() *Timestamp {
	if x != nil {
		return x.ExpirationTime
	}
	return nil
}

func (x *ContractUpdateTransactionBody) GetAdminKey() *Key {
	if x != nil {
		return x.AdminKey
	}
	return nil
}

// Deprecated: Marked as deprecated in contract_update.proto.
func (x *ContractUpdateTransactionBody) GetProxyAccountID() *AccountID {
	if x != nil {
		return x.ProxyAccountID
	}
	return nil
}

func (x *ContractUpdateTransactionBody) GetAutoRenewPeriod() *Duration {
	if x != nil {
		return x.AutoRenewPeriod
	}
	return nil
}

// Deprecated: Marked as deprecated in contract_update.proto.
func (x *ContractUpdateTransactionBody) GetFileID() *FileID {
	if x != nil {
		return x.FileID
	}
	return nil
}

func (m *ContractUpdateTransactionBody) GetMemoField() isContractUpdateTransactionBody_MemoField {
	if m != nil {
		return m.MemoField
	}
	return nil
}

// Deprecated: Marked as deprecated in contract_update.proto.
func (x *ContractUpdateTransactionBody) GetMemo() string {
	if x, ok := x.GetMemoField().(*ContractUpdateTransactionBody_Memo); ok {
		return x.Memo
	}
	return ""
}

func (x *ContractUpdateTransactionBody) GetMemoWrapper() *wrapperspb.StringValue {
	if x, ok := x.GetMemoField().(*ContractUpdateTransactionBody_MemoWrapper); ok {
		return x.MemoWrapper
	}
	return nil
}

func (x *ContractUpdateTransactionBody) GetMaxAutomaticTokenAssociations() *wrapperspb.Int32Value {
	if x != nil {
		return x.MaxAutomaticTokenAssociations
	}
	return nil
}

func (x *ContractUpdateTransactionBody) GetAutoRenewAccountId() *AccountID {
	if x != nil {
		return x.AutoRenewAccountId
	}
	return nil
}

func (m *ContractUpdateTransactionBody) GetStakedId() isContractUpdateTransactionBody_StakedId {
	if m != nil {
		return m.StakedId
	}
	return nil
}

func (x *ContractUpdateTransactionBody) GetStakedAccountId() *AccountID {
	if x, ok := x.GetStakedId().(*ContractUpdateTransactionBody_StakedAccountId); ok {
		return x.StakedAccountId
	}
	return nil
}

func (x *ContractUpdateTransactionBody) GetStakedNodeId() int64 {
	if x, ok := x.GetStakedId().(*ContractUpdateTransactionBody_StakedNodeId); ok {
		return x.StakedNodeId
	}
	return 0
}

func (x *ContractUpdateTransactionBody) GetDeclineReward() *wrapperspb.BoolValue {
	if x != nil {
		return x.DeclineReward
	}
	return nil
}

type isContractUpdateTransactionBody_MemoField interface {
	isContractUpdateTransactionBody_MemoField()
}

type ContractUpdateTransactionBody_Memo struct {
	// *
	// [Deprecated] If set with a non-zero length, the new memo to be associated with the account
	// (UTF-8 encoding max 100 bytes)
	//
	// Deprecated: Marked as deprecated in contract_update.proto.
	Memo string `protobuf:"bytes,9,opt,name=memo,proto3,oneof"`
}

type ContractUpdateTransactionBody_MemoWrapper struct {
	// *
	// If set, the new memo to be associated with the account (UTF-8 encoding max 100 bytes)
	MemoWrapper *wrapperspb.StringValue `protobuf:"bytes,10,opt,name=memoWrapper,proto3,oneof"`
}

func (*ContractUpdateTransactionBody_Memo) isContractUpdateTransactionBody_MemoField() {}

func (*ContractUpdateTransactionBody_MemoWrapper) isContractUpdateTransactionBody_MemoField() {}

type isContractUpdateTransactionBody_StakedId interface {
	isContractUpdateTransactionBody_StakedId()
}

type ContractUpdateTransactionBody_StakedAccountId struct {
	// *
	// ID of the new account to which this contract is staking. If set to the sentinel <tt>0.0.0</tt> AccountID,
	// this field removes the contract's staked account ID.
	StakedAccountId *AccountID `protobuf:"bytes,13,opt,name=staked_account_id,json=stakedAccountId,proto3,oneof"`
}

type ContractUpdateTransactionBody_StakedNodeId struct {
	// *
	// ID of the new node this contract is staked to. If set to the sentinel <tt>-1</tt>, this field
	// removes the contract's staked node ID.
	StakedNodeId int64 `protobuf:"varint,14,opt,name=staked_node_id,json=stakedNodeId,proto3,oneof"`
}

func (*ContractUpdateTransactionBody_StakedAccountId) isContractUpdateTransactionBody_StakedId() {}

func (*ContractUpdateTransactionBody_StakedNodeId) isContractUpdateTransactionBody_StakedId() {}

var File_contract_update_proto protoreflect.FileDescriptor

var file_contract_update_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0e, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa4, 0x06, 0x0a, 0x1d, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x6f, 0x64, 0x79, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x52, 0x0a, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x12, 0x38, 0x0a, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x26, 0x0a, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52,
	0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x3c, 0x0a, 0x0e, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x44, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x39, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x52,
	0x65, 0x6e, 0x65, 0x77, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x12, 0x29, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49,
	0x44, 0x42, 0x02, 0x18, 0x01, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x18, 0x0a,
	0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x48,
	0x00, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x40, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x57,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x6d, 0x65,
	0x6d, 0x6f, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x64, 0x0a, 0x20, 0x6d, 0x61, 0x78,
	0x5f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x1d, 0x6d, 0x61, 0x78, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x43, 0x0a, 0x15, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44,
	0x52, 0x12, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x64, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x44, 0x48, 0x01, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x64, 0x5f, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x0c,
	0x73, 0x74, 0x61, 0x6b, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x0e,
	0x64, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0d, 0x64, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x42,
	0x0b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x6f, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x0b, 0x0a, 0x09,
	0x73, 0x74, 0x61, 0x6b, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d,
	0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50,
	0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contract_update_proto_rawDescOnce sync.Once
	file_contract_update_proto_rawDescData = file_contract_update_proto_rawDesc
)

func file_contract_update_proto_rawDescGZIP() []byte {
	file_contract_update_proto_rawDescOnce.Do(func() {
		file_contract_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_update_proto_rawDescData)
	})
	return file_contract_update_proto_rawDescData
}

var file_contract_update_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_contract_update_proto_goTypes = []interface{}{
	(*ContractUpdateTransactionBody)(nil), // 0: proto.ContractUpdateTransactionBody
	(*ContractID)(nil),                    // 1: proto.ContractID
	(*Timestamp)(nil),                     // 2: proto.Timestamp
	(*Key)(nil),                           // 3: proto.Key
	(*AccountID)(nil),                     // 4: proto.AccountID
	(*Duration)(nil),                      // 5: proto.Duration
	(*FileID)(nil),                        // 6: proto.FileID
	(*wrapperspb.StringValue)(nil),        // 7: google.protobuf.StringValue
	(*wrapperspb.Int32Value)(nil),         // 8: google.protobuf.Int32Value
	(*wrapperspb.BoolValue)(nil),          // 9: google.protobuf.BoolValue
}
var file_contract_update_proto_depIdxs = []int32{
	1,  // 0: proto.ContractUpdateTransactionBody.contractID:type_name -> proto.ContractID
	2,  // 1: proto.ContractUpdateTransactionBody.expirationTime:type_name -> proto.Timestamp
	3,  // 2: proto.ContractUpdateTransactionBody.adminKey:type_name -> proto.Key
	4,  // 3: proto.ContractUpdateTransactionBody.proxyAccountID:type_name -> proto.AccountID
	5,  // 4: proto.ContractUpdateTransactionBody.autoRenewPeriod:type_name -> proto.Duration
	6,  // 5: proto.ContractUpdateTransactionBody.fileID:type_name -> proto.FileID
	7,  // 6: proto.ContractUpdateTransactionBody.memoWrapper:type_name -> google.protobuf.StringValue
	8,  // 7: proto.ContractUpdateTransactionBody.max_automatic_token_associations:type_name -> google.protobuf.Int32Value
	4,  // 8: proto.ContractUpdateTransactionBody.auto_renew_account_id:type_name -> proto.AccountID
	4,  // 9: proto.ContractUpdateTransactionBody.staked_account_id:type_name -> proto.AccountID
	9,  // 10: proto.ContractUpdateTransactionBody.decline_reward:type_name -> google.protobuf.BoolValue
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_contract_update_proto_init() }
func file_contract_update_proto_init() {
	if File_contract_update_proto != nil {
		return
	}
	file_basic_types_proto_init()
	file_duration_proto_init()
	file_timestamp_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_contract_update_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractUpdateTransactionBody); i {
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
	file_contract_update_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ContractUpdateTransactionBody_Memo)(nil),
		(*ContractUpdateTransactionBody_MemoWrapper)(nil),
		(*ContractUpdateTransactionBody_StakedAccountId)(nil),
		(*ContractUpdateTransactionBody_StakedNodeId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_contract_update_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contract_update_proto_goTypes,
		DependencyIndexes: file_contract_update_proto_depIdxs,
		MessageInfos:      file_contract_update_proto_msgTypes,
	}.Build()
	File_contract_update_proto = out.File
	file_contract_update_proto_rawDesc = nil
	file_contract_update_proto_goTypes = nil
	file_contract_update_proto_depIdxs = nil
}
