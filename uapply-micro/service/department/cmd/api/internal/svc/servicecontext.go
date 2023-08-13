package svc

import (
	mysqlx "github.com/jmoiron/sqlx"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
	"uapply-micro/common/sqlxUtil"
	"uapply-micro/service/department/cmd/api/internal/config"
	"uapply-micro/service/department/model"
	"uapply-micro/service/interviewer/cmd/rpc/interviewer"
	"uapply-micro/service/user/cmd/rpc/userclient"
)

type ServiceContext struct {
	Config         config.Config
	DepModel       model.DepartmentModel
	RegModel       model.RegisterModel
	InterConfModel model.InterConfModel
	UserRpc        userclient.User
	InterRpc       interviewer.Interviewer
	Mysqlx         *mysqlx.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	db := sqlxUtil.InitMySqlx(c.Mysql.DataSource)
	return &ServiceContext{
		Config:         c,
		DepModel:       model.NewDepartmentModel(conn, c.CacheRedis),
		RegModel:       model.NewRegisterModel(conn, c.CacheRedis),
		InterConfModel: model.NewInterConfModel(conn, c.CacheRedis),
		UserRpc:        userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		InterRpc:       interviewer.NewInterviewer(zrpc.MustNewClient(c.InterRpc)),
		Mysqlx:         db,
	}
}
