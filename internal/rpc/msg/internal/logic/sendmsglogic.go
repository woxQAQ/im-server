package logic

import (
	"context"
	"github.com/woxQAQ/im-service/internal/rpc/msg/internal/svc"
	"github.com/woxQAQ/im-service/internal/rpc/msg/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMsgLogic {
	return &SendMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMsgLogic) SendMsg(in *pb.SendMessageReq) (*pb.SendMessageResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SendMessageResp{}, nil
}
