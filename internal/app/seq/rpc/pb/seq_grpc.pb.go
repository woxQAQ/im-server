// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: seq.proto

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
	Seq_GetSessionId_FullMethodName     = "/seq.Seq/GetSessionId"
	Seq_GetMessageIds_FullMethodName    = "/seq.Seq/GetMessageIds"
	Seq_GetSessionSeq_FullMethodName    = "/seq.Seq/GetSessionSeq"
	Seq_UpdateSessionSeq_FullMethodName = "/seq.Seq/UpdateSessionSeq"
)

// SeqClient is the client API for Seq service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SeqClient interface {
	GetSessionId(ctx context.Context, in *GetSessionIdRequest, opts ...grpc.CallOption) (*GetSessionIdResponse, error)
	GetMessageIds(ctx context.Context, in *GetMessageIdsRequest, opts ...grpc.CallOption) (*GetMessageIdsResponse, error)
	GetSessionSeq(ctx context.Context, in *GetSeqRequest, opts ...grpc.CallOption) (*GetSeqResponse, error)
	UpdateSessionSeq(ctx context.Context, in *UpdateSessionSeqRequest, opts ...grpc.CallOption) (*UpdateSessionSeqResponse, error)
}

type seqClient struct {
	cc grpc.ClientConnInterface
}

func NewSeqClient(cc grpc.ClientConnInterface) SeqClient {
	return &seqClient{cc}
}

func (c *seqClient) GetSessionId(ctx context.Context, in *GetSessionIdRequest, opts ...grpc.CallOption) (*GetSessionIdResponse, error) {
	out := new(GetSessionIdResponse)
	err := c.cc.Invoke(ctx, Seq_GetSessionId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seqClient) GetMessageIds(ctx context.Context, in *GetMessageIdsRequest, opts ...grpc.CallOption) (*GetMessageIdsResponse, error) {
	out := new(GetMessageIdsResponse)
	err := c.cc.Invoke(ctx, Seq_GetMessageIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seqClient) GetSessionSeq(ctx context.Context, in *GetSeqRequest, opts ...grpc.CallOption) (*GetSeqResponse, error) {
	out := new(GetSeqResponse)
	err := c.cc.Invoke(ctx, Seq_GetSessionSeq_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seqClient) UpdateSessionSeq(ctx context.Context, in *UpdateSessionSeqRequest, opts ...grpc.CallOption) (*UpdateSessionSeqResponse, error) {
	out := new(UpdateSessionSeqResponse)
	err := c.cc.Invoke(ctx, Seq_UpdateSessionSeq_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SeqServer is the server API for Seq service.
// All implementations must embed UnimplementedSeqServer
// for forward compatibility
type SeqServer interface {
	GetSessionId(context.Context, *GetSessionIdRequest) (*GetSessionIdResponse, error)
	GetMessageIds(context.Context, *GetMessageIdsRequest) (*GetMessageIdsResponse, error)
	GetSessionSeq(context.Context, *GetSeqRequest) (*GetSeqResponse, error)
	UpdateSessionSeq(context.Context, *UpdateSessionSeqRequest) (*UpdateSessionSeqResponse, error)
	mustEmbedUnimplementedSeqServer()
}

// UnimplementedSeqServer must be embedded to have forward compatible implementations.
type UnimplementedSeqServer struct {
}

func (UnimplementedSeqServer) GetSessionId(context.Context, *GetSessionIdRequest) (*GetSessionIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSessionId not implemented")
}
func (UnimplementedSeqServer) GetMessageIds(context.Context, *GetMessageIdsRequest) (*GetMessageIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageIds not implemented")
}
func (UnimplementedSeqServer) GetSessionSeq(context.Context, *GetSeqRequest) (*GetSeqResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSessionSeq not implemented")
}
func (UnimplementedSeqServer) UpdateSessionSeq(context.Context, *UpdateSessionSeqRequest) (*UpdateSessionSeqResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSessionSeq not implemented")
}
func (UnimplementedSeqServer) mustEmbedUnimplementedSeqServer() {}

// UnsafeSeqServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SeqServer will
// result in compilation errors.
type UnsafeSeqServer interface {
	mustEmbedUnimplementedSeqServer()
}

func RegisterSeqServer(s grpc.ServiceRegistrar, srv SeqServer) {
	s.RegisterService(&Seq_ServiceDesc, srv)
}

func _Seq_GetSessionId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeqServer).GetSessionId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Seq_GetSessionId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeqServer).GetSessionId(ctx, req.(*GetSessionIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seq_GetMessageIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeqServer).GetMessageIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Seq_GetMessageIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeqServer).GetMessageIds(ctx, req.(*GetMessageIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seq_GetSessionSeq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSeqRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeqServer).GetSessionSeq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Seq_GetSessionSeq_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeqServer).GetSessionSeq(ctx, req.(*GetSeqRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seq_UpdateSessionSeq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSessionSeqRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeqServer).UpdateSessionSeq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Seq_UpdateSessionSeq_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeqServer).UpdateSessionSeq(ctx, req.(*UpdateSessionSeqRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Seq_ServiceDesc is the grpc.ServiceDesc for Seq service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Seq_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "seq.Seq",
	HandlerType: (*SeqServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSessionId",
			Handler:    _Seq_GetSessionId_Handler,
		},
		{
			MethodName: "GetMessageIds",
			Handler:    _Seq_GetMessageIds_Handler,
		},
		{
			MethodName: "GetSessionSeq",
			Handler:    _Seq_GetSessionSeq_Handler,
		},
		{
			MethodName: "UpdateSessionSeq",
			Handler:    _Seq_UpdateSessionSeq_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "seq.proto",
}