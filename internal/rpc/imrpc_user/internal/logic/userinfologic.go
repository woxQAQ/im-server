package logic

import (
	"context"
	"github.com/pkg/errors"
	model "github.com/woxQAQ/im-service/pkg/common/model/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/woxQAQ/im-service/internal/rpc/imrpc_user/internal/svc"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_user/pb"

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
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "User not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.UserInfoResp{
		Id:     user.Id,
		Name:   user.Name,
		Gender: pb.Gender(pb.Gender_value[user.Gender]),
		Email:  user.Email,
		Mobile: user.MobilePhone,
	}, nil
}
