package logic

import (
	"context"

	"github.com/woxQAQ/im-service/internal/api/user/internal/svc"
	"github.com/woxQAQ/im-service/internal/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyInfoLogic {
	return &ModifyInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyInfoLogic) ModifyInfo(req *types.ModifyRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
