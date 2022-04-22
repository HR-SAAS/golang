// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: counter.proto

package proto

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

// RecruitCounterServiceClient is the client API for RecruitCounterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecruitCounterServiceClient interface {
	// 统计简历
	CountPost(ctx context.Context, in *CountPostRequest, opts ...grpc.CallOption) (*CountPostResponse, error)
	CountUserPost(ctx context.Context, in *CountUserPostRequest, opts ...grpc.CallOption) (*CountUserPostResponse, error)
}

type recruitCounterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecruitCounterServiceClient(cc grpc.ClientConnInterface) RecruitCounterServiceClient {
	return &recruitCounterServiceClient{cc}
}

func (c *recruitCounterServiceClient) CountPost(ctx context.Context, in *CountPostRequest, opts ...grpc.CallOption) (*CountPostResponse, error) {
	out := new(CountPostResponse)
	err := c.cc.Invoke(ctx, "/RecruitCounterService/CountPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recruitCounterServiceClient) CountUserPost(ctx context.Context, in *CountUserPostRequest, opts ...grpc.CallOption) (*CountUserPostResponse, error) {
	out := new(CountUserPostResponse)
	err := c.cc.Invoke(ctx, "/RecruitCounterService/CountUserPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecruitCounterServiceServer is the server API for RecruitCounterService service.
// All implementations must embed UnimplementedRecruitCounterServiceServer
// for forward compatibility
type RecruitCounterServiceServer interface {
	// 统计简历
	CountPost(context.Context, *CountPostRequest) (*CountPostResponse, error)
	CountUserPost(context.Context, *CountUserPostRequest) (*CountUserPostResponse, error)
	mustEmbedUnimplementedRecruitCounterServiceServer()
}

// UnimplementedRecruitCounterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRecruitCounterServiceServer struct {
}

func (UnimplementedRecruitCounterServiceServer) CountPost(context.Context, *CountPostRequest) (*CountPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountPost not implemented")
}
func (UnimplementedRecruitCounterServiceServer) CountUserPost(context.Context, *CountUserPostRequest) (*CountUserPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountUserPost not implemented")
}
func (UnimplementedRecruitCounterServiceServer) mustEmbedUnimplementedRecruitCounterServiceServer() {}

// UnsafeRecruitCounterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecruitCounterServiceServer will
// result in compilation errors.
type UnsafeRecruitCounterServiceServer interface {
	mustEmbedUnimplementedRecruitCounterServiceServer()
}

func RegisterRecruitCounterServiceServer(s grpc.ServiceRegistrar, srv RecruitCounterServiceServer) {
	s.RegisterService(&RecruitCounterService_ServiceDesc, srv)
}

func _RecruitCounterService_CountPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruitCounterServiceServer).CountPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecruitCounterService/CountPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruitCounterServiceServer).CountPost(ctx, req.(*CountPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecruitCounterService_CountUserPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountUserPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecruitCounterServiceServer).CountUserPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecruitCounterService/CountUserPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecruitCounterServiceServer).CountUserPost(ctx, req.(*CountUserPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RecruitCounterService_ServiceDesc is the grpc.ServiceDesc for RecruitCounterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecruitCounterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RecruitCounterService",
	HandlerType: (*RecruitCounterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CountPost",
			Handler:    _RecruitCounterService_CountPost_Handler,
		},
		{
			MethodName: "CountUserPost",
			Handler:    _RecruitCounterService_CountUserPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "counter.proto",
}
