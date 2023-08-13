package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}

	App struct {
		Appid  string
		Secret string
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	CacheRedis cache.CacheConf
	DepRpc     zrpc.RpcClientConf
	UserRpc    zrpc.RpcClientConf
}
