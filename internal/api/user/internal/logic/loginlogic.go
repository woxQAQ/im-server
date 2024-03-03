package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_user/user"
	"github.com/woxQAQ/im-service/pkg/common/jwt"
	"time"

	"github.com/woxQAQ/im-service/internal/api/user/internal/svc"
	"github.com/woxQAQ/im-service/internal/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUnknownLoginMethod = errors.New("Login Method Unavailable")

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

func fetchTokenFromServer(secretkey string, req *types.LoginRequest) (string, error) {
	var (
		token string
		err   error
	)
	switch req.Method {
	case 0:
		// email login
		token, err = jwt.GetTokenWithEmail(secretkey, req.Email)
	case 1:
		// uid login
		token, err = jwt.GetTokenWithUid(secretkey, req.UserId)
	case 2:
		// phone login
		token, err = jwt.GetTokenWithPhone(secretkey, req.Mobile, 0)
	default:
		return "", ErrUnknownLoginMethod
	}
	return token, err
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

	token, err := fetchTokenFromServer(l.svcCtx.Config.Auth.AccessSecret, req)
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		Token:    token,
		ExpireAt: time.Now().Unix() + l.svcCtx.Config.Auth.AccessExpire,
	}, nil
}
