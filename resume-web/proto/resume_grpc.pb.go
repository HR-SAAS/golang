// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: resume.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ResumeClient is the client API for Resume service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResumeClient interface {
	// 获取简历列表
	GetResumeList(ctx context.Context, in *GetResumeListRequest, opts ...grpc.CallOption) (*ResumeListResponse, error)
	// 创建简历
	CreateResume(ctx context.Context, in *CreateResumeRequest, opts ...grpc.CallOption) (*ResumeResponse, error)
	// 更新简历
	UpdateResume(ctx context.Context, in *UpdateResumeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 删除简历
	DeleteResume(ctx context.Context, in *DeleteResumeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	//详情
	GetResumeDetail(ctx context.Context, in *GetResumeDetailRequest, opts ...grpc.CallOption) (*ResumeResponse, error)
}

type resumeClient struct {
	cc grpc.ClientConnInterface
}

func NewResumeClient(cc grpc.ClientConnInterface) ResumeClient {
	return &resumeClient{cc}
}

func (c *resumeClient) GetResumeList(ctx context.Context, in *GetResumeListRequest, opts ...grpc.CallOption) (*ResumeListResponse, error) {
	out := new(ResumeListResponse)
	err := c.cc.Invoke(ctx, "/Resume/GetResumeList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resumeClient) CreateResume(ctx context.Context, in *CreateResumeRequest, opts ...grpc.CallOption) (*ResumeResponse, error) {
	out := new(ResumeResponse)
	err := c.cc.Invoke(ctx, "/Resume/CreateResume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resumeClient) UpdateResume(ctx context.Context, in *UpdateResumeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Resume/UpdateResume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resumeClient) DeleteResume(ctx context.Context, in *DeleteResumeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Resume/DeleteResume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resumeClient) GetResumeDetail(ctx context.Context, in *GetResumeDetailRequest, opts ...grpc.CallOption) (*ResumeResponse, error) {
	out := new(ResumeResponse)
	err := c.cc.Invoke(ctx, "/Resume/GetResumeDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResumeServer is the server API for Resume service.
// All implementations must embed UnimplementedResumeServer
// for forward compatibility
type ResumeServer interface {
	// 获取简历列表
	GetResumeList(context.Context, *GetResumeListRequest) (*ResumeListResponse, error)
	// 创建简历
	CreateResume(context.Context, *CreateResumeRequest) (*ResumeResponse, error)
	// 更新简历
	UpdateResume(context.Context, *UpdateResumeRequest) (*emptypb.Empty, error)
	// 删除简历
	DeleteResume(context.Context, *DeleteResumeRequest) (*emptypb.Empty, error)
	//详情
	GetResumeDetail(context.Context, *GetResumeDetailRequest) (*ResumeResponse, error)
	mustEmbedUnimplementedResumeServer()
}

// UnimplementedResumeServer must be embedded to have forward compatible implementations.
type UnimplementedResumeServer struct {
}

func (UnimplementedResumeServer) GetResumeList(context.Context, *GetResumeListRequest) (*ResumeListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResumeList not implemented")
}
func (UnimplementedResumeServer) CreateResume(context.Context, *CreateResumeRequest) (*ResumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResume not implemented")
}
func (UnimplementedResumeServer) UpdateResume(context.Context, *UpdateResumeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateResume not implemented")
}
func (UnimplementedResumeServer) DeleteResume(context.Context, *DeleteResumeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteResume not implemented")
}
func (UnimplementedResumeServer) GetResumeDetail(context.Context, *GetResumeDetailRequest) (*ResumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResumeDetail not implemented")
}
func (UnimplementedResumeServer) mustEmbedUnimplementedResumeServer() {}

// UnsafeResumeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResumeServer will
// result in compilation errors.
type UnsafeResumeServer interface {
	mustEmbedUnimplementedResumeServer()
}

func RegisterResumeServer(s grpc.ServiceRegistrar, srv ResumeServer) {
	s.RegisterService(&Resume_ServiceDesc, srv)
}

func _Resume_GetResumeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResumeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeServer).GetResumeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Resume/GetResumeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeServer).GetResumeList(ctx, req.(*GetResumeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resume_CreateResume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateResumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeServer).CreateResume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Resume/CreateResume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeServer).CreateResume(ctx, req.(*CreateResumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resume_UpdateResume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateResumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeServer).UpdateResume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Resume/UpdateResume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeServer).UpdateResume(ctx, req.(*UpdateResumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resume_DeleteResume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteResumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeServer).DeleteResume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Resume/DeleteResume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeServer).DeleteResume(ctx, req.(*DeleteResumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resume_GetResumeDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResumeDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeServer).GetResumeDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Resume/GetResumeDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeServer).GetResumeDetail(ctx, req.(*GetResumeDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Resume_ServiceDesc is the grpc.ServiceDesc for Resume service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Resume_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Resume",
	HandlerType: (*ResumeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetResumeList",
			Handler:    _Resume_GetResumeList_Handler,
		},
		{
			MethodName: "CreateResume",
			Handler:    _Resume_CreateResume_Handler,
		},
		{
			MethodName: "UpdateResume",
			Handler:    _Resume_UpdateResume_Handler,
		},
		{
			MethodName: "DeleteResume",
			Handler:    _Resume_DeleteResume_Handler,
		},
		{
			MethodName: "GetResumeDetail",
			Handler:    _Resume_GetResumeDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resume.proto",
}
