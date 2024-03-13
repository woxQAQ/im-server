package svc

import (
	"github.com/apache/rocketmq-clients/golang/v5"
	"github.com/bwmarrin/snowflake"
	"github.com/woxQAQ/im-service/internal/app/seq/model"
	"github.com/woxQAQ/im-service/internal/app/seq/rpc/internal/config"
	"github.com/woxQAQ/im-service/pkg/common/mq"
)

type ServiceContext struct {
	Config config.Config

	Producer  golang.Producer
	SnowFlake *snowflake.Node

	SessionIdModel model.SessionIdModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// conn := sqlx.NewMysql(c.Mysql.DataSource)
	snow, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}

	rmqConf, err := c.Rmq.GetConf()
	if err != nil {
		panic(err)
	}

	p, err := mq.NewProducer(rmqConf, c.Topic)

	return &ServiceContext{
		Config:    c,
		SnowFlake: snow,
		Producer:  p,
	}
}
