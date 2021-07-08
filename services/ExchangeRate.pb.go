// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: ExchangeRate.proto

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
//An exchange rate between hbar and cents (USD) and the time at which the exchange rate will expire, and be superseded by a new exchange rate.
type ExchangeRate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Denominator in calculation of exchange rate between hbar and cents
	HbarEquiv int32 `protobuf:"varint,1,opt,name=hbarEquiv,proto3" json:"hbarEquiv,omitempty"`
	// Numerator in calculation of exchange rate between hbar and cents
	CentEquiv int32 `protobuf:"varint,2,opt,name=centEquiv,proto3" json:"centEquiv,omitempty"`
	// Expiration time in seconds for this exchange rate
	ExpirationTime *TimestampSeconds `protobuf:"bytes,3,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`
}

func (x *ExchangeRate) Reset() {
	*x = ExchangeRate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ExchangeRate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeRate) ProtoMessage() {}

func (x *ExchangeRate) ProtoReflect() protoreflect.Message {
	mi := &file_ExchangeRate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeRate.ProtoReflect.Descriptor instead.
func (*ExchangeRate) Descriptor() ([]byte, []int) {
	return file_ExchangeRate_proto_rawDescGZIP(), []int{0}
}

func (x *ExchangeRate) GetHbarEquiv() int32 {
	if x != nil {
		return x.HbarEquiv
	}
	return 0
}

func (x *ExchangeRate) GetCentEquiv() int32 {
	if x != nil {
		return x.CentEquiv
	}
	return 0
}

func (x *ExchangeRate) GetExpirationTime() *TimestampSeconds {
	if x != nil {
		return x.ExpirationTime
	}
	return nil
}

// Two sets of exchange rates
type ExchangeRateSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Current exchange rate
	CurrentRate *ExchangeRate `protobuf:"bytes,1,opt,name=currentRate,proto3" json:"currentRate,omitempty"`
	// Next exchange rate which will take effect when current rate expires
	NextRate *ExchangeRate `protobuf:"bytes,2,opt,name=nextRate,proto3" json:"nextRate,omitempty"`
}

func (x *ExchangeRateSet) Reset() {
	*x = ExchangeRateSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ExchangeRate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeRateSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeRateSet) ProtoMessage() {}

func (x *ExchangeRateSet) ProtoReflect() protoreflect.Message {
	mi := &file_ExchangeRate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeRateSet.ProtoReflect.Descriptor instead.
func (*ExchangeRateSet) Descriptor() ([]byte, []int) {
	return file_ExchangeRate_proto_rawDescGZIP(), []int{1}
}

func (x *ExchangeRateSet) GetCurrentRate() *ExchangeRate {
	if x != nil {
		return x.CurrentRate
	}
	return nil
}

func (x *ExchangeRateSet) GetNextRate() *ExchangeRate {
	if x != nil {
		return x.NextRate
	}
	return nil
}

var File_ExchangeRate_proto protoreflect.FileDescriptor

var file_ExchangeRate_proto_rawDesc = []byte{
	0x0a, 0x12, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a,
	0x0c, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x68, 0x62, 0x61, 0x72, 0x45, 0x71, 0x75, 0x69, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x68, 0x62, 0x61, 0x72, 0x45, 0x71, 0x75, 0x69, 0x76, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x65, 0x6e, 0x74, 0x45, 0x71, 0x75, 0x69, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x63, 0x65, 0x6e, 0x74, 0x45, 0x71, 0x75, 0x69, 0x76, 0x12, 0x3f, 0x0a, 0x0e, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x79, 0x0a, 0x0f, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x12, 0x35, 0x0a,
	0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x52, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x52, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x08, 0x6e, 0x65, 0x78,
	0x74, 0x52, 0x61, 0x74, 0x65, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ExchangeRate_proto_rawDescOnce sync.Once
	file_ExchangeRate_proto_rawDescData = file_ExchangeRate_proto_rawDesc
)

func file_ExchangeRate_proto_rawDescGZIP() []byte {
	file_ExchangeRate_proto_rawDescOnce.Do(func() {
		file_ExchangeRate_proto_rawDescData = protoimpl.X.CompressGZIP(file_ExchangeRate_proto_rawDescData)
	})
	return file_ExchangeRate_proto_rawDescData
}

var file_ExchangeRate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ExchangeRate_proto_goTypes = []interface{}{
	(*ExchangeRate)(nil),     // 0: proto.ExchangeRate
	(*ExchangeRateSet)(nil),  // 1: proto.ExchangeRateSet
	(*TimestampSeconds)(nil), // 2: proto.TimestampSeconds
}
var file_ExchangeRate_proto_depIdxs = []int32{
	2, // 0: proto.ExchangeRate.expirationTime:type_name -> proto.TimestampSeconds
	0, // 1: proto.ExchangeRateSet.currentRate:type_name -> proto.ExchangeRate
	0, // 2: proto.ExchangeRateSet.nextRate:type_name -> proto.ExchangeRate
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ExchangeRate_proto_init() }
func file_ExchangeRate_proto_init() {
	if File_ExchangeRate_proto != nil {
		return
	}
	file_Timestamp_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ExchangeRate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeRate); i {
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
		file_ExchangeRate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeRateSet); i {
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
			RawDescriptor: file_ExchangeRate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ExchangeRate_proto_goTypes,
		DependencyIndexes: file_ExchangeRate_proto_depIdxs,
		MessageInfos:      file_ExchangeRate_proto_msgTypes,
	}.Build()
	File_ExchangeRate_proto = out.File
	file_ExchangeRate_proto_rawDesc = nil
	file_ExchangeRate_proto_goTypes = nil
	file_ExchangeRate_proto_depIdxs = nil
}
