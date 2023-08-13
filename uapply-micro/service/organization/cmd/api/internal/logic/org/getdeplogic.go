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

type GetDepLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDepLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetDepLogic {
	return GetDepLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDepLogic) GetDep(req types.IdReq) (*types.DepRsp, error) {
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
	id, _ := strconv.Atoi(req.Id)
	dep, err := l.svcCtx.DepRpc.GetDep(l.ctx, &departmentclient.DepReq{
		OrgId: int32(orgId),
		Id:    []int32{int32(id)},
	})
	if err != nil {
		logx.Error(err)
		err := errorx.HandleGrpcErrorToHttp(err)
		return nil, err
	}
	return &types.DepRsp{
		Name:     dep.Name,
		Account:  dep.Account,
		Password: dep.Password,
		ID:       int(dep.Id),
	}, nil
}
