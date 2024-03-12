// Code generated by goctl. DO NOT EDIT.
// Source: seq.proto

package server

import (
	"context"

	"github.com/woxQAQ/im-service/internal/rpc/imrpc_seq/internal/logic"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_seq/internal/svc"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_seq/pb"
)

type SeqServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedSeqServer
}

func NewSeqServer(svcCtx *svc.ServiceContext) *SeqServer {
	return &SeqServer{
		svcCtx: svcCtx,
	}
}

func (s *SeqServer) GetSessionId(ctx context.Context, in *pb.GetSessionIdRequest) (*pb.GetSessionIdResponse, error) {
	l := logic.NewGetSessionIdLogic(ctx, s.svcCtx)
	return l.GetSessionId(in)
}

func (s *SeqServer) GetMessageIds(ctx context.Context, in *pb.GetMessageIdsRequest) (*pb.GetMessageIdsResponse, error) {
	l := logic.NewGetMessageIdsLogic(ctx, s.svcCtx)
	return l.GetMessageIds(in)
}

func (s *SeqServer) GetSessionSeq(ctx context.Context, in *pb.GetSeqRequest) (*pb.GetSeqResponse, error) {
	l := logic.NewGetSessionSeqLogic(ctx, s.svcCtx)
	return l.GetSessionSeq(in)
}

func (s *SeqServer) UpdateSessionSeq(ctx context.Context, in *pb.UpdateSessionSeqRequest) (*pb.UpdateSessionSeqResponse, error) {
	l := logic.NewUpdateSessionSeqLogic(ctx, s.svcCtx)
	return l.UpdateSessionSeq(in)
}
