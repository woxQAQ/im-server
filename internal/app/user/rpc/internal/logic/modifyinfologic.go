package logic

import (
	"context"

	"github.com/woxQAQ/im-service/internal/rpc/imrpc_user/internal/svc"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_user/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyInfoLogic {
	return &ModifyInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifyInfoLogic) ModifyInfo(in *pb.ModifyInfoRequest) (*pb.ModifyInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ModifyInfoResp{}, nil
}
