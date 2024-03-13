package logic

import (
	"context"

	"github.com/woxQAQ/im-service/internal/rpc/imrpc_seq/internal/svc"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_seq/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMessageIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageIdsLogic {
	return &GetMessageIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMessageIdsLogic) GetMessageIds(in *pb.GetMessageIdsRequest) (*pb.GetMessageIdsResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetMessageIdsResponse{}, nil
}
