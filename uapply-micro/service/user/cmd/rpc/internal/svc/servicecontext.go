package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"uapply-micro/service/user/cmd/rpc/internal/config"
	"uapply-micro/service/user/model"

	_ "github.com/go-sql-driver/mysql"
	mysqlx "github.com/jmoiron/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	UserModel  model.UserInfoModel
	LoginModel model.LoginInfoModel
	Mysqlx     *mysqlx.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	connect, _ := mysqlx.Connect("mysql", c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		UserModel:  model.NewUserInfoModel(conn, c.CacheRedis),
		LoginModel: model.NewLoginInfoModel(conn, c.CacheRedis),
		Mysqlx:     connect,
	}
}
