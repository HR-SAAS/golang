// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: user_post.proto

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

// UserPostClient is the client API for UserPost service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserPostClient interface {
	// 获取岗位 投递列表
	GetUserPostList(ctx context.Context, in *GetUserPostListRequest, opts ...grpc.CallOption) (*UserPostListResponse, error)
	// 投递简历
	CreateUserPost(ctx context.Context, in *CreateUserPostRequest, opts ...grpc.CallOption) (*UserPostResponse, error)
	// 修改投递
	UpdateUserPost(ctx context.Context, in *UpdateUserPostRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 取消投递
	DeleteUserPost(ctx context.Context, in *DeleteUserPostRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 获取投递详情
	GetUserPostDetail(ctx context.Context, in *GetUserPostDetailRequest, opts ...grpc.CallOption) (*UserPostResponse, error)
	// 批量修改状态
	BatchUpdateUserPost(ctx context.Context, in *BatchUpdateUserPostRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type userPostClient struct {
	cc grpc.ClientConnInterface
}

func NewUserPostClient(cc grpc.ClientConnInterface) UserPostClient {
	return &userPostClient{cc}
}

func (c *userPostClient) GetUserPostList(ctx context.Context, in *GetUserPostListRequest, opts ...grpc.CallOption) (*UserPostListResponse, error) {
	out := new(UserPostListResponse)
	err := c.cc.Invoke(ctx, "/UserPost/GetUserPostList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPostClient) CreateUserPost(ctx context.Context, in *CreateUserPostRequest, opts ...grpc.CallOption) (*UserPostResponse, error) {
	out := new(UserPostResponse)
	err := c.cc.Invoke(ctx, "/UserPost/CreateUserPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPostClient) UpdateUserPost(ctx context.Context, in *UpdateUserPostRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/UserPost/UpdateUserPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPostClient) DeleteUserPost(ctx context.Context, in *DeleteUserPostRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/UserPost/DeleteUserPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPostClient) GetUserPostDetail(ctx context.Context, in *GetUserPostDetailRequest, opts ...grpc.CallOption) (*UserPostResponse, error) {
	out := new(UserPostResponse)
	err := c.cc.Invoke(ctx, "/UserPost/GetUserPostDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPostClient) BatchUpdateUserPost(ctx context.Context, in *BatchUpdateUserPostRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/UserPost/BatchUpdateUserPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserPostServer is the server API for UserPost service.
// All implementations must embed UnimplementedUserPostServer
// for forward compatibility
type UserPostServer interface {
	// 获取岗位 投递列表
	GetUserPostList(context.Context, *GetUserPostListRequest) (*UserPostListResponse, error)
	// 投递简历
	CreateUserPost(context.Context, *CreateUserPostRequest) (*UserPostResponse, error)
	// 修改投递
	UpdateUserPost(context.Context, *UpdateUserPostRequest) (*emptypb.Empty, error)
	// 取消投递
	DeleteUserPost(context.Context, *DeleteUserPostRequest) (*emptypb.Empty, error)
	// 获取投递详情
	GetUserPostDetail(context.Context, *GetUserPostDetailRequest) (*UserPostResponse, error)
	// 批量修改状态
	BatchUpdateUserPost(context.Context, *BatchUpdateUserPostRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedUserPostServer()
}

// UnimplementedUserPostServer must be embedded to have forward compatible implementations.
type UnimplementedUserPostServer struct {
}

func (UnimplementedUserPostServer) GetUserPostList(context.Context, *GetUserPostListRequest) (*UserPostListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPostList not implemented")
}
func (UnimplementedUserPostServer) CreateUserPost(context.Context, *CreateUserPostRequest) (*UserPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserPost not implemented")
}
func (UnimplementedUserPostServer) UpdateUserPost(context.Context, *UpdateUserPostRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserPost not implemented")
}
func (UnimplementedUserPostServer) DeleteUserPost(context.Context, *DeleteUserPostRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserPost not implemented")
}
func (UnimplementedUserPostServer) GetUserPostDetail(context.Context, *GetUserPostDetailRequest) (*UserPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPostDetail not implemented")
}
func (UnimplementedUserPostServer) BatchUpdateUserPost(context.Context, *BatchUpdateUserPostRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchUpdateUserPost not implemented")
}
func (UnimplementedUserPostServer) mustEmbedUnimplementedUserPostServer() {}

// UnsafeUserPostServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserPostServer will
// result in compilation errors.
type UnsafeUserPostServer interface {
	mustEmbedUnimplementedUserPostServer()
}

func RegisterUserPostServer(s grpc.ServiceRegistrar, srv UserPostServer) {
	s.RegisterService(&UserPost_ServiceDesc, srv)
}

func _UserPost_GetUserPostList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserPostListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPostServer).GetUserPostList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserPost/GetUserPostList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPostServer).GetUserPostList(ctx, req.(*GetUserPostListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPost_CreateUserPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPostServer).CreateUserPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserPost/CreateUserPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPostServer).CreateUserPost(ctx, req.(*CreateUserPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPost_UpdateUserPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPostServer).UpdateUserPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserPost/UpdateUserPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPostServer).UpdateUserPost(ctx, req.(*UpdateUserPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPost_DeleteUserPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPostServer).DeleteUserPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserPost/DeleteUserPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPostServer).DeleteUserPost(ctx, req.(*DeleteUserPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPost_GetUserPostDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserPostDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPostServer).GetUserPostDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserPost/GetUserPostDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPostServer).GetUserPostDetail(ctx, req.(*GetUserPostDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPost_BatchUpdateUserPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchUpdateUserPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPostServer).BatchUpdateUserPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserPost/BatchUpdateUserPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPostServer).BatchUpdateUserPost(ctx, req.(*BatchUpdateUserPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserPost_ServiceDesc is the grpc.ServiceDesc for UserPost service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserPost_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserPost",
	HandlerType: (*UserPostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserPostList",
			Handler:    _UserPost_GetUserPostList_Handler,
		},
		{
			MethodName: "CreateUserPost",
			Handler:    _UserPost_CreateUserPost_Handler,
		},
		{
			MethodName: "UpdateUserPost",
			Handler:    _UserPost_UpdateUserPost_Handler,
		},
		{
			MethodName: "DeleteUserPost",
			Handler:    _UserPost_DeleteUserPost_Handler,
		},
		{
			MethodName: "GetUserPostDetail",
			Handler:    _UserPost_GetUserPostDetail_Handler,
		},
		{
			MethodName: "BatchUpdateUserPost",
			Handler:    _UserPost_BatchUpdateUserPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_post.proto",
}
