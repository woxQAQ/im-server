package logic

import (
	"context"
	"errors"
	model2 "github.com/woxQAQ/im-service/internal/app/user/model"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_user/pb"
	"github.com/woxQAQ/im-service/pkg/common/crypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/woxQAQ/im-service/internal/rpc/imrpc_user/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var (
	methodPhone = "phone"
	methodEmail = "email"
)

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func loginRequestIsValid(in *pb.LoginRequest) (bool, string) {
	// todo: add your logic here and delete this line
	if in.Email == "" && in.Mobile == "" {
		// email and mobile phone are empty
		return false, ""
	} else if in.Email == "" && in.Mobile != "" {
		if in.Validate == "" {
			return false, ""
		}
		return true, methodPhone
	} else if in.Email != "" && in.Mobile != "" {
		return false, ""
	} else {
		if in.Password == "" {
			return false, ""
		}
		return true, methodEmail
	}
}

func (l *LoginLogic) Login(in *pb.LoginRequest) (*pb.LoginResp, error) {
	// we will use phone OR email to login, so the LoginRequest will just flitter phone OR email
	// a request with phone and email will be rejected
	ok, method := loginRequestIsValid(in)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "Login request is invalid")
	}

	var (
		user *model2.Userbasic
		err  error
	)
	switch method {
	case methodEmail:
		user, err = l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	case methodPhone:
		user, err = l.svcCtx.UserModel.FindOneByMobilePhone(l.ctx, in.Mobile)
	}

	if err != nil {
		if errors.Is(err, model2.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "User not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	password := crypt.PasswordEncrypt(in.Password, l.svcCtx.Config.Salt)
	if password != user.Password {
		return nil, status.Error(codes.Unauthenticated, "Password is wrong")
	}

	// todo: get token
	return &pb.LoginResp{}, nil
}
