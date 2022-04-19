// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: counter.proto

package proto

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ResumeCounterServiceClient is the client API for ResumeCounterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResumeCounterServiceClient interface {
}

type resumeCounterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResumeCounterServiceClient(cc grpc.ClientConnInterface) ResumeCounterServiceClient {
	return &resumeCounterServiceClient{cc}
}

// ResumeCounterServiceServer is the server API for ResumeCounterService service.
// All implementations must embed UnimplementedResumeCounterServiceServer
// for forward compatibility
type ResumeCounterServiceServer interface {
	mustEmbedUnimplementedResumeCounterServiceServer()
}

// UnimplementedResumeCounterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedResumeCounterServiceServer struct {
}

func (UnimplementedResumeCounterServiceServer) mustEmbedUnimplementedResumeCounterServiceServer() {}

// UnsafeResumeCounterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResumeCounterServiceServer will
// result in compilation errors.
type UnsafeResumeCounterServiceServer interface {
	mustEmbedUnimplementedResumeCounterServiceServer()
}

func RegisterResumeCounterServiceServer(s grpc.ServiceRegistrar, srv ResumeCounterServiceServer) {
	s.RegisterService(&ResumeCounterService_ServiceDesc, srv)
}

// ResumeCounterService_ServiceDesc is the grpc.ServiceDesc for ResumeCounterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResumeCounterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ResumeCounterService",
	HandlerType: (*ResumeCounterServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "counter.proto",
}