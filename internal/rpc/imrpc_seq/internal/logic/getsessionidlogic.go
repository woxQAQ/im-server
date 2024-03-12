package logic

import (
	"context"

	"github.com/woxQAQ/im-service/internal/rpc/imrpc_seq/internal/svc"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_seq/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSessionIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSessionIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSessionIdLogic {
	return &GetSessionIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSessionIdLogic) GetSessionId(in *pb.GetSessionIdRequest) (*pb.GetSessionIdResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetSessionIdResponse{}, nil
}
