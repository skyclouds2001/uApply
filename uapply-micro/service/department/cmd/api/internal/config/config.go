package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Mysql struct {
		DataSource string
	}
	Sms struct {
		SecretId  string
		SecretKey string
		Region    string
		AppId     string
		SignName  string
		Templates struct {
			First  string
			Second string
			Enroll string
			Out    string
		}
	}
	UserRpc    zrpc.RpcClientConf
	InterRpc   zrpc.RpcClientConf
	CacheRedis cache.CacheConf
	Redis      redis.RedisConf
}
