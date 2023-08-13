package svc

import (
	mysqlx "github.com/jmoiron/sqlx"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"uapply-micro/common/sqlxUtil"
	"uapply-micro/service/department/cmd/rpc/internal/config"
	"uapply-micro/service/department/model"
)

type ServiceContext struct {
	Config   config.Config
	DepModel model.DepartmentModel
	RegModel model.RegisterModel
	Mysqlx   *mysqlx.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	db := sqlxUtil.InitMySqlx(c.Mysql.DataSource)
	return &ServiceContext{
		Config:   c,
		DepModel: model.NewDepartmentModel(conn, c.CacheRedis),
		RegModel: model.NewRegisterModel(conn, c.CacheRedis),
		Mysqlx:   db,
	}
}
