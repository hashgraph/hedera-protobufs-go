// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: CryptoGetInfo.proto

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

// Get all the information about an account, including the balance. This does not get the list of account records.
type CryptoGetInfoQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header    *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`       // Standard info sent from client to node, including the signed payment, and what kind of response is requested (cost, state proof, both, or neither).
	AccountID *AccountID   `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"` // The account ID for which information is requested
}

func (x *CryptoGetInfoQuery) Reset() {
	*x = CryptoGetInfoQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CryptoGetInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoGetInfoQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoGetInfoQuery) ProtoMessage() {}

func (x *CryptoGetInfoQuery) ProtoReflect() protoreflect.Message {
	mi := &file_CryptoGetInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoGetInfoQuery.ProtoReflect.Descriptor instead.
func (*CryptoGetInfoQuery) Descriptor() ([]byte, []int) {
	return file_CryptoGetInfo_proto_rawDescGZIP(), []int{0}
}

func (x *CryptoGetInfoQuery) GetHeader() *QueryHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *CryptoGetInfoQuery) GetAccountID() *AccountID {
	if x != nil {
		return x.AccountID
	}
	return nil
}

// Response when the client sends the node CryptoGetInfoQuery
type CryptoGetInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header      *ResponseHeader                    `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`           //Standard response from node to client, including the requested fields: cost, or state proof, or both, or neither
	AccountInfo *CryptoGetInfoResponse_AccountInfo `protobuf:"bytes,2,opt,name=accountInfo,proto3" json:"accountInfo,omitempty"` // Info about the account (a state proof can be generated for this)
}

func (x *CryptoGetInfoResponse) Reset() {
	*x = CryptoGetInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CryptoGetInfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoGetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoGetInfoResponse) ProtoMessage() {}

func (x *CryptoGetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_CryptoGetInfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoGetInfoResponse.ProtoReflect.Descriptor instead.
func (*CryptoGetInfoResponse) Descriptor() ([]byte, []int) {
	return file_CryptoGetInfo_proto_rawDescGZIP(), []int{1}
}

func (x *CryptoGetInfoResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *CryptoGetInfoResponse) GetAccountInfo() *CryptoGetInfoResponse_AccountInfo {
	if x != nil {
		return x.AccountInfo
	}
	return nil
}

type CryptoGetInfoResponse_AccountInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID         *AccountID `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`                 // The account ID for which this information applies
	ContractAccountID string     `protobuf:"bytes,2,opt,name=contractAccountID,proto3" json:"contractAccountID,omitempty"` // The Contract Account ID comprising of both the contract instance and the cryptocurrency account owned by the contract instance, in the format used by Solidity
	Deleted           bool       `protobuf:"varint,3,opt,name=deleted,proto3" json:"deleted,omitempty"`                    // If true, then this account has been deleted, it will disappear when it expires, and all transactions for it will fail except the transaction to extend its expiration date
	ProxyAccountID    *AccountID `protobuf:"bytes,4,opt,name=proxyAccountID,proto3" json:"proxyAccountID,omitempty"`       // The Account ID of the account to which this is proxy staked. If proxyAccountID is null, or is an invalid account, or is an account that isn't a node, then this account is automatically proxy staked to a node chosen by the network, but without earning payments. If the proxyAccountID account refuses to accept proxy staking , or if it is not currently running a node, then it will behave as if proxyAccountID was null.
	ProxyReceived     int64      `protobuf:"varint,6,opt,name=proxyReceived,proto3" json:"proxyReceived,omitempty"`        // The total number of tinybars proxy staked to this account
	Key               *Key       `protobuf:"bytes,7,opt,name=key,proto3" json:"key,omitempty"`                             // The key for the account, which must sign in order to transfer out, or to modify the account in any way other than extending its expiration date.
	Balance           uint64     `protobuf:"varint,8,opt,name=balance,proto3" json:"balance,omitempty"`                    // The current balance of account in tinybars
	// [Deprecated]. The threshold amount, in tinybars, at which a record is created of any transaction that decreases the balance of this account by more than the threshold
	//
	// Deprecated: Do not use.
	GenerateSendRecordThreshold uint64 `protobuf:"varint,9,opt,name=generateSendRecordThreshold,proto3" json:"generateSendRecordThreshold,omitempty"`
	// [Deprecated]. The threshold amount, in tinybars, at which a record is created of any transaction that increases the balance of this account by more than the threshold
	//
	// Deprecated: Do not use.
	GenerateReceiveRecordThreshold uint64               `protobuf:"varint,10,opt,name=generateReceiveRecordThreshold,proto3" json:"generateReceiveRecordThreshold,omitempty"`
	ReceiverSigRequired            bool                 `protobuf:"varint,11,opt,name=receiverSigRequired,proto3" json:"receiverSigRequired,omitempty"` // If true, no transaction can transfer to this account unless signed by this account's key
	ExpirationTime                 *Timestamp           `protobuf:"bytes,12,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`            // The TimeStamp time at which this account is set to expire
	AutoRenewPeriod                *Duration            `protobuf:"bytes,13,opt,name=autoRenewPeriod,proto3" json:"autoRenewPeriod,omitempty"`          // The duration for expiration time will extend every this many seconds. If there are insufficient funds, then it extends as long as possible. If it is empty when it expires, then it is deleted.
	LiveHashes                     []*LiveHash          `protobuf:"bytes,14,rep,name=liveHashes,proto3" json:"liveHashes,omitempty"`                    // All of the livehashes attached to the account (each of which is a hash along with the keys that authorized it and can delete it)
	TokenRelationships             []*TokenRelationship `protobuf:"bytes,15,rep,name=tokenRelationships,proto3" json:"tokenRelationships,omitempty"`    // All tokens related to this account
	Memo                           string               `protobuf:"bytes,16,opt,name=memo,proto3" json:"memo,omitempty"`                                // The memo associated with the account
	OwnedNfts                      int64                `protobuf:"varint,17,opt,name=ownedNfts,proto3" json:"ownedNfts,omitempty"`                     // The number of NFTs owned by this account
}

func (x *CryptoGetInfoResponse_AccountInfo) Reset() {
	*x = CryptoGetInfoResponse_AccountInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CryptoGetInfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoGetInfoResponse_AccountInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoGetInfoResponse_AccountInfo) ProtoMessage() {}

func (x *CryptoGetInfoResponse_AccountInfo) ProtoReflect() protoreflect.Message {
	mi := &file_CryptoGetInfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoGetInfoResponse_AccountInfo.ProtoReflect.Descriptor instead.
func (*CryptoGetInfoResponse_AccountInfo) Descriptor() ([]byte, []int) {
	return file_CryptoGetInfo_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CryptoGetInfoResponse_AccountInfo) GetAccountID() *AccountID {
	if x != nil {
		return x.AccountID
	}
	return nil
}

func (x *CryptoGetInfoResponse_AccountInfo) GetContractAccountID() string {
	if x != nil {
		return x.ContractAccountID
	}
	return ""
}

func (x *CryptoGetInfoResponse_AccountInfo) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

func (x *CryptoGetInfoResponse_AccountInfo) GetProxyAccountID() *AccountID {
	if x != nil {
		return x.ProxyAccountID
	}
	return nil
}

func (x *CryptoGetInfoResponse_AccountInfo) GetProxyReceived() int64 {
	if x != nil {
		return x.ProxyReceived
	}
	return 0
}

func (x *CryptoGetInfoResponse_AccountInfo) GetKey() *Key {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *CryptoGetInfoResponse_AccountInfo) GetBalance() uint64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

// Deprecated: Do not use.
func (x *CryptoGetInfoResponse_AccountInfo) GetGenerateSendRecordThreshold() uint64 {
	if x != nil {
		return x.GenerateSendRecordThreshold
	}
	return 0
}

// Deprecated: Do not use.
func (x *CryptoGetInfoResponse_AccountInfo) GetGenerateReceiveRecordThreshold() uint64 {
	if x != nil {
		return x.GenerateReceiveRecordThreshold
	}
	return 0
}

func (x *CryptoGetInfoResponse_AccountInfo) GetReceiverSigRequired() bool {
	if x != nil {
		return x.ReceiverSigRequired
	}
	return false
}

func (x *CryptoGetInfoResponse_AccountInfo) GetExpirationTime() *Timestamp {
	if x != nil {
		return x.ExpirationTime
	}
	return nil
}

func (x *CryptoGetInfoResponse_AccountInfo) GetAutoRenewPeriod() *Duration {
	if x != nil {
		return x.AutoRenewPeriod
	}
	return nil
}

func (x *CryptoGetInfoResponse_AccountInfo) GetLiveHashes() []*LiveHash {
	if x != nil {
		return x.LiveHashes
	}
	return nil
}

func (x *CryptoGetInfoResponse_AccountInfo) GetTokenRelationships() []*TokenRelationship {
	if x != nil {
		return x.TokenRelationships
	}
	return nil
}

func (x *CryptoGetInfoResponse_AccountInfo) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *CryptoGetInfoResponse_AccountInfo) GetOwnedNfts() int64 {
	if x != nil {
		return x.OwnedNfts
	}
	return 0
}

var File_CryptoGetInfo_proto protoreflect.FileDescriptor

var file_CryptoGetInfo_proto_rawDesc = []byte{
	0x0a, 0x13, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x42,
	0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x14, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x41, 0x64, 0x64, 0x4c, 0x69, 0x76, 0x65, 0x48, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x70, 0x0a, 0x12, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x44, 0x22, 0x98, 0x07, 0x0a, 0x15, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x0b,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x83, 0x06, 0x0a, 0x0b, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x12, 0x38, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64,
	0x12, 0x1c, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x1b, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x18,
	0x01, 0x52, 0x1b, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x4a,
	0x0a, 0x1e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x18, 0x01, 0x52, 0x1e, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x53, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x53, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x0e,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65,
	0x6e, 0x65, 0x77, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x50, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x12, 0x2f, 0x0a, 0x0a, 0x6c, 0x69, 0x76, 0x65, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18,
	0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69,
	0x76, 0x65, 0x48, 0x61, 0x73, 0x68, 0x52, 0x0a, 0x6c, 0x69, 0x76, 0x65, 0x48, 0x61, 0x73, 0x68,
	0x65, 0x73, 0x12, 0x48, 0x0a, 0x12, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x12, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f,
	0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x64, 0x4e, 0x66, 0x74, 0x73, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x64, 0x4e, 0x66, 0x74, 0x73, 0x42, 0x26,
	0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CryptoGetInfo_proto_rawDescOnce sync.Once
	file_CryptoGetInfo_proto_rawDescData = file_CryptoGetInfo_proto_rawDesc
)

func file_CryptoGetInfo_proto_rawDescGZIP() []byte {
	file_CryptoGetInfo_proto_rawDescOnce.Do(func() {
		file_CryptoGetInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_CryptoGetInfo_proto_rawDescData)
	})
	return file_CryptoGetInfo_proto_rawDescData
}

var file_CryptoGetInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_CryptoGetInfo_proto_goTypes = []interface{}{
	(*CryptoGetInfoQuery)(nil),                // 0: proto.CryptoGetInfoQuery
	(*CryptoGetInfoResponse)(nil),             // 1: proto.CryptoGetInfoResponse
	(*CryptoGetInfoResponse_AccountInfo)(nil), // 2: proto.CryptoGetInfoResponse.AccountInfo
	(*QueryHeader)(nil),                       // 3: proto.QueryHeader
	(*AccountID)(nil),                         // 4: proto.AccountID
	(*ResponseHeader)(nil),                    // 5: proto.ResponseHeader
	(*Key)(nil),                               // 6: proto.Key
	(*Timestamp)(nil),                         // 7: proto.Timestamp
	(*Duration)(nil),                          // 8: proto.Duration
	(*LiveHash)(nil),                          // 9: proto.LiveHash
	(*TokenRelationship)(nil),                 // 10: proto.TokenRelationship
}
var file_CryptoGetInfo_proto_depIdxs = []int32{
	3,  // 0: proto.CryptoGetInfoQuery.header:type_name -> proto.QueryHeader
	4,  // 1: proto.CryptoGetInfoQuery.accountID:type_name -> proto.AccountID
	5,  // 2: proto.CryptoGetInfoResponse.header:type_name -> proto.ResponseHeader
	2,  // 3: proto.CryptoGetInfoResponse.accountInfo:type_name -> proto.CryptoGetInfoResponse.AccountInfo
	4,  // 4: proto.CryptoGetInfoResponse.AccountInfo.accountID:type_name -> proto.AccountID
	4,  // 5: proto.CryptoGetInfoResponse.AccountInfo.proxyAccountID:type_name -> proto.AccountID
	6,  // 6: proto.CryptoGetInfoResponse.AccountInfo.key:type_name -> proto.Key
	7,  // 7: proto.CryptoGetInfoResponse.AccountInfo.expirationTime:type_name -> proto.Timestamp
	8,  // 8: proto.CryptoGetInfoResponse.AccountInfo.autoRenewPeriod:type_name -> proto.Duration
	9,  // 9: proto.CryptoGetInfoResponse.AccountInfo.liveHashes:type_name -> proto.LiveHash
	10, // 10: proto.CryptoGetInfoResponse.AccountInfo.tokenRelationships:type_name -> proto.TokenRelationship
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_CryptoGetInfo_proto_init() }
func file_CryptoGetInfo_proto_init() {
	if File_CryptoGetInfo_proto != nil {
		return
	}
	file_Timestamp_proto_init()
	file_Duration_proto_init()
	file_BasicTypes_proto_init()
	file_QueryHeader_proto_init()
	file_ResponseHeader_proto_init()
	file_CryptoAddLiveHash_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CryptoGetInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoGetInfoQuery); i {
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
		file_CryptoGetInfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoGetInfoResponse); i {
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
		file_CryptoGetInfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoGetInfoResponse_AccountInfo); i {
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
			RawDescriptor: file_CryptoGetInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CryptoGetInfo_proto_goTypes,
		DependencyIndexes: file_CryptoGetInfo_proto_depIdxs,
		MessageInfos:      file_CryptoGetInfo_proto_msgTypes,
	}.Build()
	File_CryptoGetInfo_proto = out.File
	file_CryptoGetInfo_proto_rawDesc = nil
	file_CryptoGetInfo_proto_goTypes = nil
	file_CryptoGetInfo_proto_depIdxs = nil
}
