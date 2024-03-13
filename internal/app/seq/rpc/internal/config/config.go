package config

import (
	"github.com/woxQAQ/im-service/config"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	service.ServiceConf
	Seq struct {
		Step          int
		DefaultMaxSeq int
	}
	Mysql struct {
		DataSource string
	}
	CacheRedis cache.CacheConf
	Rmq        config.RmqConfig
	Topic      string
	// Topic      string
}
