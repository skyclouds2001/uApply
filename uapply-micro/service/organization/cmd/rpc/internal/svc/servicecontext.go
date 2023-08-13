package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"uapply-micro/service/organization/cmd/rpc/internal/config"
	"uapply-micro/service/organization/model"
)

type ServiceContext struct {
	Config    config.Config
	TimeModel model.EnrollTimeModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		TimeModel: model.NewEnrollTimeModel(conn, c.CacheRedis),
	}
}
