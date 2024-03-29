package logic

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_user/user"

	"github.com/woxQAQ/im-service/internal/api/api/internal/svc"
	"github.com/woxQAQ/im-service/internal/api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{Id: uid})
	if err != nil {
		return nil, err
	}
	return &types.UserInfoResp{
		Id:     res.Id,
		Name:   res.Name,
		Gender: res.Gender.String(),
		Email:  res.Email,
	}, nil
}
