// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: pkg/protobuf/doge/v1/doge_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	uuid "my-go/pkg/protobuf/uuid"
	doge "my-go/pkg/protobuf/doge"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DogeService_GetDoge_FullMethodName    = "/doge.v1.DogeService/GetDoge"
	DogeService_CreateDoge_FullMethodName = "/doge.v1.DogeService/CreateDoge"
)

// DogeServiceClient is the client API for DogeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DogeServiceClient interface {
	GetDoge(ctx context.Context, in *uuid.Uuid, opts ...grpc.CallOption) (*doge.Doge, error)
	CreateDoge(ctx context.Context, in *CreateDogeRequest, opts ...grpc.CallOption) (*doge.Doge, error)
}

type dogeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDogeServiceClient(cc grpc.ClientConnInterface) DogeServiceClient {
	return &dogeServiceClient{cc}
}

func (c *dogeServiceClient) GetDoge(ctx context.Context, in *uuid.Uuid, opts ...grpc.CallOption) (*doge.Doge, error) {
	out := new(doge.Doge)
	err := c.cc.Invoke(ctx, DogeService_GetDoge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dogeServiceClient) CreateDoge(ctx context.Context, in *CreateDogeRequest, opts ...grpc.CallOption) (*doge.Doge, error) {
	out := new(doge.Doge)
	err := c.cc.Invoke(ctx, DogeService_CreateDoge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DogeServiceServer is the server API for DogeService service.
// All implementations must embed UnimplementedDogeServiceServer
// for forward compatibility
type DogeServiceServer interface {
	GetDoge(context.Context, *uuid.Uuid) (*doge.Doge, error)
	CreateDoge(context.Context, *CreateDogeRequest) (*doge.Doge, error)
	mustEmbedUnimplementedDogeServiceServer()
}

// UnimplementedDogeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDogeServiceServer struct {
}

func (UnimplementedDogeServiceServer) GetDoge(context.Context, *uuid.Uuid) (*doge.Doge, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDoge not implemented")
}
func (UnimplementedDogeServiceServer) CreateDoge(context.Context, *CreateDogeRequest) (*doge.Doge, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDoge not implemented")
}
func (UnimplementedDogeServiceServer) mustEmbedUnimplementedDogeServiceServer() {}

// UnsafeDogeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DogeServiceServer will
// result in compilation errors.
type UnsafeDogeServiceServer interface {
	mustEmbedUnimplementedDogeServiceServer()
}

func RegisterDogeServiceServer(s grpc.ServiceRegistrar, srv DogeServiceServer) {
	s.RegisterService(&DogeService_ServiceDesc, srv)
}

func _DogeService_GetDoge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(uuid.Uuid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DogeServiceServer).GetDoge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DogeService_GetDoge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DogeServiceServer).GetDoge(ctx, req.(*uuid.Uuid))
	}
	return interceptor(ctx, in, info, handler)
}

func _DogeService_CreateDoge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDogeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DogeServiceServer).CreateDoge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DogeService_CreateDoge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DogeServiceServer).CreateDoge(ctx, req.(*CreateDogeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DogeService_ServiceDesc is the grpc.ServiceDesc for DogeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DogeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "doge.v1.DogeService",
	HandlerType: (*DogeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDoge",
			Handler:    _DogeService_GetDoge_Handler,
		},
		{
			MethodName: "CreateDoge",
			Handler:    _DogeService_CreateDoge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/protobuf/doge/v1/doge_service.proto",
}
