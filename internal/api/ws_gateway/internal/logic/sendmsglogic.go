package logic

import (
	"context"

	"github.com/woxQAQ/im-service/internal/api/ws_gateway/internal/svc"
	"github.com/woxQAQ/im-service/internal/api/ws_gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMsgLogic {
	return &SendMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendMsgLogic) SendMsg(req *types.SendRequest) (resp *types.SendResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
