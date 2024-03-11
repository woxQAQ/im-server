package svc

import (
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_message/internal/config"
	"github.com/woxQAQ/im-service/pkg/common/model/single"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	SingleMsg single.MessageDtlModel
	Rds redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	singleConn := sqlx.NewMysql(c.Mysql.DataSourceSingle)
	return &ServiceContext{
		Config: c,
		SingleMsg: single.NewMessageDtlModel(singleConn,c.CacheRedis),
		Rds: *redis.MustNewRedis(c.Redis),
	}
}
