package logic

import (
	"context"
	"github.com/woxQAQ/im-service/internal/rpc/user/user"

	"github.com/woxQAQ/im-service/internal/api/user/internal/svc"
	"github.com/woxQAQ/im-service/internal/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line
	register, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterRequest{
		Name:     req.Name,
		Email:    req.Email,
		Gender:   req.Gender,
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.RegisterResp{
		Id:     register.Id,
		Name:   register.Name,
		Gender: register.Gender,
		Email:  register.Email,
	}, nil
}
