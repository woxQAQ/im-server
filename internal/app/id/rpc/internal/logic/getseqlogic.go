package logic

import (
	"context"
	"github.com/woxQAQ/im-service/internal/app/id/rpc/internal/svc"
	"github.com/woxQAQ/im-service/internal/app/id/rpc/pb"

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

func (l *GetSeqLogic) GetSeq(in *pb.GetSeqReq) (*pb.GetSeqResp, error) {
	// todo: add your logic here and delete this line
	switch in.CallerType {
	case pb.RequestCaller_REQUEST_CALLER_USER_CREATE:
	}
	return &pb.GetSeqResp{}, nil
}
