// Code generated by goctl. DO NOT EDIT.
// Source: seq.proto

package seq

import (
	"context"

	"github.com/woxQAQ/im-service/internal/rpc/imrpc_seq/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetMessageIdsRequest     = pb.GetMessageIdsRequest
	GetMessageIdsResponse    = pb.GetMessageIdsResponse
	GetSeqRequest            = pb.GetSeqRequest
	GetSeqResponse           = pb.GetSeqResponse
	GetSessionIdRequest      = pb.GetSessionIdRequest
	GetSessionIdResponse     = pb.GetSessionIdResponse
	RespBase                 = pb.RespBase
	UpdateSessionSeqRequest  = pb.UpdateSessionSeqRequest
	UpdateSessionSeqResponse = pb.UpdateSessionSeqResponse

	Seq interface {
		GetSessionId(ctx context.Context, in *GetSessionIdRequest, opts ...grpc.CallOption) (*GetSessionIdResponse, error)
		GetMessageIds(ctx context.Context, in *GetMessageIdsRequest, opts ...grpc.CallOption) (*GetMessageIdsResponse, error)
		GetSessionSeq(ctx context.Context, in *GetSeqRequest, opts ...grpc.CallOption) (*GetSeqResponse, error)
		UpdateSessionSeq(ctx context.Context, in *UpdateSessionSeqRequest, opts ...grpc.CallOption) (*UpdateSessionSeqResponse, error)
	}

	defaultSeq struct {
		cli zrpc.Client
	}
)

func NewSeq(cli zrpc.Client) Seq {
	return &defaultSeq{
		cli: cli,
	}
}

func (m *defaultSeq) GetSessionId(ctx context.Context, in *GetSessionIdRequest, opts ...grpc.CallOption) (*GetSessionIdResponse, error) {
	client := pb.NewSeqClient(m.cli.Conn())
	return client.GetSessionId(ctx, in, opts...)
}

func (m *defaultSeq) GetMessageIds(ctx context.Context, in *GetMessageIdsRequest, opts ...grpc.CallOption) (*GetMessageIdsResponse, error) {
	client := pb.NewSeqClient(m.cli.Conn())
	return client.GetMessageIds(ctx, in, opts...)
}

func (m *defaultSeq) GetSessionSeq(ctx context.Context, in *GetSeqRequest, opts ...grpc.CallOption) (*GetSeqResponse, error) {
	client := pb.NewSeqClient(m.cli.Conn())
	return client.GetSessionSeq(ctx, in, opts...)
}

func (m *defaultSeq) UpdateSessionSeq(ctx context.Context, in *UpdateSessionSeqRequest, opts ...grpc.CallOption) (*UpdateSessionSeqResponse, error) {
	client := pb.NewSeqClient(m.cli.Conn())
	return client.UpdateSessionSeq(ctx, in, opts...)
}
