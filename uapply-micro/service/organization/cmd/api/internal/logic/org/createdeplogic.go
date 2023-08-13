package org

import (
	"context"
	"encoding/json"
	"fmt"
	"uapply-micro/common/errorx"
	"uapply-micro/common/password"
	"uapply-micro/service/department/cmd/rpc/departmentclient"
	"uapply-micro/service/organization/cmd/api/internal/logic"

	"uapply-micro/service/organization/cmd/api/internal/svc"
	"uapply-micro/service/organization/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateDepLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDepLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateDepLogic {
	return CreateDepLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CreateDep 创建部门
func (l *CreateDepLogic) CreateDep(req types.RegReq) (*types.CommonReply, error) {
	orgIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value(logic.KeyID)))
	orgId, err := orgIdNumber.Int64()
	if err != nil {
		logx.Error(err)
		return nil, errorx.NewSysError()
	}
	if orgId <= 0 {
		return nil, errorx.NewError("登录信息错误", errorx.Fail)
	}

	orgNameNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value(logic.KeyName)))
	orgName := orgNameNumber.String()
	if orgName == "" {
		return nil, errorx.NewError("登录信息错误", errorx.Fail)
	}

	// 检查密码是否合法
	ok := password.VerifyPwd(req.Password, 0)
	if !ok {
		return nil, errorx.NewError("密码不合法，请检查并修改", errorx.PasswordInvalid)
	}
	_, err = l.svcCtx.DepRpc.CreateDep(l.ctx, &departmentclient.CreReq{
		Name:     req.Name,
		Account:  req.Account,
		Password: req.Password,
		OrgId:    int32(orgId),
		OrgName:  orgName,
	})
	if err != nil {
		logx.Error(err)
		err = errorx.HandleGrpcErrorToHttp(err)
		return nil, err
	}
	return &types.CommonReply{
		Name: req.Name,
		Id:   orgId,
	}, nil
}
