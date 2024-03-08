package svc

import (
	"github.com/apache/rocketmq-clients/golang/v5"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_seq/internal/config"
	"github.com/woxQAQ/im-service/pkg/common/model/seq"
	"github.com/woxQAQ/im-service/pkg/common/mq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config   config.Config
	SeqModel seq.SequenceModel
	Rds      *redis.Redis
	producer golang.Producer
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	rds := redis.MustNewRedis(c.Redis.RedisConf)

	p, err := mq.NewProducer(&c.Rmq, c.Topic)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:   c,
		SeqModel: seq.NewSequenceModel(conn),
		Rds:      rds,
		producer: p,
	}
}
