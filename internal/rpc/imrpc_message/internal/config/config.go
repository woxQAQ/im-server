package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSourceSingle string
		DataSourceGroup  string
		DataSourceSender string
	}
	CacheRedis cache.CacheConf
	//Redis redis.RedisConf
	SeqRpc zrpc.RpcClientConf
}
