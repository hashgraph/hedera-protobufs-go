// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.9
// source: token_service.proto

package services

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_token_service_proto protoreflect.FileDescriptor

var file_token_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd0, 0x09, 0x0a, 0x0c, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x62, 0x75, 0x72, 0x6e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x42, 0x0a, 0x10, 0x77, 0x69, 0x70, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x12, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x14, 0x75, 0x6e,
	0x66, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x48, 0x0a, 0x16, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x4b, 0x79, 0x63, 0x54, 0x6f,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x19,
	0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x4b, 0x79, 0x63, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0f, 0x61, 0x73, 0x73,
	0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x10,
	0x64, 0x69, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x48, 0x0a, 0x16, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x46,
	0x65, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x0c, 0x67, 0x65,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x12, 0x67, 0x65, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03,
	0x88, 0x02, 0x01, 0x12, 0x30, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4e,
	0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x88, 0x02, 0x01, 0x12, 0x3c, 0x0a,
	0x0a, 0x70, 0x61, 0x75, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x75,
	0x6e, 0x70, 0x61, 0x75, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x28, 0x0a, 0x26, 0x63,
	0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x6a, 0x61, 0x76, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_token_service_proto_goTypes = []interface{}{
	(*Transaction)(nil),         // 0: proto.Transaction
	(*Query)(nil),               // 1: proto.Query
	(*TransactionResponse)(nil), // 2: proto.TransactionResponse
	(*Response)(nil),            // 3: proto.Response
}
var file_token_service_proto_depIdxs = []int32{
	0,  // 0: proto.TokenService.createToken:input_type -> proto.Transaction
	0,  // 1: proto.TokenService.updateToken:input_type -> proto.Transaction
	0,  // 2: proto.TokenService.mintToken:input_type -> proto.Transaction
	0,  // 3: proto.TokenService.burnToken:input_type -> proto.Transaction
	0,  // 4: proto.TokenService.deleteToken:input_type -> proto.Transaction
	0,  // 5: proto.TokenService.wipeTokenAccount:input_type -> proto.Transaction
	0,  // 6: proto.TokenService.freezeTokenAccount:input_type -> proto.Transaction
	0,  // 7: proto.TokenService.unfreezeTokenAccount:input_type -> proto.Transaction
	0,  // 8: proto.TokenService.grantKycToTokenAccount:input_type -> proto.Transaction
	0,  // 9: proto.TokenService.revokeKycFromTokenAccount:input_type -> proto.Transaction
	0,  // 10: proto.TokenService.associateTokens:input_type -> proto.Transaction
	0,  // 11: proto.TokenService.dissociateTokens:input_type -> proto.Transaction
	0,  // 12: proto.TokenService.updateTokenFeeSchedule:input_type -> proto.Transaction
	1,  // 13: proto.TokenService.getTokenInfo:input_type -> proto.Query
	1,  // 14: proto.TokenService.getAccountNftInfos:input_type -> proto.Query
	1,  // 15: proto.TokenService.getTokenNftInfo:input_type -> proto.Query
	1,  // 16: proto.TokenService.getTokenNftInfos:input_type -> proto.Query
	0,  // 17: proto.TokenService.pauseToken:input_type -> proto.Transaction
	0,  // 18: proto.TokenService.unpauseToken:input_type -> proto.Transaction
	2,  // 19: proto.TokenService.createToken:output_type -> proto.TransactionResponse
	2,  // 20: proto.TokenService.updateToken:output_type -> proto.TransactionResponse
	2,  // 21: proto.TokenService.mintToken:output_type -> proto.TransactionResponse
	2,  // 22: proto.TokenService.burnToken:output_type -> proto.TransactionResponse
	2,  // 23: proto.TokenService.deleteToken:output_type -> proto.TransactionResponse
	2,  // 24: proto.TokenService.wipeTokenAccount:output_type -> proto.TransactionResponse
	2,  // 25: proto.TokenService.freezeTokenAccount:output_type -> proto.TransactionResponse
	2,  // 26: proto.TokenService.unfreezeTokenAccount:output_type -> proto.TransactionResponse
	2,  // 27: proto.TokenService.grantKycToTokenAccount:output_type -> proto.TransactionResponse
	2,  // 28: proto.TokenService.revokeKycFromTokenAccount:output_type -> proto.TransactionResponse
	2,  // 29: proto.TokenService.associateTokens:output_type -> proto.TransactionResponse
	2,  // 30: proto.TokenService.dissociateTokens:output_type -> proto.TransactionResponse
	2,  // 31: proto.TokenService.updateTokenFeeSchedule:output_type -> proto.TransactionResponse
	3,  // 32: proto.TokenService.getTokenInfo:output_type -> proto.Response
	3,  // 33: proto.TokenService.getAccountNftInfos:output_type -> proto.Response
	3,  // 34: proto.TokenService.getTokenNftInfo:output_type -> proto.Response
	3,  // 35: proto.TokenService.getTokenNftInfos:output_type -> proto.Response
	2,  // 36: proto.TokenService.pauseToken:output_type -> proto.TransactionResponse
	2,  // 37: proto.TokenService.unpauseToken:output_type -> proto.TransactionResponse
	19, // [19:38] is the sub-list for method output_type
	0,  // [0:19] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_token_service_proto_init() }
func file_token_service_proto_init() {
	if File_token_service_proto != nil {
		return
	}
	file_query_proto_init()
	file_response_proto_init()
	file_transaction_response_proto_init()
	file_transaction_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_token_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_token_service_proto_goTypes,
		DependencyIndexes: file_token_service_proto_depIdxs,
	}.Build()
	File_token_service_proto = out.File
	file_token_service_proto_rawDesc = nil
	file_token_service_proto_goTypes = nil
	file_token_service_proto_depIdxs = nil
}
