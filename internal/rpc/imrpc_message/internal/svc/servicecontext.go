package svc

import (
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_message/internal/config"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_seq/seq"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_session/session"
	"github.com/woxQAQ/im-service/pkg/common/model/single"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	SingleMsg single.MessageDtlModel
	//Rds redis.Redis

	// PreMap store map[sessionId]previousId
	PreMap map[string]int64

	SeqServer seq.Seq

	SessionServer session.Session
}

func NewServiceContext(c config.Config) *ServiceContext {
	singleConn := sqlx.NewMysql(c.Mysql.DataSourceSingle)
	return &ServiceContext{
		Config:    c,
		SingleMsg: single.NewMessageDtlModel(singleConn, c.CacheRedis),
		//Rds: *redis.MustNewRedis(c.Redis),
		SeqServer:     seq.NewSeq(zrpc.MustNewClient(c.SeqRpc)),
		SessionServer: session.NewSession(zrpc.MustNewClient(c.SessionRpc)),
	}
}
