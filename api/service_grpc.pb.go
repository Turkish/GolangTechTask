// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// VotingServiceClient is the client API for VotingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VotingServiceClient interface {
	CreateVoteable(ctx context.Context, in *CreateVoteableRequest, opts ...grpc.CallOption) (*CreateVoteableResponse, error)
	ListVoteables(ctx context.Context, in *ListVoteableRequest, opts ...grpc.CallOption) (*ListVoteableResponse, error)
	CastVote(ctx context.Context, in *CastVoteRequest, opts ...grpc.CallOption) (*CastVoteResponse, error)
}

type votingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVotingServiceClient(cc grpc.ClientConnInterface) VotingServiceClient {
	return &votingServiceClient{cc}
}

func (c *votingServiceClient) CreateVoteable(ctx context.Context, in *CreateVoteableRequest, opts ...grpc.CallOption) (*CreateVoteableResponse, error) {
	out := new(CreateVoteableResponse)
	err := c.cc.Invoke(ctx, "/VotingService/CreateVoteable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *votingServiceClient) ListVoteables(ctx context.Context, in *ListVoteableRequest, opts ...grpc.CallOption) (*ListVoteableResponse, error) {
	out := new(ListVoteableResponse)
	err := c.cc.Invoke(ctx, "/VotingService/ListVoteables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *votingServiceClient) CastVote(ctx context.Context, in *CastVoteRequest, opts ...grpc.CallOption) (*CastVoteResponse, error) {
	out := new(CastVoteResponse)
	err := c.cc.Invoke(ctx, "/VotingService/CastVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VotingServiceServer is the server API for VotingService service.
// All implementations must embed UnimplementedVotingServiceServer
// for forward compatibility
type VotingServiceServer interface {
	CreateVoteable(context.Context, *CreateVoteableRequest) (*CreateVoteableResponse, error)
	ListVoteables(context.Context, *ListVoteableRequest) (*ListVoteableResponse, error)
	CastVote(context.Context, *CastVoteRequest) (*CastVoteResponse, error)
	mustEmbedUnimplementedVotingServiceServer()
}

// UnimplementedVotingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVotingServiceServer struct {
}

func (UnimplementedVotingServiceServer) CreateVoteable(context.Context, *CreateVoteableRequest) (*CreateVoteableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVoteable not implemented")
}
func (UnimplementedVotingServiceServer) ListVoteables(context.Context, *ListVoteableRequest) (*ListVoteableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVoteables not implemented")
}
func (UnimplementedVotingServiceServer) CastVote(context.Context, *CastVoteRequest) (*CastVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CastVote not implemented")
}
func (UnimplementedVotingServiceServer) mustEmbedUnimplementedVotingServiceServer() {}

// UnsafeVotingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VotingServiceServer will
// result in compilation errors.
type UnsafeVotingServiceServer interface {
	mustEmbedUnimplementedVotingServiceServer()
}

func RegisterVotingServiceServer(s grpc.ServiceRegistrar, srv VotingServiceServer) {
	s.RegisterService(&VotingService_ServiceDesc, srv)
}

func _VotingService_CreateVoteable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVoteableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VotingServiceServer).CreateVoteable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VotingService/CreateVoteable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VotingServiceServer).CreateVoteable(ctx, req.(*CreateVoteableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VotingService_ListVoteables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVoteableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VotingServiceServer).ListVoteables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VotingService/ListVoteables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VotingServiceServer).ListVoteables(ctx, req.(*ListVoteableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VotingService_CastVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CastVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VotingServiceServer).CastVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VotingService/CastVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VotingServiceServer).CastVote(ctx, req.(*CastVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VotingService_ServiceDesc is the grpc.ServiceDesc for VotingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VotingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "VotingService",
	HandlerType: (*VotingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVoteable",
			Handler:    _VotingService_CreateVoteable_Handler,
		},
		{
			MethodName: "ListVoteables",
			Handler:    _VotingService_ListVoteables_Handler,
		},
		{
			MethodName: "CastVote",
			Handler:    _VotingService_CastVote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/service.proto",
}
