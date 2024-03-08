package logic

import (
	"context"

	"github.com/woxQAQ/im-service/internal/rpc/imrpc_seq/internal/svc"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_seq/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetDefaultSeqLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetDefaultSeqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetDefaultSeqLogic {
	return &SetDefaultSeqLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetDefaultSeqLogic) SetDefaultSeq(in *pb.SetDefaultSeqRequest) (*pb.SetDefaultSeqResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.SetDefaultSeqResponse{}, nil
}
