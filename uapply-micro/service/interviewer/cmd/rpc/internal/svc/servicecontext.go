package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"uapply-micro/service/interviewer/cmd/rpc/internal/config"
	"uapply-micro/service/interviewer/model"
)

type ServiceContext struct {
	InterModel model.InterInfoModel
	LoginModel model.LoginInfoModel
	Config     config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		InterModel: model.NewInterInfoModel(conn, c.CacheRedis),
		LoginModel: model.NewLoginInfoModel(conn, c.CacheRedis),
		Config:     c,
	}
}
