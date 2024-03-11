package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Seq struct {
		Step          int
		DefaultMaxSeq int
	}
	Mysql struct {
		DataSource string
	}
	CacheRedis cache.CacheConf
	// Rmq        golang.Config
	// Topic      string
}
