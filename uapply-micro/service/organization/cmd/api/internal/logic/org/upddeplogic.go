package org

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"uapply-micro/common/errorx"
	"uapply-micro/service/department/cmd/rpc/departmentclient"
	"uapply-micro/service/organization/cmd/api/internal/logic"

	"uapply-micro/service/organization/cmd/api/internal/svc"
	"uapply-micro/service/organization/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdDepLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdDepLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdDepLogic {
	return UpdDepLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdDepLogic) UpdDep(req types.UdpReq) (*types.DepRsp, error) {
	orgIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value(logic.KeyID)))
	orgId, err := orgIdNumber.Int64()
	if err != nil {
		logx.Error(err)
		return nil, errorx.NewSysError()
	}
	if orgId <= 0 {
		return nil, errorx.NewError("登录信息错误", errorx.Fail)
	}

	id, _ := strconv.Atoi(req.Id)
	dep, err := l.svcCtx.DepRpc.UpdDep(l.ctx, &departmentclient.UdpReq{
		Id:       int32(id),
		OrgId:    int32(orgId),
		Name:     req.Name,
		Account:  req.Account,
		Password: req.Password,
	})
	if err != nil {
		err = errorx.HandleGrpcErrorToHttp(err)
		return nil, err
	}
	return &types.DepRsp{
		ID:       id,
		Name:     dep.Name,
		Account:  dep.Account,
		Password: dep.Password,
	}, nil
}
