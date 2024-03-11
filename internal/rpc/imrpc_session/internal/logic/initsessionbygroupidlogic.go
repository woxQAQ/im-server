package logic

import (
	"context"

	"github.com/woxQAQ/im-service/internal/rpc/imrpc_session/internal/svc"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_session/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitSessionByGroupIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitSessionByGroupIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitSessionByGroupIdLogic {
	return &InitSessionByGroupIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitSessionByGroupIdLogic) InitSessionByGroupId(in *pb.InitSessionByGroupIdRequest) (*pb.InitSessionResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.InitSessionResponse{}, nil
}
