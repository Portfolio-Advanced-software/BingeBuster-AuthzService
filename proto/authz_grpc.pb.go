// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: proto/authz.proto

package authzpb

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
	AuthzService_VerifyRole_FullMethodName = "/authz.AuthzService/VerifyRole"
)

// AuthzServiceClient is the client API for AuthzService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthzServiceClient interface {
	VerifyRole(ctx context.Context, in *VerifyRoleRequest, opts ...grpc.CallOption) (*VerifyRoleResponse, error)
}

type authzServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthzServiceClient(cc grpc.ClientConnInterface) AuthzServiceClient {
	return &authzServiceClient{cc}
}

func (c *authzServiceClient) VerifyRole(ctx context.Context, in *VerifyRoleRequest, opts ...grpc.CallOption) (*VerifyRoleResponse, error) {
	out := new(VerifyRoleResponse)
	err := c.cc.Invoke(ctx, AuthzService_VerifyRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthzServiceServer is the server API for AuthzService service.
// All implementations must embed UnimplementedAuthzServiceServer
// for forward compatibility
type AuthzServiceServer interface {
	VerifyRole(context.Context, *VerifyRoleRequest) (*VerifyRoleResponse, error)
	mustEmbedUnimplementedAuthzServiceServer()
}

// UnimplementedAuthzServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthzServiceServer struct {
}

func (UnimplementedAuthzServiceServer) VerifyRole(context.Context, *VerifyRoleRequest) (*VerifyRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyRole not implemented")
}
func (UnimplementedAuthzServiceServer) mustEmbedUnimplementedAuthzServiceServer() {}

// UnsafeAuthzServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthzServiceServer will
// result in compilation errors.
type UnsafeAuthzServiceServer interface {
	mustEmbedUnimplementedAuthzServiceServer()
}

func RegisterAuthzServiceServer(s grpc.ServiceRegistrar, srv AuthzServiceServer) {
	s.RegisterService(&AuthzService_ServiceDesc, srv)
}

func _AuthzService_VerifyRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthzServiceServer).VerifyRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthzService_VerifyRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthzServiceServer).VerifyRole(ctx, req.(*VerifyRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthzService_ServiceDesc is the grpc.ServiceDesc for AuthzService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthzService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authz.AuthzService",
	HandlerType: (*AuthzServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyRole",
			Handler:    _AuthzService_VerifyRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/authz.proto",
}
