package svc

import (
	"bytes"
	"sync"

	"github.com/woxQAQ/im-service/internal/rpc/imrpc_seq/internal/config"
	"github.com/woxQAQ/im-service/pkg/common/model"
	"github.com/woxQAQ/im-service/pkg/common/model/seq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/sharding"
)

type ServiceContext struct {
	Config          config.Config
	SessionSeqModel seq.SessionSequenceModel
	Rds             *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	// conn := sqlx.NewMysql(c.Mysql.DataSource)

	dbSession, err := gorm.Open(mysql.Open(c.Mysql.DataSource))
	if err != nil {
		panic(err)
	}

	sessionConn, err := model.UseSharding(dbSession, "session_id", 100, sharding.PKSnowflake, "session_sequence")
	if err != nil {
		panic(err)
	}
	rds := redis.MustNewRedis(c.Redis.RedisConf)

	return &ServiceContext{
		Config:          c,
		SessionSeqModel: seq.NewSessionSequenceModel(*sessionConn, c.CacheRedis),
		Rds:             rds,
	}
}
