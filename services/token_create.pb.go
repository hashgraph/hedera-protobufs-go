// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: token_create.proto

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
// Create a new token. After the token is created, the Token ID for it is in the receipt.
// The specified Treasury Account is receiving the initial supply of tokens as-well as the tokens
// from the Token Mint operation once executed. The balance of the treasury account is decreased
// when the Token Burn operation is executed.
//
// The <tt>initialSupply</tt> is the initial supply of the smallest parts of a token (like a
// tinybar, not an hbar). These are the smallest units of the token which may be transferred.
//
// The supply can change over time. If the total supply at some moment is <i>S</i> parts of tokens,
// and the token is using <i>D</i> decimals, then <i>S</i> must be less than or equal to
// 2<sup>63</sup>-1, which is 9,223,372,036,854,775,807. The number of whole tokens (not parts) will
// be <i>S / 10<sup>D</sup></i>.
//
// If decimals is 8 or 11, then the number of whole tokens can be at most a few billions or
// millions, respectively. For example, it could match Bitcoin (21 million whole tokens with 8
// decimals) or hbars (50 billion whole tokens with 8 decimals). It could even match Bitcoin with
// milli-satoshis (21 million whole tokens with 11 decimals).
//
// Note that a created token is <i>immutable</i> if the <tt>adminKey</tt> is omitted. No property of
// an immutable token can ever change, with the sole exception of its expiry. Anyone can pay to
// extend the expiry time of an immutable token.
//
// A token can be either <i>FUNGIBLE_COMMON</i> or <i>NON_FUNGIBLE_UNIQUE</i>, based on its
// <i>TokenType</i>. If it has been omitted, <i>FUNGIBLE_COMMON</i> type is used.
//
// A token can have either <i>INFINITE</i> or <i>FINITE</i> supply type, based on its
// <i>TokenType</i>. If it has been omitted, <i>INFINITE</i> type is used.
//
// If a <i>FUNGIBLE</i> TokenType is used, <i>initialSupply</i> should explicitly be set to a
// non-negative. If not, the transaction will resolve to INVALID_TOKEN_INITIAL_SUPPLY.
//
// If a <i>NON_FUNGIBLE_UNIQUE</i> TokenType is used, <i>initialSupply</i> should explicitly be set
// to 0. If not, the transaction will resolve to INVALID_TOKEN_INITIAL_SUPPLY.
//
// If an <i>INFINITE</i> TokenSupplyType is used, <i>maxSupply</i> should explicitly be set to 0. If
// it is not 0, the transaction will resolve to INVALID_TOKEN_MAX_SUPPLY.
//
// If a <i>FINITE</i> TokenSupplyType is used, <i>maxSupply</i> should be explicitly set to a
// non-negative value. If it is not, the transaction will resolve to INVALID_TOKEN_MAX_SUPPLY.
type TokenCreateTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// The publicly visible name of the token. The token name is specified as a Unicode string.
	// Its UTF-8 encoding cannot exceed 100 bytes, and cannot contain the 0 byte (NUL).
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// *
	// The publicly visible token symbol. The token symbol is specified as a Unicode string.
	// Its UTF-8 encoding cannot exceed 100 bytes, and cannot contain the 0 byte (NUL).
	Symbol string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// *
	// For tokens of type FUNGIBLE_COMMON - the number of decimal places a
	// token is divisible by. For tokens of type NON_FUNGIBLE_UNIQUE - value
	// must be 0
	Decimals uint32 `protobuf:"varint,3,opt,name=decimals,proto3" json:"decimals,omitempty"`
	// *
	// Specifies the initial supply of tokens to be put in circulation. The
	// initial supply is sent to the Treasury Account. The supply is in the
	// lowest denomination possible. In the case for NON_FUNGIBLE_UNIQUE Type
	// the value must be 0
	InitialSupply uint64 `protobuf:"varint,4,opt,name=initialSupply,proto3" json:"initialSupply,omitempty"`
	// *
	// The account which will act as a treasury for the token. This account
	// will receive the specified initial supply or the newly minted NFTs in
	// the case for NON_FUNGIBLE_UNIQUE Type
	Treasury *AccountID `protobuf:"bytes,5,opt,name=treasury,proto3" json:"treasury,omitempty"`
	// *
	// The key which can perform update/delete operations on the token. If empty, the token can be
	// perceived as immutable (not being able to be updated/deleted)
	AdminKey *Key `protobuf:"bytes,6,opt,name=adminKey,proto3" json:"adminKey,omitempty"`
	// *
	// The key which can grant or revoke KYC of an account for the token's transactions. If empty,
	// KYC is not required, and KYC grant or revoke operations are not possible.
	KycKey *Key `protobuf:"bytes,7,opt,name=kycKey,proto3" json:"kycKey,omitempty"`
	// *
	// The key which can sign to freeze or unfreeze an account for token transactions. If empty,
	// freezing is not possible
	FreezeKey *Key `protobuf:"bytes,8,opt,name=freezeKey,proto3" json:"freezeKey,omitempty"`
	// *
	// The key which can wipe the token balance of an account. If empty, wipe is not possible
	WipeKey *Key `protobuf:"bytes,9,opt,name=wipeKey,proto3" json:"wipeKey,omitempty"`
	// *
	// The key which can change the supply of a token. The key is used to sign Token Mint/Burn
	// operations
	SupplyKey *Key `protobuf:"bytes,10,opt,name=supplyKey,proto3" json:"supplyKey,omitempty"`
	// *
	// The default Freeze status (frozen or unfrozen) of Hedera accounts relative to this token. If
	// true, an account must be unfrozen before it can receive the token
	FreezeDefault bool `protobuf:"varint,11,opt,name=freezeDefault,proto3" json:"freezeDefault,omitempty"`
	// *
	// The epoch second at which the token should expire; if an auto-renew account and period are
	// specified, this is coerced to the current epoch second plus the autoRenewPeriod
	Expiry *Timestamp `protobuf:"bytes,13,opt,name=expiry,proto3" json:"expiry,omitempty"`
	// *
	// An account which will be automatically charged to renew the token's expiration, at
	// autoRenewPeriod interval
	AutoRenewAccount *AccountID `protobuf:"bytes,14,opt,name=autoRenewAccount,proto3" json:"autoRenewAccount,omitempty"`
	// *
	// The interval at which the auto-renew account will be charged to extend the token's expiry
	AutoRenewPeriod *Duration `protobuf:"bytes,15,opt,name=autoRenewPeriod,proto3" json:"autoRenewPeriod,omitempty"`
	// *
	// The memo associated with the token (UTF-8 encoding max 100 bytes)
	Memo string `protobuf:"bytes,16,opt,name=memo,proto3" json:"memo,omitempty"`
	// *
	// IWA compatibility. Specifies the token type. Defaults to FUNGIBLE_COMMON
	TokenType TokenType `protobuf:"varint,17,opt,name=tokenType,proto3,enum=proto.TokenType" json:"tokenType,omitempty"`
	// *
	// IWA compatibility. Specified the token supply type. Defaults to INFINITE
	SupplyType TokenSupplyType `protobuf:"varint,18,opt,name=supplyType,proto3,enum=proto.TokenSupplyType" json:"supplyType,omitempty"`
	// *
	// IWA Compatibility. Depends on TokenSupplyType. For tokens of type FUNGIBLE_COMMON - the
	// maximum number of tokens that can be in circulation. For tokens of type NON_FUNGIBLE_UNIQUE -
	// the maximum number of NFTs (serial numbers) that can be minted. This field can never be
	// changed!
	MaxSupply int64 `protobuf:"varint,19,opt,name=maxSupply,proto3" json:"maxSupply,omitempty"`
	// *
	// The key which can change the token's custom fee schedule; must sign a TokenFeeScheduleUpdate
	// transaction
	FeeScheduleKey *Key `protobuf:"bytes,20,opt,name=fee_schedule_key,json=feeScheduleKey,proto3" json:"fee_schedule_key,omitempty"`
	// *
	// The custom fees to be assessed during a CryptoTransfer that transfers units of this token
	CustomFees []*CustomFee `protobuf:"bytes,21,rep,name=custom_fees,json=customFees,proto3" json:"custom_fees,omitempty"`
	// *
	// The Key which can pause and unpause the Token.
	// If Empty the token pause status defaults to PauseNotApplicable, otherwise Unpaused.
	PauseKey *Key `protobuf:"bytes,22,opt,name=pause_key,json=pauseKey,proto3" json:"pause_key,omitempty"`
	// *
	// Metadata of the created token definition.
	Metadata []byte `protobuf:"bytes,23,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// *
	// The key which can change the metadata of a token
	// (token definition, partition definition, and individual NFTs).
	MetadataKey *Key `protobuf:"bytes,24,opt,name=metadata_key,json=metadataKey,proto3" json:"metadata_key,omitempty"`
}

func (x *TokenCreateTransactionBody) Reset() {
	*x = TokenCreateTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenCreateTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenCreateTransactionBody) ProtoMessage() {}

func (x *TokenCreateTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_token_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenCreateTransactionBody.ProtoReflect.Descriptor instead.
func (*TokenCreateTransactionBody) Descriptor() ([]byte, []int) {
	return file_token_create_proto_rawDescGZIP(), []int{0}
}

func (x *TokenCreateTransactionBody) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TokenCreateTransactionBody) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *TokenCreateTransactionBody) GetDecimals() uint32 {
	if x != nil {
		return x.Decimals
	}
	return 0
}

func (x *TokenCreateTransactionBody) GetInitialSupply() uint64 {
	if x != nil {
		return x.InitialSupply
	}
	return 0
}

func (x *TokenCreateTransactionBody) GetTreasury() *AccountID {
	if x != nil {
		return x.Treasury
	}
	return nil
}

func (x *TokenCreateTransactionBody) GetAdminKey() *Key {
	if x != nil {
		return x.AdminKey
	}
	return nil
}

func (x *TokenCreateTransactionBody) GetKycKey() *Key {
	if x != nil {
		return x.KycKey
	}
	return nil
}

func (x *TokenCreateTransactionBody) GetFreezeKey() *Key {
	if x != nil {
		return x.FreezeKey
	}
	return nil
}

func (x *TokenCreateTransactionBody) GetWipeKey() *Key {
	if x != nil {
		return x.WipeKey
	}
	return nil
}

func (x *TokenCreateTransactionBody) GetSupplyKey() *Key {
	if x != nil {
		return x.SupplyKey
	}
	return nil
}

func (x *TokenCreateTransactionBody) GetFreezeDefault() bool {
	if x != nil {
		return x.FreezeDefault
	}
	return false
}

func (x *TokenCreateTransactionBody) GetExpiry() *Timestamp {
	if x != nil {
		return x.Expiry
	}
	return nil
}

func (x *TokenCreateTransactionBody) GetAutoRenewAccount() *AccountID {
	if x != nil {
		return x.AutoRenewAccount
	}
	return nil
}

func (x *TokenCreateTransactionBody) GetAutoRenewPeriod() *Duration {
	if x != nil {
		return x.AutoRenewPeriod
	}
	return nil
}

func (x *TokenCreateTransactionBody) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *TokenCreateTransactionBody) GetTokenType() TokenType {
	if x != nil {
		return x.TokenType
	}
	return TokenType_FUNGIBLE_COMMON
}

func (x *TokenCreateTransactionBody) GetSupplyType() TokenSupplyType {
	if x != nil {
		return x.SupplyType
	}
	return TokenSupplyType_INFINITE
}

func (x *TokenCreateTransactionBody) GetMaxSupply() int64 {
	if x != nil {
		return x.MaxSupply
	}
	return 0
}

func (x *TokenCreateTransactionBody) GetFeeScheduleKey() *Key {
	if x != nil {
		return x.FeeScheduleKey
	}
	return nil
}

func (x *TokenCreateTransactionBody) GetCustomFees() []*CustomFee {
	if x != nil {
		return x.CustomFees
	}
	return nil
}

func (x *TokenCreateTransactionBody) GetPauseKey() *Key {
	if x != nil {
		return x.PauseKey
	}
	return nil
}

func (x *TokenCreateTransactionBody) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *TokenCreateTransactionBody) GetMetadataKey() *Key {
	if x != nil {
		return x.MetadataKey
	}
	return nil
}

var File_token_create_proto protoreflect.FileDescriptor

var file_token_create_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x66, 0x65, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xbe, 0x07, 0x0a, 0x1a, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12,
	0x2c, 0x0a, 0x08, 0x74, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x44, 0x52, 0x08, 0x74, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x12, 0x26, 0x0a,
	0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x08, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x06, 0x6b, 0x79, 0x63, 0x4b, 0x65, 0x79, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65,
	0x79, 0x52, 0x06, 0x6b, 0x79, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x09, 0x66, 0x72, 0x65,
	0x65, 0x7a, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x65,
	0x4b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x07, 0x77, 0x69, 0x70, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79,
	0x52, 0x07, 0x77, 0x69, 0x70, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x09, 0x73, 0x75, 0x70,
	0x70, 0x6c, 0x79, 0x4b, 0x65, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79,
	0x4b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x66, 0x72, 0x65, 0x65,
	0x7a, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x79, 0x12, 0x3c, 0x0a, 0x10, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52,
	0x10, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x39, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x61, 0x75, 0x74,
	0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f,
	0x12, 0x2e, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x36, 0x0a, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x73, 0x75,
	0x70, 0x70, 0x6c, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x53,
	0x75, 0x70, 0x70, 0x6c, 0x79, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x61, 0x78,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x34, 0x0a, 0x10, 0x66, 0x65, 0x65, 0x5f, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x0e, 0x66, 0x65,
	0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x0b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x66, 0x65, 0x65, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x46, 0x65, 0x65, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x46, 0x65, 0x65, 0x73, 0x12,
	0x27, 0x0a, 0x09, 0x70, 0x61, 0x75, 0x73, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x08,
	0x70, 0x61, 0x75, 0x73, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x4b, 0x65, 0x79, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_token_create_proto_rawDescOnce sync.Once
	file_token_create_proto_rawDescData = file_token_create_proto_rawDesc
)

func file_token_create_proto_rawDescGZIP() []byte {
	file_token_create_proto_rawDescOnce.Do(func() {
		file_token_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_token_create_proto_rawDescData)
	})
	return file_token_create_proto_rawDescData
}

var file_token_create_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_token_create_proto_goTypes = []interface{}{
	(*TokenCreateTransactionBody)(nil), // 0: proto.TokenCreateTransactionBody
	(*AccountID)(nil),                  // 1: proto.AccountID
	(*Key)(nil),                        // 2: proto.Key
	(*Timestamp)(nil),                  // 3: proto.Timestamp
	(*Duration)(nil),                   // 4: proto.Duration
	(TokenType)(0),                     // 5: proto.TokenType
	(TokenSupplyType)(0),               // 6: proto.TokenSupplyType
	(*CustomFee)(nil),                  // 7: proto.CustomFee
}
var file_token_create_proto_depIdxs = []int32{
	1,  // 0: proto.TokenCreateTransactionBody.treasury:type_name -> proto.AccountID
	2,  // 1: proto.TokenCreateTransactionBody.adminKey:type_name -> proto.Key
	2,  // 2: proto.TokenCreateTransactionBody.kycKey:type_name -> proto.Key
	2,  // 3: proto.TokenCreateTransactionBody.freezeKey:type_name -> proto.Key
	2,  // 4: proto.TokenCreateTransactionBody.wipeKey:type_name -> proto.Key
	2,  // 5: proto.TokenCreateTransactionBody.supplyKey:type_name -> proto.Key
	3,  // 6: proto.TokenCreateTransactionBody.expiry:type_name -> proto.Timestamp
	1,  // 7: proto.TokenCreateTransactionBody.autoRenewAccount:type_name -> proto.AccountID
	4,  // 8: proto.TokenCreateTransactionBody.autoRenewPeriod:type_name -> proto.Duration
	5,  // 9: proto.TokenCreateTransactionBody.tokenType:type_name -> proto.TokenType
	6,  // 10: proto.TokenCreateTransactionBody.supplyType:type_name -> proto.TokenSupplyType
	2,  // 11: proto.TokenCreateTransactionBody.fee_schedule_key:type_name -> proto.Key
	7,  // 12: proto.TokenCreateTransactionBody.custom_fees:type_name -> proto.CustomFee
	2,  // 13: proto.TokenCreateTransactionBody.pause_key:type_name -> proto.Key
	2,  // 14: proto.TokenCreateTransactionBody.metadata_key:type_name -> proto.Key
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_token_create_proto_init() }
func file_token_create_proto_init() {
	if File_token_create_proto != nil {
		return
	}
	file_duration_proto_init()
	file_basic_types_proto_init()
	file_custom_fees_proto_init()
	file_timestamp_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_token_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenCreateTransactionBody); i {
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
			RawDescriptor: file_token_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_token_create_proto_goTypes,
		DependencyIndexes: file_token_create_proto_depIdxs,
		MessageInfos:      file_token_create_proto_msgTypes,
	}.Build()
	File_token_create_proto = out.File
	file_token_create_proto_rawDesc = nil
	file_token_create_proto_goTypes = nil
	file_token_create_proto_depIdxs = nil
}
