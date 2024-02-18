package logic

import (
	"context"

	"github.com/woxQAQ/im-service/internal/rpc/user/internal/svc"
	"github.com/woxQAQ/im-service/internal/rpc/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *pb.UserInfoRequest) (*pb.UserInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UserInfoResp{}, nil
}
