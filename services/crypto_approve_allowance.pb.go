// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: crypto_approve_allowance.proto

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
// Creates one or more hbar/token approved allowances <b>relative to the owner account specified in the allowances of
// this transaction</b>. Each allowance grants a spender the right to transfer a pre-determined amount of the owner's
// hbar/token to any other account of the spender's choice. If the owner is not specified in any allowance, the payer
// of transaction is considered to be the owner for that particular allowance.
// Setting the amount to zero in CryptoAllowance or TokenAllowance will remove the respective allowance for the spender.
//
// (So if account <tt>0.0.X</tt> pays for this transaction and owner is not specified in the allowance,
// then at consensus each spender account will have new allowances to spend hbar or tokens from <tt>0.0.X</tt>).
type CryptoApproveAllowanceTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// List of hbar allowances approved by the account owner.
	CryptoAllowances []*CryptoAllowance `protobuf:"bytes,1,rep,name=cryptoAllowances,proto3" json:"cryptoAllowances,omitempty"`
	// *
	// List of non-fungible token allowances approved by the account owner.
	NftAllowances []*NftAllowance `protobuf:"bytes,2,rep,name=nftAllowances,proto3" json:"nftAllowances,omitempty"`
	// *
	// List of fungible token allowances approved by the account owner.
	TokenAllowances []*TokenAllowance `protobuf:"bytes,3,rep,name=tokenAllowances,proto3" json:"tokenAllowances,omitempty"`
}

func (x *CryptoApproveAllowanceTransactionBody) Reset() {
	*x = CryptoApproveAllowanceTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_approve_allowance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoApproveAllowanceTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoApproveAllowanceTransactionBody) ProtoMessage() {}

func (x *CryptoApproveAllowanceTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_approve_allowance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoApproveAllowanceTransactionBody.ProtoReflect.Descriptor instead.
func (*CryptoApproveAllowanceTransactionBody) Descriptor() ([]byte, []int) {
	return file_crypto_approve_allowance_proto_rawDescGZIP(), []int{0}
}

func (x *CryptoApproveAllowanceTransactionBody) GetCryptoAllowances() []*CryptoAllowance {
	if x != nil {
		return x.CryptoAllowances
	}
	return nil
}

func (x *CryptoApproveAllowanceTransactionBody) GetNftAllowances() []*NftAllowance {
	if x != nil {
		return x.NftAllowances
	}
	return nil
}

func (x *CryptoApproveAllowanceTransactionBody) GetTokenAllowances() []*TokenAllowance {
	if x != nil {
		return x.TokenAllowances
	}
	return nil
}

// *
// An approved allowance of hbar transfers for a spender.
type CryptoAllowance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// The account ID of the hbar owner (ie. the grantor of the allowance).
	Owner *AccountID `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// *
	// The account ID of the spender of the hbar allowance.
	Spender *AccountID `protobuf:"bytes,2,opt,name=spender,proto3" json:"spender,omitempty"`
	// *
	// The amount of the spender's allowance in tinybars.
	Amount int64 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *CryptoAllowance) Reset() {
	*x = CryptoAllowance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_approve_allowance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoAllowance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoAllowance) ProtoMessage() {}

func (x *CryptoAllowance) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_approve_allowance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoAllowance.ProtoReflect.Descriptor instead.
func (*CryptoAllowance) Descriptor() ([]byte, []int) {
	return file_crypto_approve_allowance_proto_rawDescGZIP(), []int{1}
}

func (x *CryptoAllowance) GetOwner() *AccountID {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *CryptoAllowance) GetSpender() *AccountID {
	if x != nil {
		return x.Spender
	}
	return nil
}

func (x *CryptoAllowance) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

// *
// An approved allowance of non-fungible token transfers for a spender.
type NftAllowance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// The NFT token type that the allowance pertains to.
	TokenId *TokenID `protobuf:"bytes,1,opt,name=tokenId,proto3" json:"tokenId,omitempty"`
	// *
	// The account ID of the token owner (ie. the grantor of the allowance).
	Owner *AccountID `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// *
	// The account ID of the token allowance spender.
	Spender *AccountID `protobuf:"bytes,3,opt,name=spender,proto3" json:"spender,omitempty"`
	// *
	// The list of serial numbers that the spender is permitted to transfer.
	SerialNumbers []int64 `protobuf:"varint,4,rep,packed,name=serial_numbers,json=serialNumbers,proto3" json:"serial_numbers,omitempty"`
	// *
	// If true, the spender has access to all of the owner's NFT units of type tokenId (currently
	// owned and any in the future).
	ApprovedForAll *wrapperspb.BoolValue `protobuf:"bytes,5,opt,name=approved_for_all,json=approvedForAll,proto3" json:"approved_for_all,omitempty"`
	// *
	// The account ID of the spender who is granted approvedForAll allowance and granting
	// approval on an NFT serial to another spender.
	DelegatingSpender *AccountID `protobuf:"bytes,6,opt,name=delegating_spender,json=delegatingSpender,proto3" json:"delegating_spender,omitempty"`
}

func (x *NftAllowance) Reset() {
	*x = NftAllowance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_approve_allowance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NftAllowance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftAllowance) ProtoMessage() {}

func (x *NftAllowance) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_approve_allowance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftAllowance.ProtoReflect.Descriptor instead.
func (*NftAllowance) Descriptor() ([]byte, []int) {
	return file_crypto_approve_allowance_proto_rawDescGZIP(), []int{2}
}

func (x *NftAllowance) GetTokenId() *TokenID {
	if x != nil {
		return x.TokenId
	}
	return nil
}

func (x *NftAllowance) GetOwner() *AccountID {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *NftAllowance) GetSpender() *AccountID {
	if x != nil {
		return x.Spender
	}
	return nil
}

func (x *NftAllowance) GetSerialNumbers() []int64 {
	if x != nil {
		return x.SerialNumbers
	}
	return nil
}

func (x *NftAllowance) GetApprovedForAll() *wrapperspb.BoolValue {
	if x != nil {
		return x.ApprovedForAll
	}
	return nil
}

func (x *NftAllowance) GetDelegatingSpender() *AccountID {
	if x != nil {
		return x.DelegatingSpender
	}
	return nil
}

// *
// An approved allowance of fungible token transfers for a spender.
type TokenAllowance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// The token that the allowance pertains to.
	TokenId *TokenID `protobuf:"bytes,1,opt,name=tokenId,proto3" json:"tokenId,omitempty"`
	// *
	// The account ID of the token owner (ie. the grantor of the allowance).
	Owner *AccountID `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// *
	// The account ID of the token allowance spender.
	Spender *AccountID `protobuf:"bytes,3,opt,name=spender,proto3" json:"spender,omitempty"`
	// *
	// The amount of the spender's token allowance.
	Amount int64 `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TokenAllowance) Reset() {
	*x = TokenAllowance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_approve_allowance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenAllowance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenAllowance) ProtoMessage() {}

func (x *TokenAllowance) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_approve_allowance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenAllowance.ProtoReflect.Descriptor instead.
func (*TokenAllowance) Descriptor() ([]byte, []int) {
	return file_crypto_approve_allowance_proto_rawDescGZIP(), []int{3}
}

func (x *TokenAllowance) GetTokenId() *TokenID {
	if x != nil {
		return x.TokenId
	}
	return nil
}

func (x *TokenAllowance) GetOwner() *AccountID {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *TokenAllowance) GetSpender() *AccountID {
	if x != nil {
		return x.Spender
	}
	return nil
}

func (x *TokenAllowance) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_crypto_approve_allowance_proto protoreflect.FileDescriptor

var file_crypto_approve_allowance_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x01, 0x0a, 0x25, 0x43,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x41, 0x6c, 0x6c, 0x6f,
	0x77, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x6f, 0x64, 0x79, 0x12, 0x42, 0x0a, 0x10, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x41, 0x6c,
	0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x41, 0x6c, 0x6c,
	0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x10, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x41, 0x6c,
	0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0d, 0x6e, 0x66, 0x74, 0x41,
	0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x66, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x77,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x0d, 0x6e, 0x66, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x0f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x6c, 0x6c, 0x6f,
	0x77, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x0f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61,
	0x6e, 0x63, 0x65, 0x73, 0x22, 0x7d, 0x0a, 0x0f, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x41, 0x6c,
	0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12,
	0x2a, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x44, 0x52, 0x07, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0xba, 0x02, 0x0a, 0x0c, 0x4e, 0x66, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x77,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x49, 0x44, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x26,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x07, 0x73, 0x70, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x44, 0x0a, 0x10, 0x61, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x65, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x61, 0x6c, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x41, 0x6c, 0x6c, 0x12,
	0x3f, 0x0a, 0x12, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x70,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x11, 0x64,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x22, 0xa6, 0x01, 0x0a, 0x0e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x49, 0x44, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x26, 0x0a,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x07, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d,
	0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50,
	0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crypto_approve_allowance_proto_rawDescOnce sync.Once
	file_crypto_approve_allowance_proto_rawDescData = file_crypto_approve_allowance_proto_rawDesc
)

func file_crypto_approve_allowance_proto_rawDescGZIP() []byte {
	file_crypto_approve_allowance_proto_rawDescOnce.Do(func() {
		file_crypto_approve_allowance_proto_rawDescData = protoimpl.X.CompressGZIP(file_crypto_approve_allowance_proto_rawDescData)
	})
	return file_crypto_approve_allowance_proto_rawDescData
}

var file_crypto_approve_allowance_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_crypto_approve_allowance_proto_goTypes = []interface{}{
	(*CryptoApproveAllowanceTransactionBody)(nil), // 0: proto.CryptoApproveAllowanceTransactionBody
	(*CryptoAllowance)(nil),                       // 1: proto.CryptoAllowance
	(*NftAllowance)(nil),                          // 2: proto.NftAllowance
	(*TokenAllowance)(nil),                        // 3: proto.TokenAllowance
	(*AccountID)(nil),                             // 4: proto.AccountID
	(*TokenID)(nil),                               // 5: proto.TokenID
	(*wrapperspb.BoolValue)(nil),                  // 6: google.protobuf.BoolValue
}
var file_crypto_approve_allowance_proto_depIdxs = []int32{
	1,  // 0: proto.CryptoApproveAllowanceTransactionBody.cryptoAllowances:type_name -> proto.CryptoAllowance
	2,  // 1: proto.CryptoApproveAllowanceTransactionBody.nftAllowances:type_name -> proto.NftAllowance
	3,  // 2: proto.CryptoApproveAllowanceTransactionBody.tokenAllowances:type_name -> proto.TokenAllowance
	4,  // 3: proto.CryptoAllowance.owner:type_name -> proto.AccountID
	4,  // 4: proto.CryptoAllowance.spender:type_name -> proto.AccountID
	5,  // 5: proto.NftAllowance.tokenId:type_name -> proto.TokenID
	4,  // 6: proto.NftAllowance.owner:type_name -> proto.AccountID
	4,  // 7: proto.NftAllowance.spender:type_name -> proto.AccountID
	6,  // 8: proto.NftAllowance.approved_for_all:type_name -> google.protobuf.BoolValue
	4,  // 9: proto.NftAllowance.delegating_spender:type_name -> proto.AccountID
	5,  // 10: proto.TokenAllowance.tokenId:type_name -> proto.TokenID
	4,  // 11: proto.TokenAllowance.owner:type_name -> proto.AccountID
	4,  // 12: proto.TokenAllowance.spender:type_name -> proto.AccountID
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_crypto_approve_allowance_proto_init() }
func file_crypto_approve_allowance_proto_init() {
	if File_crypto_approve_allowance_proto != nil {
		return
	}
	file_basic_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_crypto_approve_allowance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoApproveAllowanceTransactionBody); i {
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
		file_crypto_approve_allowance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoAllowance); i {
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
		file_crypto_approve_allowance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NftAllowance); i {
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
		file_crypto_approve_allowance_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenAllowance); i {
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
			RawDescriptor: file_crypto_approve_allowance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_crypto_approve_allowance_proto_goTypes,
		DependencyIndexes: file_crypto_approve_allowance_proto_depIdxs,
		MessageInfos:      file_crypto_approve_allowance_proto_msgTypes,
	}.Build()
	File_crypto_approve_allowance_proto = out.File
	file_crypto_approve_allowance_proto_rawDesc = nil
	file_crypto_approve_allowance_proto_goTypes = nil
	file_crypto_approve_allowance_proto_depIdxs = nil
}
