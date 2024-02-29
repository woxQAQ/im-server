package svc

import (
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_user/internal/config"
	model "github.com/woxQAQ/im-service/pkg/common/sql/user"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	UserModel model.UserbasicModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserbasicModel(conn, c.CacheRedis),
	}
}
