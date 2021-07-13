// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: TokenCreate.proto

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

//
//Create a new token. After the token is created, the Token ID for it is in the receipt.
//The specified Treasury Account is receiving the initial supply of tokens as-well as the tokens from the Token Mint operation once executed. The balance of the treasury account is decreased when the Token Burn operation is executed.
//
//The <tt>initialSupply</tt> is the initial supply of the smallest parts of a token (like a tinybar, not an hbar). These are the smallest units of the token which may be transferred.
//
//The supply can change over time. If the total supply at some moment is <i>S</i> parts of tokens, and the token is using <i>D</i> decimals, then <i>S</i> must be less than or equal to 2<sup>63</sup>-1, which is 9,223,372,036,854,775,807. The number of whole tokens (not parts) will be <i>S / 10<sup>D</sup></i>.
//
//If decimals is 8 or 11, then the number of whole tokens can be at most a few billions or millions, respectively. For example, it could match Bitcoin (21 million whole tokens with 8 decimals) or hbars (50 billion whole tokens with 8 decimals). It could even match Bitcoin with milli-satoshis (21 million whole tokens with 11 decimals).
//
//Note that a created token is <i>immutable</i> if the <tt>adminKey</tt> is omitted. No property of an immutable token can ever change, with the sole exception of its expiry. Anyone can pay to extend the expiry time of an immutable token.
type TokenCreateTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                                              // The publicly visible name of the token, limited to a UTF-8 encoding of length <tt>tokens.maxSymbolUtf8Bytes</tt>.
	Symbol           string          `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`                                          // The publicly visible token symbol, limited to a UTF-8 encoding of length <tt>tokens.maxTokenNameUtf8Bytes</tt>.
	Decimals         uint32          `protobuf:"varint,3,opt,name=decimals,proto3" json:"decimals,omitempty"`                                     // For tokens of type FUNGIBLE_COMMON - the number of decimal places a token is divisible by. For tokens of type NON_FUNGIBLE_UNIQUE - value must be 0
	InitialSupply    uint64          `protobuf:"varint,4,opt,name=initialSupply,proto3" json:"initialSupply,omitempty"`                           // Specifies the initial supply of tokens to be put in circulation. The initial supply is sent to the Treasury Account. The supply is in the lowest denomination possible. In the case for NON_FUNGIBLE_UNIQUE Type the value must be 0
	Treasury         *AccountID      `protobuf:"bytes,5,opt,name=treasury,proto3" json:"treasury,omitempty"`                                      // The account which will act as a treasury for the token. This account will receive the specified initial supply or the newly minted NFTs in the case for NON_FUNGIBLE_UNIQUE Type
	AdminKey         *Key            `protobuf:"bytes,6,opt,name=adminKey,proto3" json:"adminKey,omitempty"`                                      // The key which can perform update/delete operations on the token. If empty, the token can be perceived as immutable (not being able to be updated/deleted)
	KycKey           *Key            `protobuf:"bytes,7,opt,name=kycKey,proto3" json:"kycKey,omitempty"`                                          // The key which can grant or revoke KYC of an account for the token's transactions. If empty, KYC is not required, and KYC grant or revoke operations are not possible.
	FreezeKey        *Key            `protobuf:"bytes,8,opt,name=freezeKey,proto3" json:"freezeKey,omitempty"`                                    // The key which can sign to freeze or unfreeze an account for token transactions. If empty, freezing is not possible
	WipeKey          *Key            `protobuf:"bytes,9,opt,name=wipeKey,proto3" json:"wipeKey,omitempty"`                                        // The key which can wipe the token balance of an account. If empty, wipe is not possible
	SupplyKey        *Key            `protobuf:"bytes,10,opt,name=supplyKey,proto3" json:"supplyKey,omitempty"`                                   // The key which can change the supply of a token. The key is used to sign Token Mint/Burn operations
	FreezeDefault    bool            `protobuf:"varint,11,opt,name=freezeDefault,proto3" json:"freezeDefault,omitempty"`                          // The default Freeze status (frozen or unfrozen) of Hedera accounts relative to this token. If true, an account must be unfrozen before it can receive the token
	Expiry           *Timestamp      `protobuf:"bytes,13,opt,name=expiry,proto3" json:"expiry,omitempty"`                                         // The epoch second at which the token should expire; if an auto-renew account and period are specified, this is coerced to the current epoch second plus the autoRenewPeriod
	AutoRenewAccount *AccountID      `protobuf:"bytes,14,opt,name=autoRenewAccount,proto3" json:"autoRenewAccount,omitempty"`                     // An account which will be automatically charged to renew the token's expiration, at autoRenewPeriod interval
	AutoRenewPeriod  *Duration       `protobuf:"bytes,15,opt,name=autoRenewPeriod,proto3" json:"autoRenewPeriod,omitempty"`                       // The interval at which the auto-renew account will be charged to extend the token's expiry
	Memo             string          `protobuf:"bytes,16,opt,name=memo,proto3" json:"memo,omitempty"`                                             // The memo associated with the token (UTF-8 encoding max 100 bytes)
	TokenType        TokenType       `protobuf:"varint,17,opt,name=tokenType,proto3,enum=proto.TokenType" json:"tokenType,omitempty"`             // IWA compatibility. Specifies the token type. Defaults to FUNGIBLE_COMMON
	SupplyType       TokenSupplyType `protobuf:"varint,18,opt,name=supplyType,proto3,enum=proto.TokenSupplyType" json:"supplyType,omitempty"`     // IWA compatibility. Specified the token supply type. Defaults to INFINITE
	MaxSupply        int64           `protobuf:"varint,19,opt,name=maxSupply,proto3" json:"maxSupply,omitempty"`                                  // IWA Compatibility. Depends on TokenSupplyType. For tokens of type FUNGIBLE_COMMON - the maximum number of tokens that can be in circulation. For tokens of type NON_FUNGIBLE_UNIQUE - the maximum number of NFTs (serial numbers) that can be minted. This field can never be changed!
	FeeScheduleKey   *Key            `protobuf:"bytes,20,opt,name=fee_schedule_key,json=feeScheduleKey,proto3" json:"fee_schedule_key,omitempty"` // The key which can change the token's custom fee schedule; must sign a TokenFeeScheduleUpdate transaction
	CustomFees       []*CustomFee    `protobuf:"bytes,21,rep,name=custom_fees,json=customFees,proto3" json:"custom_fees,omitempty"`               // The custom fees to be assessed during a CryptoTransfer that transfers units of this token
}

func (x *TokenCreateTransactionBody) Reset() {
	*x = TokenCreateTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TokenCreate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenCreateTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenCreateTransactionBody) ProtoMessage() {}

func (x *TokenCreateTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_TokenCreate_proto_msgTypes[0]
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
	return file_TokenCreate_proto_rawDescGZIP(), []int{0}
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

var File_TokenCreate_proto protoreflect.FileDescriptor

var file_TokenCreate_proto_rawDesc = []byte{
	0x0a, 0x11, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x46, 0x65, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xca, 0x06, 0x0a, 0x1a, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65,
	0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x65,
	0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x2c, 0x0a, 0x08,
	0x74, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44,
	0x52, 0x08, 0x74, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x12, 0x26, 0x0a, 0x08, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4b,
	0x65, 0x79, 0x12, 0x22, 0x0a, 0x06, 0x6b, 0x79, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x06,
	0x6b, 0x79, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x65,
	0x4b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x4b, 0x65, 0x79,
	0x12, 0x24, 0x0a, 0x07, 0x77, 0x69, 0x70, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x07, 0x77,
	0x69, 0x70, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x09, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79,
	0x4b, 0x65, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x4b, 0x65, 0x79,
	0x12, 0x24, 0x0a, 0x0d, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x44,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79,
	0x12, 0x3c, 0x0a, 0x10, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x10, 0x61, 0x75,
	0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x39,
	0x0a, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x50, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65,
	0x6e, 0x65, 0x77, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d,
	0x6f, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x2e, 0x0a,
	0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a,
	0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53,
	0x75, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x53, 0x75, 0x70, 0x70,
	0x6c, 0x79, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x53, 0x75, 0x70,
	0x70, 0x6c, 0x79, 0x12, 0x34, 0x0a, 0x10, 0x66, 0x65, 0x65, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x0e, 0x66, 0x65, 0x65, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x0b, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x5f, 0x66, 0x65, 0x65, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x46, 0x65, 0x65,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x46, 0x65, 0x65, 0x73, 0x42, 0x26, 0x0a, 0x22,
	0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61,
	0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TokenCreate_proto_rawDescOnce sync.Once
	file_TokenCreate_proto_rawDescData = file_TokenCreate_proto_rawDesc
)

func file_TokenCreate_proto_rawDescGZIP() []byte {
	file_TokenCreate_proto_rawDescOnce.Do(func() {
		file_TokenCreate_proto_rawDescData = protoimpl.X.CompressGZIP(file_TokenCreate_proto_rawDescData)
	})
	return file_TokenCreate_proto_rawDescData
}

var file_TokenCreate_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TokenCreate_proto_goTypes = []interface{}{
	(*TokenCreateTransactionBody)(nil), // 0: proto.TokenCreateTransactionBody
	(*AccountID)(nil),                  // 1: proto.AccountID
	(*Key)(nil),                        // 2: proto.Key
	(*Timestamp)(nil),                  // 3: proto.Timestamp
	(*Duration)(nil),                   // 4: proto.Duration
	(TokenType)(0),                     // 5: proto.TokenType
	(TokenSupplyType)(0),               // 6: proto.TokenSupplyType
	(*CustomFee)(nil),                  // 7: proto.CustomFee
}
var file_TokenCreate_proto_depIdxs = []int32{
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
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_TokenCreate_proto_init() }
func file_TokenCreate_proto_init() {
	if File_TokenCreate_proto != nil {
		return
	}
	file_Duration_proto_init()
	file_BasicTypes_proto_init()
	file_CustomFees_proto_init()
	file_Timestamp_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TokenCreate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_TokenCreate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TokenCreate_proto_goTypes,
		DependencyIndexes: file_TokenCreate_proto_depIdxs,
		MessageInfos:      file_TokenCreate_proto_msgTypes,
	}.Build()
	File_TokenCreate_proto = out.File
	file_TokenCreate_proto_rawDesc = nil
	file_TokenCreate_proto_goTypes = nil
	file_TokenCreate_proto_depIdxs = nil
}
