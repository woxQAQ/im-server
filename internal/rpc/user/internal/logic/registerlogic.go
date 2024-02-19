package logic

import (
	"context"
	"errors"
	"github.com/woxQAQ/im-service/internal/rpc/user/user"
	"github.com/woxQAQ/im-service/pkg/common/crypt"
	model "github.com/woxQAQ/im-service/pkg/common/sql/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/woxQAQ/im-service/internal/rpc/user/internal/svc"
	"github.com/woxQAQ/im-service/internal/rpc/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterRequest) (*pb.RegisterResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err == nil {
		return nil, status.Error(codes.AlreadyExists, "User exists")
	}

	if errors.Is(err, model.ErrNotFound) {
		newUser := model.Userbasic{
			Name:        in.Name,
			Gender:      in.Gender.String(),
			MobilePhone: in.Mobile,
			Email:       in.Email,
			Password:    crypt.PasswordEncrypt(in.Password, l.svcCtx.Config.Salt),
		}
		res, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		newUser.Id, err = res.LastInsertId()
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		return &user.RegisterResp{
			Id:     newUser.Id,
			Name:   newUser.Name,
			Gender: in.Gender,
			Email:  newUser.Email,
		}, nil
	}
	return nil, status.Error(500, err.Error())
}
