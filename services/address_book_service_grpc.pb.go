// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: address_book_service.proto

package services

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AddressBookServiceClient is the client API for AddressBookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AddressBookServiceClient interface {
	// *
	// A transaction to create a new consensus node in the network.
	// address book.
	// <p>
	// This transaction, once complete, SHALL add a new consensus node to the
	// network state.<br/>
	// The new consensus node SHALL remain in state, but SHALL NOT participate
	// in network consensus until the network updates the network configuration.
	// <p>
	// Hedera governing council authorization is REQUIRED for this transaction.
	CreateNode(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// *
	// A transaction to remove a consensus node from the network address
	// book.
	// <p>
	// This transaction, once complete, SHALL remove the identified consensus
	// node from the network state.
	// <p>
	// Hedera governing council authorization is REQUIRED for this transaction.
	DeleteNode(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// *
	// A transaction to update an existing consensus node from the network
	// address book.
	// <p>
	// This transaction, once complete, SHALL modify the identified consensus
	// node state as requested.
	// <p>
	// This transaction MAY be authorized by either the node operator OR the
	// Hedera governing council.
	UpdateNode(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type addressBookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAddressBookServiceClient(cc grpc.ClientConnInterface) AddressBookServiceClient {
	return &addressBookServiceClient{cc}
}

func (c *addressBookServiceClient) CreateNode(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.AddressBookService/createNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressBookServiceClient) DeleteNode(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.AddressBookService/deleteNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressBookServiceClient) UpdateNode(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.AddressBookService/updateNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressBookServiceServer is the server API for AddressBookService service.
// All implementations must embed UnimplementedAddressBookServiceServer
// for forward compatibility
type AddressBookServiceServer interface {
	// *
	// A transaction to create a new consensus node in the network.
	// address book.
	// <p>
	// This transaction, once complete, SHALL add a new consensus node to the
	// network state.<br/>
	// The new consensus node SHALL remain in state, but SHALL NOT participate
	// in network consensus until the network updates the network configuration.
	// <p>
	// Hedera governing council authorization is REQUIRED for this transaction.
	CreateNode(context.Context, *Transaction) (*TransactionResponse, error)
	// *
	// A transaction to remove a consensus node from the network address
	// book.
	// <p>
	// This transaction, once complete, SHALL remove the identified consensus
	// node from the network state.
	// <p>
	// Hedera governing council authorization is REQUIRED for this transaction.
	DeleteNode(context.Context, *Transaction) (*TransactionResponse, error)
	// *
	// A transaction to update an existing consensus node from the network
	// address book.
	// <p>
	// This transaction, once complete, SHALL modify the identified consensus
	// node state as requested.
	// <p>
	// This transaction MAY be authorized by either the node operator OR the
	// Hedera governing council.
	UpdateNode(context.Context, *Transaction) (*TransactionResponse, error)
	mustEmbedUnimplementedAddressBookServiceServer()
}

// UnimplementedAddressBookServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAddressBookServiceServer struct {
}

func (UnimplementedAddressBookServiceServer) CreateNode(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNode not implemented")
}
func (UnimplementedAddressBookServiceServer) DeleteNode(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNode not implemented")
}
func (UnimplementedAddressBookServiceServer) UpdateNode(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNode not implemented")
}
func (UnimplementedAddressBookServiceServer) mustEmbedUnimplementedAddressBookServiceServer() {}

// UnsafeAddressBookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AddressBookServiceServer will
// result in compilation errors.
type UnsafeAddressBookServiceServer interface {
	mustEmbedUnimplementedAddressBookServiceServer()
}

func RegisterAddressBookServiceServer(s grpc.ServiceRegistrar, srv AddressBookServiceServer) {
	s.RegisterService(&AddressBookService_ServiceDesc, srv)
}

func _AddressBookService_CreateNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressBookServiceServer).CreateNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddressBookService/createNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressBookServiceServer).CreateNode(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressBookService_DeleteNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressBookServiceServer).DeleteNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddressBookService/deleteNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressBookServiceServer).DeleteNode(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressBookService_UpdateNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressBookServiceServer).UpdateNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddressBookService/updateNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressBookServiceServer).UpdateNode(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

// AddressBookService_ServiceDesc is the grpc.ServiceDesc for AddressBookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AddressBookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AddressBookService",
	HandlerType: (*AddressBookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createNode",
			Handler:    _AddressBookService_CreateNode_Handler,
		},
		{
			MethodName: "deleteNode",
			Handler:    _AddressBookService_DeleteNode_Handler,
		},
		{
			MethodName: "updateNode",
			Handler:    _AddressBookService_UpdateNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "address_book_service.proto",
}