// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: node_create.proto

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
// A transaction body to add a new node to the network.
// After the node is created, the node_id for it is in the receipt.
//
// This transaction body SHALL be considered a "privileged transaction".
//
// This message supports a transaction to create a new node in the network.
// The transaction, once complete, enables a new consensus node
// to join the network, and requires governing council authorization.
//
// A `NodeCreateTransactionBody` MUST be signed by the governing council.<br/>
// The newly created node information will be used to generate config.txt and
// a-pulbic-NodeAlias.pem file per each node during phase 2,<br>
// not active until next freeze upgrade.
type NodeCreateTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// Node account id, mandatory field, ALIAS is not allowed, only ACCOUNT_NUM.
	// If account_id does not exist, it will reject the transaction.
	// Multiple nodes can have the same account_id.
	AccountId *AccountID `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// *
	// Description of the node with UTF-8 encoding up to 100 bytes, optional field.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// *
	// Ip address and port, mandatory field. Fully qualified domain name is
	// not allowed here. Maximum number of these endpoints is 10.
	// The first in the list is used as the Internal IP address in config.txt,
	// the second in the list is used as the External IP address in config.txt,
	// the rest of IP addresses are ignored for DAB phase 2.
	GossipEndpoint []*ServiceEndpoint `protobuf:"bytes,3,rep,name=gossip_endpoint,json=gossipEndpoint,proto3" json:"gossip_endpoint,omitempty"`
	// *
	// A node's grpc service IP addresses and ports, IP:Port is mandatory,
	// fully qualified domain name is optional. Maximum number of these endpoints is 8.
	ServiceEndpoint []*ServiceEndpoint `protobuf:"bytes,4,rep,name=service_endpoint,json=serviceEndpoint,proto3" json:"service_endpoint,omitempty"`
	// *
	// The node's X509 certificate used to sign stream files (e.g., record stream
	// files). Precisely, this field is the DER encoding of gossip X509 certificate.
	// This is a mandatory field.
	GossipCaCertificate []byte `protobuf:"bytes,5,opt,name=gossip_ca_certificate,json=gossipCaCertificate,proto3" json:"gossip_ca_certificate,omitempty"`
	// *
	// Hash of the node's TLS certificate. Precisely, this field is a string of
	// hexadecimal characters which translated to binary, are the SHA-384 hash of
	// the UTF-8 NFKD encoding of the node's TLS cert in PEM format.
	// Its value can be used to verify the node's certificate it presents
	// during TLS negotiations.node x509 certificate hash, optional field.
	GrpcCertificateHash []byte `protobuf:"bytes,6,opt,name=grpc_certificate_hash,json=grpcCertificateHash,proto3" json:"grpc_certificate_hash,omitempty"`
}

func (x *NodeCreateTransactionBody) Reset() {
	*x = NodeCreateTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeCreateTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeCreateTransactionBody) ProtoMessage() {}

func (x *NodeCreateTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_node_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeCreateTransactionBody.ProtoReflect.Descriptor instead.
func (*NodeCreateTransactionBody) Descriptor() ([]byte, []int) {
	return file_node_create_proto_rawDescGZIP(), []int{0}
}

func (x *NodeCreateTransactionBody) GetAccountId() *AccountID {
	if x != nil {
		return x.AccountId
	}
	return nil
}

func (x *NodeCreateTransactionBody) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NodeCreateTransactionBody) GetGossipEndpoint() []*ServiceEndpoint {
	if x != nil {
		return x.GossipEndpoint
	}
	return nil
}

func (x *NodeCreateTransactionBody) GetServiceEndpoint() []*ServiceEndpoint {
	if x != nil {
		return x.ServiceEndpoint
	}
	return nil
}

func (x *NodeCreateTransactionBody) GetGossipCaCertificate() []byte {
	if x != nil {
		return x.GossipCaCertificate
	}
	return nil
}

func (x *NodeCreateTransactionBody) GetGrpcCertificateHash() []byte {
	if x != nil {
		return x.GrpcCertificateHash
	}
	return nil
}

var File_node_create_proto protoreflect.FileDescriptor

var file_node_create_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x02,
	0x0a, 0x19, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x2f, 0x0a, 0x0a, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x44, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f,
	0x0a, 0x0f, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x0e, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x41, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x5f, 0x63, 0x61, 0x5f,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x13, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x43, 0x61, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x13, 0x67, 0x72, 0x70, 0x63, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x48, 0x61, 0x73, 0x68, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f,
	0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61,
	0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_create_proto_rawDescOnce sync.Once
	file_node_create_proto_rawDescData = file_node_create_proto_rawDesc
)

func file_node_create_proto_rawDescGZIP() []byte {
	file_node_create_proto_rawDescOnce.Do(func() {
		file_node_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_create_proto_rawDescData)
	})
	return file_node_create_proto_rawDescData
}

var file_node_create_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_node_create_proto_goTypes = []interface{}{
	(*NodeCreateTransactionBody)(nil), // 0: proto.NodeCreateTransactionBody
	(*AccountID)(nil),                 // 1: proto.AccountID
	(*ServiceEndpoint)(nil),           // 2: proto.ServiceEndpoint
}
var file_node_create_proto_depIdxs = []int32{
	1, // 0: proto.NodeCreateTransactionBody.account_id:type_name -> proto.AccountID
	2, // 1: proto.NodeCreateTransactionBody.gossip_endpoint:type_name -> proto.ServiceEndpoint
	2, // 2: proto.NodeCreateTransactionBody.service_endpoint:type_name -> proto.ServiceEndpoint
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_node_create_proto_init() }
func file_node_create_proto_init() {
	if File_node_create_proto != nil {
		return
	}
	file_basic_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_node_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeCreateTransactionBody); i {
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
			RawDescriptor: file_node_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_node_create_proto_goTypes,
		DependencyIndexes: file_node_create_proto_depIdxs,
		MessageInfos:      file_node_create_proto_msgTypes,
	}.Build()
	File_node_create_proto = out.File
	file_node_create_proto_rawDesc = nil
	file_node_create_proto_goTypes = nil
	file_node_create_proto_depIdxs = nil
}
