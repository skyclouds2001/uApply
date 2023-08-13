package org

import (
	"context"
	"encoding/json"
	"fmt"
	"uapply-micro/common/errorx"
	"uapply-micro/service/department/cmd/rpc/departmentclient"
	"uapply-micro/service/organization/cmd/api/internal/logic"

	"uapply-micro/service/organization/cmd/api/internal/svc"
	"uapply-micro/service/organization/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAllDepLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllDepLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllDepLogic {
	return GetAllDepLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllDepLogic) GetAllDep() (*types.DepsRsp, error) {
	// 先获取 orgId
	orgIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value(logic.KeyID)))
	orgId, err := orgIdNumber.Int64()
	if err != nil {
		logx.Error(err)
		return nil, errorx.NewSysError()
	}
	if orgId <= 0 {
		return nil, errorx.NewError("登录信息错误", errorx.Fail)
	}

	// 调用rpc
	dep, err := l.svcCtx.DepRpc.GetDeps(l.ctx, &departmentclient.DepReq{
		OrgId: int32(orgId),
	})
	if err != nil {
		logx.Error(err)
		err := errorx.HandleGrpcErrorToHttp(err)
		return nil, err
	}
	var rsp = make([]types.DepRsp, 0)
	var d = types.DepRsp{}
	for _, r := range dep.Deps {
		d.ID = int(r.Id)
		d.Name = r.Name
		d.Account = r.Account
		d.Password = r.Password
		rsp = append(rsp, d)
	}
	return &types.DepsRsp{
		Departments: rsp,
	}, nil
}
