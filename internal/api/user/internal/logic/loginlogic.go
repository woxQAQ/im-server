package logic

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_user/user"
	"github.com/woxQAQ/im-service/pkg/common/jwt"

	"github.com/woxQAQ/im-service/internal/api/user/internal/svc"
	"github.com/woxQAQ/im-service/internal/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUnknownLoginMethod = errors.New("Login Method Unavailable")
var ErrRequestUnavailable = errors.New("Request Unavailable")

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

func fetchTokenFromServer(secretkey string, req *types.LoginRequest, accessExpired int64) (string, error) {
	var (
		token string
		err   error
	)
	switch req.Method {
	case 0:
		// email login
		token, err = jwt.GetToken(secretkey, req.Email, accessExpired)
	case 1:
		// uid login
		token, err = jwt.GetToken(secretkey, req.UserId, accessExpired)
	case 2:
		// phone login
		token, err = jwt.GetToken(secretkey, req.Mobile, accessExpired)
	default:
		return "", ErrUnknownLoginMethod
	}
	return token, err
}



func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResp, err error) {
	switch req.Method {
	case 0:
		if req.Email == "" {
			resp = nil
			err = ErrRequestUnavailable
			return
		}
	case 1:
		if req.UserId == "" {
			resp = nil
			err = ErrRequestUnavailable
			return
		}
	case 2:
		if req.Mobile == "" {
			resp = nil
			err = ErrRequestUnavailable
			return
		}
	}
	_, err = l.svcCtx.UserRpc.Login(l.ctx, &user.LoginRequest{
		Userid: req.UserId,
		Mobile:   req.Mobile,
		Email:    req.Email,
		Password: req.Password,
		Validate: req.Validate,
	})
	if err != nil {
		return nil, err
	}

	token, err := fetchTokenFromServer(l.svcCtx.Config.Auth.AccessSecret, req, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		Token:    token,
		ExpireAt: time.Now().Unix() + l.svcCtx.Config.Auth.AccessExpire,
	}, nil
}
