// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: mirror/consensus_service.proto

package mirror

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

// ConsensusServiceClient is the client API for ConsensusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConsensusServiceClient interface {
	SubscribeTopic(ctx context.Context, in *ConsensusTopicQuery, opts ...grpc.CallOption) (ConsensusService_SubscribeTopicClient, error)
}

type consensusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConsensusServiceClient(cc grpc.ClientConnInterface) ConsensusServiceClient {
	return &consensusServiceClient{cc}
}

func (c *consensusServiceClient) SubscribeTopic(ctx context.Context, in *ConsensusTopicQuery, opts ...grpc.CallOption) (ConsensusService_SubscribeTopicClient, error) {
	stream, err := c.cc.NewStream(ctx, &ConsensusService_ServiceDesc.Streams[0], "/com.hedera.mirror.api.proto.ConsensusService/subscribeTopic", opts...)
	if err != nil {
		return nil, err
	}
	x := &consensusServiceSubscribeTopicClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ConsensusService_SubscribeTopicClient interface {
	Recv() (*ConsensusTopicResponse, error)
	grpc.ClientStream
}

type consensusServiceSubscribeTopicClient struct {
	grpc.ClientStream
}

func (x *consensusServiceSubscribeTopicClient) Recv() (*ConsensusTopicResponse, error) {
	m := new(ConsensusTopicResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConsensusServiceServer is the server API for ConsensusService service.
// All implementations must embed UnimplementedConsensusServiceServer
// for forward compatibility
type ConsensusServiceServer interface {
	SubscribeTopic(*ConsensusTopicQuery, ConsensusService_SubscribeTopicServer) error
	mustEmbedUnimplementedConsensusServiceServer()
}

// UnimplementedConsensusServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConsensusServiceServer struct {
}

func (UnimplementedConsensusServiceServer) SubscribeTopic(*ConsensusTopicQuery, ConsensusService_SubscribeTopicServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeTopic not implemented")
}
func (UnimplementedConsensusServiceServer) mustEmbedUnimplementedConsensusServiceServer() {}

// UnsafeConsensusServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConsensusServiceServer will
// result in compilation errors.
type UnsafeConsensusServiceServer interface {
	mustEmbedUnimplementedConsensusServiceServer()
}

func RegisterConsensusServiceServer(s grpc.ServiceRegistrar, srv ConsensusServiceServer) {
	s.RegisterService(&ConsensusService_ServiceDesc, srv)
}

func _ConsensusService_SubscribeTopic_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConsensusTopicQuery)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConsensusServiceServer).SubscribeTopic(m, &consensusServiceSubscribeTopicServer{stream})
}

type ConsensusService_SubscribeTopicServer interface {
	Send(*ConsensusTopicResponse) error
	grpc.ServerStream
}

type consensusServiceSubscribeTopicServer struct {
	grpc.ServerStream
}

func (x *consensusServiceSubscribeTopicServer) Send(m *ConsensusTopicResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ConsensusService_ServiceDesc is the grpc.ServiceDesc for ConsensusService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConsensusService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.hedera.mirror.api.proto.ConsensusService",
	HandlerType: (*ConsensusServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "subscribeTopic",
			Handler:       _ConsensusService_SubscribeTopic_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "mirror/consensus_service.proto",
}
