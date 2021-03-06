// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: company.proto

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

// CompanyClient is the client API for Company service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompanyClient interface {
	// 获取公司列表
	GetCompanyList(ctx context.Context, in *GetCompanyListRequest, opts ...grpc.CallOption) (*CompanyListResponse, error)
	// 获取公司详情
	GetCompanyDetail(ctx context.Context, in *GetCompanyDetailRequest, opts ...grpc.CallOption) (*CompanyResponse, error)
	// 创建公司
	CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*CompanyResponse, error)
	// 更新公司信息
	UpdateCompany(ctx context.Context, in *UpdateCompanyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 删除公司
	DeleteCompany(ctx context.Context, in *DeleteCompanyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 获取我都公司列表
	GetMyCompanyList(ctx context.Context, in *GetMyCompanyListRequest, opts ...grpc.CallOption) (*CompanyListResponse, error)
	// 公司下所有用户,分页
	GetCompanyUserIdList(ctx context.Context, in *GetCompanyUserListRequest, opts ...grpc.CallOption) (*GetCompanyUserIdListResponse, error)
	// 加入公司
	CreateUserCompany(ctx context.Context, in *SaveUserCompanyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 关系更新
	UpdateUserCompany(ctx context.Context, in *SaveUserCompanyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 删除用户公司
	DeleteUserCompany(ctx context.Context, in *DeleteUserCompanyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type companyClient struct {
	cc grpc.ClientConnInterface
}

func NewCompanyClient(cc grpc.ClientConnInterface) CompanyClient {
	return &companyClient{cc}
}

func (c *companyClient) GetCompanyList(ctx context.Context, in *GetCompanyListRequest, opts ...grpc.CallOption) (*CompanyListResponse, error) {
	out := new(CompanyListResponse)
	err := c.cc.Invoke(ctx, "/Company/GetCompanyList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyClient) GetCompanyDetail(ctx context.Context, in *GetCompanyDetailRequest, opts ...grpc.CallOption) (*CompanyResponse, error) {
	out := new(CompanyResponse)
	err := c.cc.Invoke(ctx, "/Company/GetCompanyDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyClient) CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*CompanyResponse, error) {
	out := new(CompanyResponse)
	err := c.cc.Invoke(ctx, "/Company/CreateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyClient) UpdateCompany(ctx context.Context, in *UpdateCompanyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Company/UpdateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyClient) DeleteCompany(ctx context.Context, in *DeleteCompanyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Company/DeleteCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyClient) GetMyCompanyList(ctx context.Context, in *GetMyCompanyListRequest, opts ...grpc.CallOption) (*CompanyListResponse, error) {
	out := new(CompanyListResponse)
	err := c.cc.Invoke(ctx, "/Company/GetMyCompanyList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyClient) GetCompanyUserIdList(ctx context.Context, in *GetCompanyUserListRequest, opts ...grpc.CallOption) (*GetCompanyUserIdListResponse, error) {
	out := new(GetCompanyUserIdListResponse)
	err := c.cc.Invoke(ctx, "/Company/GetCompanyUserIdList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyClient) CreateUserCompany(ctx context.Context, in *SaveUserCompanyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Company/CreateUserCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyClient) UpdateUserCompany(ctx context.Context, in *SaveUserCompanyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Company/UpdateUserCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyClient) DeleteUserCompany(ctx context.Context, in *DeleteUserCompanyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Company/DeleteUserCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyServer is the server API for Company service.
// All implementations must embed UnimplementedCompanyServer
// for forward compatibility
type CompanyServer interface {
	// 获取公司列表
	GetCompanyList(context.Context, *GetCompanyListRequest) (*CompanyListResponse, error)
	// 获取公司详情
	GetCompanyDetail(context.Context, *GetCompanyDetailRequest) (*CompanyResponse, error)
	// 创建公司
	CreateCompany(context.Context, *CreateCompanyRequest) (*CompanyResponse, error)
	// 更新公司信息
	UpdateCompany(context.Context, *UpdateCompanyRequest) (*emptypb.Empty, error)
	// 删除公司
	DeleteCompany(context.Context, *DeleteCompanyRequest) (*emptypb.Empty, error)
	// 获取我都公司列表
	GetMyCompanyList(context.Context, *GetMyCompanyListRequest) (*CompanyListResponse, error)
	// 公司下所有用户,分页
	GetCompanyUserIdList(context.Context, *GetCompanyUserListRequest) (*GetCompanyUserIdListResponse, error)
	// 加入公司
	CreateUserCompany(context.Context, *SaveUserCompanyRequest) (*emptypb.Empty, error)
	// 关系更新
	UpdateUserCompany(context.Context, *SaveUserCompanyRequest) (*emptypb.Empty, error)
	// 删除用户公司
	DeleteUserCompany(context.Context, *DeleteUserCompanyRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedCompanyServer()
}

// UnimplementedCompanyServer must be embedded to have forward compatible implementations.
type UnimplementedCompanyServer struct {
}

func (UnimplementedCompanyServer) GetCompanyList(context.Context, *GetCompanyListRequest) (*CompanyListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyList not implemented")
}
func (UnimplementedCompanyServer) GetCompanyDetail(context.Context, *GetCompanyDetailRequest) (*CompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyDetail not implemented")
}
func (UnimplementedCompanyServer) CreateCompany(context.Context, *CreateCompanyRequest) (*CompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompany not implemented")
}
func (UnimplementedCompanyServer) UpdateCompany(context.Context, *UpdateCompanyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompany not implemented")
}
func (UnimplementedCompanyServer) DeleteCompany(context.Context, *DeleteCompanyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCompany not implemented")
}
func (UnimplementedCompanyServer) GetMyCompanyList(context.Context, *GetMyCompanyListRequest) (*CompanyListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyCompanyList not implemented")
}
func (UnimplementedCompanyServer) GetCompanyUserIdList(context.Context, *GetCompanyUserListRequest) (*GetCompanyUserIdListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyUserIdList not implemented")
}
func (UnimplementedCompanyServer) CreateUserCompany(context.Context, *SaveUserCompanyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserCompany not implemented")
}
func (UnimplementedCompanyServer) UpdateUserCompany(context.Context, *SaveUserCompanyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserCompany not implemented")
}
func (UnimplementedCompanyServer) DeleteUserCompany(context.Context, *DeleteUserCompanyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserCompany not implemented")
}
func (UnimplementedCompanyServer) mustEmbedUnimplementedCompanyServer() {}

// UnsafeCompanyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompanyServer will
// result in compilation errors.
type UnsafeCompanyServer interface {
	mustEmbedUnimplementedCompanyServer()
}

func RegisterCompanyServer(s grpc.ServiceRegistrar, srv CompanyServer) {
	s.RegisterService(&Company_ServiceDesc, srv)
}

func _Company_GetCompanyList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServer).GetCompanyList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Company/GetCompanyList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServer).GetCompanyList(ctx, req.(*GetCompanyListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Company_GetCompanyDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServer).GetCompanyDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Company/GetCompanyDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServer).GetCompanyDetail(ctx, req.(*GetCompanyDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Company_CreateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServer).CreateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Company/CreateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServer).CreateCompany(ctx, req.(*CreateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Company_UpdateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServer).UpdateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Company/UpdateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServer).UpdateCompany(ctx, req.(*UpdateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Company_DeleteCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServer).DeleteCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Company/DeleteCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServer).DeleteCompany(ctx, req.(*DeleteCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Company_GetMyCompanyList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyCompanyListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServer).GetMyCompanyList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Company/GetMyCompanyList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServer).GetMyCompanyList(ctx, req.(*GetMyCompanyListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Company_GetCompanyUserIdList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyUserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServer).GetCompanyUserIdList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Company/GetCompanyUserIdList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServer).GetCompanyUserIdList(ctx, req.(*GetCompanyUserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Company_CreateUserCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveUserCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServer).CreateUserCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Company/CreateUserCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServer).CreateUserCompany(ctx, req.(*SaveUserCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Company_UpdateUserCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveUserCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServer).UpdateUserCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Company/UpdateUserCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServer).UpdateUserCompany(ctx, req.(*SaveUserCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Company_DeleteUserCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServer).DeleteUserCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Company/DeleteUserCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServer).DeleteUserCompany(ctx, req.(*DeleteUserCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Company_ServiceDesc is the grpc.ServiceDesc for Company service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Company_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Company",
	HandlerType: (*CompanyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCompanyList",
			Handler:    _Company_GetCompanyList_Handler,
		},
		{
			MethodName: "GetCompanyDetail",
			Handler:    _Company_GetCompanyDetail_Handler,
		},
		{
			MethodName: "CreateCompany",
			Handler:    _Company_CreateCompany_Handler,
		},
		{
			MethodName: "UpdateCompany",
			Handler:    _Company_UpdateCompany_Handler,
		},
		{
			MethodName: "DeleteCompany",
			Handler:    _Company_DeleteCompany_Handler,
		},
		{
			MethodName: "GetMyCompanyList",
			Handler:    _Company_GetMyCompanyList_Handler,
		},
		{
			MethodName: "GetCompanyUserIdList",
			Handler:    _Company_GetCompanyUserIdList_Handler,
		},
		{
			MethodName: "CreateUserCompany",
			Handler:    _Company_CreateUserCompany_Handler,
		},
		{
			MethodName: "UpdateUserCompany",
			Handler:    _Company_UpdateUserCompany_Handler,
		},
		{
			MethodName: "DeleteUserCompany",
			Handler:    _Company_DeleteUserCompany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "company.proto",
}
