package svc

import (
	"github.com/woxQAQ/im-service/internal/app/id/model"
	"github.com/woxQAQ/im-service/internal/app/id/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	IdGen model.IdGeneratorModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		IdGen:  model.NewIdGeneratorModel(conn, c.CacheRedis),
	}
}
