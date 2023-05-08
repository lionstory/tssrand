// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.0
// source: proto/rand.proto

package pb

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

// RandClient is the client API for Rand service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RandClient interface {
	GetRand(ctx context.Context, in *RandRequest, opts ...grpc.CallOption) (*RandReply, error)
}

type randClient struct {
	cc grpc.ClientConnInterface
}

func NewRandClient(cc grpc.ClientConnInterface) RandClient {
	return &randClient{cc}
}

func (c *randClient) GetRand(ctx context.Context, in *RandRequest, opts ...grpc.CallOption) (*RandReply, error) {
	out := new(RandReply)
	err := c.cc.Invoke(ctx, "/tssrand.Rand/GetRand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RandServer is the server API for Rand service.
// All implementations must embed UnimplementedRandServer
// for forward compatibility
type RandServer interface {
	GetRand(context.Context, *RandRequest) (*RandReply, error)
	mustEmbedUnimplementedRandServer()
}

// UnimplementedRandServer must be embedded to have forward compatible implementations.
type UnimplementedRandServer struct {
}

func (UnimplementedRandServer) GetRand(context.Context, *RandRequest) (*RandReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRand not implemented")
}
func (UnimplementedRandServer) mustEmbedUnimplementedRandServer() {}

// UnsafeRandServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RandServer will
// result in compilation errors.
type UnsafeRandServer interface {
	mustEmbedUnimplementedRandServer()
}

func RegisterRandServer(s grpc.ServiceRegistrar, srv RandServer) {
	s.RegisterService(&Rand_ServiceDesc, srv)
}

func _Rand_GetRand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandServer).GetRand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tssrand.Rand/GetRand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandServer).GetRand(ctx, req.(*RandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Rand_ServiceDesc is the grpc.ServiceDesc for Rand service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rand_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tssrand.Rand",
	HandlerType: (*RandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRand",
			Handler:    _Rand_GetRand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rand.proto",
}
