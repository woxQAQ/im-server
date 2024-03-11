package logic

import (
	"context"

	"github.com/woxQAQ/im-service/internal/rpc/imrpc_session/internal/svc"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_session/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitSessionByUsersIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitSessionByUsersIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitSessionByUsersIdLogic {
	return &InitSessionByUsersIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitSessionByUsersIdLogic) InitSessionByUsersId(in *pb.InitSessionByUsersIdRequest) (*pb.InitSessionResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.InitSessionResponse{}, nil
}
