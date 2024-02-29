package logic

import (
	"context"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_user/user"
	"github.com/woxQAQ/im-service/pkg/common/jwt"
	"time"

	"github.com/woxQAQ/im-service/internal/api/user/internal/svc"
	"github.com/woxQAQ/im-service/internal/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResp, err error) {
	login, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginRequest{
		Mobile:   req.Mobile,
		Email:    req.Email,
		Password: req.Password,
		Validate: req.Validate,
	})
	if err != nil {
		return nil, err
	}

	token, err := jwt.GetToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(), l.svcCtx.Config.Auth.AccessExpire, login.Id)
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		Token:    token,
		ExpireAt: time.Now().Unix() + l.svcCtx.Config.Auth.AccessExpire,
	}, nil
}
