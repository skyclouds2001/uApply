package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	Mysql struct {
		DataSource string
	}
	CacheRedis cache.CacheConf
	zrpc.RpcServerConf
}
