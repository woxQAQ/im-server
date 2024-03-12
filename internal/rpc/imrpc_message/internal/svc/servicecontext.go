package svc

import (
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_message/internal/config"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_seq/seq"
	"github.com/woxQAQ/im-service/pkg/common/model/message"
	"github.com/woxQAQ/im-service/pkg/common/model/single"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type Session struct {
	Sender   int64
	Receiver int64
}

type ServiceContext struct {
	Config    config.Config
	SingleMsg single.MessageDtlModel
	//Rds redis.Redis

	SenderDtl message.TSenderDtlModel
	// SinglePreMap store map[sessionId]previousId
	SinglePreMap map[Session]int64

	SeqServer seq.Seq
}

func NewServiceContext(c config.Config) *ServiceContext {
	singleConn := sqlx.NewMysql(c.Mysql.DataSourceSingle)
	senderDtlConn := sqlx.NewMysql(c.Mysql.DataSourceSender)
	return &ServiceContext{
		Config:    c,
		SingleMsg: single.NewMessageDtlModel(singleConn, c.CacheRedis),
		//Rds: *redis.MustNewRedis(c.Redis),
		SeqServer: seq.NewSeq(zrpc.MustNewClient(c.SeqRpc)),

		SenderDtl: message.NewTSenderDtlModel(senderDtlConn, c.CacheRedis),
	}
}
