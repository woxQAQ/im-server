package svc

import (
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_seq/internal/config"
	"github.com/woxQAQ/im-service/pkg/common/model"
	"github.com/woxQAQ/im-service/pkg/common/model/seq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/sharding"
)

type ServiceContext struct {
	Config                      config.Config
	SessionSeqModel             seq.SessionSequenceModel
	SessionMessageSequenceModel seq.SessionMessageSequenceModel
	SessionIdModel              seq.SessionIdModel
	GroupSequenceModel          seq.GroupSequenceModel
	GroupMessageSequenceModel   seq.GroupMessageSequenceModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// conn := sqlx.NewMysql(c.Mysql.DataSource)

	dbSeq, err := gorm.Open(mysql.Open(c.Mysql.DataSource))
	if err != nil {
		panic(err)
	}

	sessionSeqConn, err := model.UseSharding(dbSeq, "session_id", 100, sharding.PKSnowflake, "session_sequence")
	if err != nil {
		panic(err)
	}

	sessionMessageSequenceConn, err := model.UseSharding(dbSeq, "session_id", 100, sharding.PKSnowflake, "session_message_sequence")
	if err != nil {
		panic(err)
	}

	GroupSequenceConn, err := model.UseSharding(dbSeq, "group_id", 100, sharding.PKSnowflake, "group_sequence")
	if err != nil {
		panic(err)
	}

	GroupMessageSequenceConn, err := model.UseSharding(dbSeq, "group_id", 100, sharding.PKSnowflake, "group_message_sequence")
	if err != nil {
		panic(err)
	}

	SessionIdConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:                      c,
		SessionSeqModel:             seq.NewSessionSequenceModel(*sessionSeqConn, c.CacheRedis),
		SessionIdModel:              seq.NewSessionIdModel(SessionIdConn, c.CacheRedis),
		SessionMessageSequenceModel: seq.NewSessionMessageSequenceModel(*sessionMessageSequenceConn, c.CacheRedis),
		GroupMessageSequenceModel:   seq.NewGroupMessageSequenceModel(*GroupMessageSequenceConn, c.CacheRedis),
		GroupSequenceModel:          seq.NewGroupSequenceModel(*GroupSequenceConn, c.CacheRedis),
	}
}
