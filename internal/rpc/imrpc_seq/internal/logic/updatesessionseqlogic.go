package logic

import (
	"context"

	"github.com/woxQAQ/im-service/internal/rpc/imrpc_seq/internal/svc"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_seq/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSessionSeqLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSessionSeqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSessionSeqLogic {
	return &UpdateSessionSeqLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSessionSeqLogic) UpdateSessionSeq(in *pb.UpdateSessionSeqRequest) (*pb.UpdateSessionSeqResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateSessionSeqResponse{}, nil
}
