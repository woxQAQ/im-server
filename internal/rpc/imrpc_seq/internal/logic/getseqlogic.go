package logic

import (
	"context"

	"github.com/woxQAQ/im-service/internal/rpc/imrpc_seq/internal/svc"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_seq/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSeqLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSeqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSeqLogic {
	return &GetSeqLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSeqLogic) GetSeq(in *pb.GetSeqRequest) (*pb.GetSeqResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetSeqResponse{}, nil
}
