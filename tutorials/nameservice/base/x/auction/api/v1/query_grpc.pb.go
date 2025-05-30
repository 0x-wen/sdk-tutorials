// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cosmos/auction/v1/query.proto

package auctionv1

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

const (
	Query_Name_FullMethodName  = "/cosmos.auction.v1.Query/Name"
	Query_Names_FullMethodName = "/cosmos.auction.v1.Query/Names"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Name is a method that takes a QueryNameRequest and returns a QueryNameResponse.
	Name(ctx context.Context, in *QueryNameRequest, opts ...grpc.CallOption) (*QueryNameResponse, error)
	Names(ctx context.Context, in *QueryNamesRequest, opts ...grpc.CallOption) (*QueryNamesResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Name(ctx context.Context, in *QueryNameRequest, opts ...grpc.CallOption) (*QueryNameResponse, error) {
	out := new(QueryNameResponse)
	err := c.cc.Invoke(ctx, Query_Name_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Names(ctx context.Context, in *QueryNamesRequest, opts ...grpc.CallOption) (*QueryNamesResponse, error) {
	out := new(QueryNamesResponse)
	err := c.cc.Invoke(ctx, Query_Names_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Name is a method that takes a QueryNameRequest and returns a QueryNameResponse.
	Name(context.Context, *QueryNameRequest) (*QueryNameResponse, error)
	Names(context.Context, *QueryNamesRequest) (*QueryNamesResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Name(context.Context, *QueryNameRequest) (*QueryNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Name not implemented")
}
func (UnimplementedQueryServer) Names(context.Context, *QueryNamesRequest) (*QueryNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Names not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Name_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Name(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Name_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Name(ctx, req.(*QueryNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Names_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Names(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Names_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Names(ctx, req.(*QueryNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.auction.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Name",
			Handler:    _Query_Name_Handler,
		},
		{
			MethodName: "Names",
			Handler:    _Query_Names_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/auction/v1/query.proto",
}
