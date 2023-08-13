package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
	"uapply-micro/service/department/cmd/rpc/departmentclient"
	"uapply-micro/service/organization/cmd/rpc/organizationclient"
	"uapply-micro/service/user/cmd/api/internal/config"
	"uapply-micro/service/user/model"
)

type ServiceContext struct {
	Config      config.Config
	OrgRpc      organizationclient.Organization
	DepRpc      departmentclient.Department
	LoginModel  model.LoginInfoModel
	ResumeModel model.UserInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		OrgRpc:      organizationclient.NewOrganization(zrpc.MustNewClient(c.OrgRpc)),
		DepRpc:      departmentclient.NewDepartment(zrpc.MustNewClient(c.DepRpc)),
		LoginModel:  model.NewLoginInfoModel(conn, c.CacheRedis),
		ResumeModel: model.NewUserInfoModel(conn, c.CacheRedis),
	}
}
