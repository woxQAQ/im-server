package logic

import (
	"context"

	"github.com/pkg/errors"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_message/internal/svc"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_message/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var (
	ErrReqDataEmpty = errors.New("Data in the Request is empty")
)

func NewSendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMsgLogic {
	return &SendMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMsgLogic) SendMsg(in *pb.SendMessageReq) (*pb.SendMessageResp, error) {
	// todo: add your logic here and delete this line
	// if in.Data == nil {
	// 	return nil, ErrReqDataEmpty
	// }
	return &pb.SendMessageResp{}, nil
}
