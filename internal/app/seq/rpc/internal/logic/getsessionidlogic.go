package logic

import (
	"context"
	"github.com/woxQAQ/im-service/internal/app/seq/model"
	"github.com/woxQAQ/im-service/internal/app/seq/rpc/internal/svc"
	"github.com/woxQAQ/im-service/internal/app/seq/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSessionIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSessionIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSessionIdLogic {
	return &GetSessionIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSessionIdLogic) GetSessionId(in *pb.GetSessionIdRequest) (*pb.GetSessionIdResponse, error) {
	// todo: add your logic here and delete this line

	isExist, _ := l.svcCtx.SessionIdModel.ExistSessions(context.Background(), in.UserId_1, in.UserId_2)
	if isExist {
		return &pb.GetSessionIdResponse{
			Base: &pb.RespBase{
				ErrCode: 13,
				ErrMsg:  "Session has been established",
			},
		}, nil
	}

	id := l.svcCtx.SnowFlake.Generate()
	insert, err := l.svcCtx.SessionIdModel.Insert(context.Background(), &model.SessionId{
		Session: id.Int64(),
		UserId1: in.UserId_1,
		UserId2: in.UserId_2,
	})
	if err != nil {
		return nil, err
	}
	return &pb.GetSessionIdResponse{}, nil
}
