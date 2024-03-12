// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: session.proto

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

const (
	Session_InitSessionByUsersId_FullMethodName = "/session.Session/InitSessionByUsersId"
	Session_InitSessionByGroupId_FullMethodName = "/session.Session/InitSessionByGroupId"
)

// SessionClient is the client API for Session service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SessionClient interface {
	InitSessionByUsersId(ctx context.Context, in *InitSessionByUsersIdRequest, opts ...grpc.CallOption) (*InitSessionResponse, error)
	InitSessionByGroupId(ctx context.Context, in *InitSessionByGroupIdRequest, opts ...grpc.CallOption) (*InitSessionResponse, error)
}

type sessionClient struct {
	cc grpc.ClientConnInterface
}

func NewSessionClient(cc grpc.ClientConnInterface) SessionClient {
	return &sessionClient{cc}
}

func (c *sessionClient) InitSessionByUsersId(ctx context.Context, in *InitSessionByUsersIdRequest, opts ...grpc.CallOption) (*InitSessionResponse, error) {
	out := new(InitSessionResponse)
	err := c.cc.Invoke(ctx, Session_InitSessionByUsersId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionClient) InitSessionByGroupId(ctx context.Context, in *InitSessionByGroupIdRequest, opts ...grpc.CallOption) (*InitSessionResponse, error) {
	out := new(InitSessionResponse)
	err := c.cc.Invoke(ctx, Session_InitSessionByGroupId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionServer is the server API for Session service.
// All implementations must embed UnimplementedSessionServer
// for forward compatibility
type SessionServer interface {
	InitSessionByUsersId(context.Context, *InitSessionByUsersIdRequest) (*InitSessionResponse, error)
	InitSessionByGroupId(context.Context, *InitSessionByGroupIdRequest) (*InitSessionResponse, error)
	mustEmbedUnimplementedSessionServer()
}

// UnimplementedSessionServer must be embedded to have forward compatible implementations.
type UnimplementedSessionServer struct {
}

func (UnimplementedSessionServer) InitSessionByUsersId(context.Context, *InitSessionByUsersIdRequest) (*InitSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitSessionByUsersId not implemented")
}
func (UnimplementedSessionServer) InitSessionByGroupId(context.Context, *InitSessionByGroupIdRequest) (*InitSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitSessionByGroupId not implemented")
}
func (UnimplementedSessionServer) mustEmbedUnimplementedSessionServer() {}

// UnsafeSessionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SessionServer will
// result in compilation errors.
type UnsafeSessionServer interface {
	mustEmbedUnimplementedSessionServer()
}

func RegisterSessionServer(s grpc.ServiceRegistrar, srv SessionServer) {
	s.RegisterService(&Session_ServiceDesc, srv)
}

func _Session_InitSessionByUsersId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitSessionByUsersIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServer).InitSessionByUsersId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Session_InitSessionByUsersId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServer).InitSessionByUsersId(ctx, req.(*InitSessionByUsersIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Session_InitSessionByGroupId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitSessionByGroupIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServer).InitSessionByGroupId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Session_InitSessionByGroupId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServer).InitSessionByGroupId(ctx, req.(*InitSessionByGroupIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Session_ServiceDesc is the grpc.ServiceDesc for Session service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Session_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "session.Session",
	HandlerType: (*SessionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitSessionByUsersId",
			Handler:    _Session_InitSessionByUsersId_Handler,
		},
		{
			MethodName: "InitSessionByGroupId",
			Handler:    _Session_InitSessionByGroupId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "session.proto",
}
