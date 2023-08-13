package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
	"uapply-micro/service/department/cmd/rpc/departmentclient"
	"uapply-micro/service/interviewer/cmd/api/internal/config"
	"uapply-micro/service/interviewer/model"
	"uapply-micro/service/user/cmd/rpc/userclient"
)

type ServiceContext struct {
	Config    config.Config
	InterInfo model.InterInfoModel
	LoginInfo model.LoginInfoModel
	DepRpc    departmentclient.Department
	UserRpc   userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		InterInfo: model.NewInterInfoModel(conn, c.CacheRedis),
		LoginInfo: model.NewLoginInfoModel(conn, c.CacheRedis),
		DepRpc:    departmentclient.NewDepartment(zrpc.MustNewClient(c.DepRpc)),
		UserRpc:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
