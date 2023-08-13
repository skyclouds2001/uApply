package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
	"uapply-micro/service/department/cmd/rpc/departmentclient"
	"uapply-micro/service/organization/cmd/api/internal/config"
	"uapply-micro/service/organization/model"
	"uapply-micro/service/user/cmd/rpc/userclient"
)

type ServiceContext struct {
	Config    config.Config
	OrgModel  model.OrganizationModel
	TimeModel model.EnrollTimeModel
	DepRpc    departmentclient.Department
	UserRpc   userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		OrgModel:  model.NewOrganizationModel(conn, c.CacheRedis),
		TimeModel: model.NewEnrollTimeModel(conn, c.CacheRedis),
		DepRpc:    departmentclient.NewDepartment(zrpc.MustNewClient(c.DepRpc)),
		UserRpc:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
