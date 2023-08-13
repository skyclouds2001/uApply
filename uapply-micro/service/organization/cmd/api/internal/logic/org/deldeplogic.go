package org

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"uapply-micro/common/errorx"
	"uapply-micro/service/department/cmd/rpc/department"
	"uapply-micro/service/organization/cmd/api/internal/logic"

	"uapply-micro/service/organization/cmd/api/internal/svc"
	"uapply-micro/service/organization/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DelDepLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelDepLogic(ctx context.Context, svcCtx *svc.ServiceContext) DelDepLogic {
	return DelDepLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelDepLogic) DelDep(req types.IdReq) (*types.DepRsp, error) {
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
	dep, err := l.svcCtx.DepRpc.DelDep(l.ctx, &department.DelReq{
		OrgId: int32(orgId),
		DepId: int32(id),
	})
	if err != nil {
		err := errorx.HandleGrpcErrorToHttp(err)
		return nil, err
	}
	return &types.DepRsp{
		Name: dep.Name,
		ID:   int(dep.Id),
	}, nil
}
