package svc

import (
	"github.com/woxQAQ/im-service/internal/app/message/model"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_message/internal/config"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_seq/seq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type Session struct {
	Sender   int64
	Receiver int64
}

type ServiceContext struct {
	Config    config.Config
	SingleMsg model.MessageDtlModel
	//Rds redis.Redis

	SenderDtl model.TSenderDtlModel
	// SinglePreMap store map[sessionId]previousId
	SinglePreMap map[Session]int64

	SeqServer seq.Seq
}

func NewServiceContext(c config.Config) *ServiceContext {
	singleConn := sqlx.NewMysql(c.Mysql.DataSourceSingle)
	senderDtlConn := sqlx.NewMysql(c.Mysql.DataSourceSender)
	return &ServiceContext{
		Config:    c,
		SingleMsg: model.NewMessageDtlModel(singleConn, c.CacheRedis),
		//Rds: *redis.MustNewRedis(c.Redis),
		SeqServer: seq.NewSeq(zrpc.MustNewClient(c.SeqRpc)),

		SenderDtl: model.NewTSenderDtlModel(senderDtlConn, c.CacheRedis),
	}
}
