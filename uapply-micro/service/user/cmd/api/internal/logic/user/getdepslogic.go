package user

import (
	"context"
	"strconv"
	"uapply-micro/common/errorx"
	"uapply-micro/service/department/cmd/rpc/department"

	"uapply-micro/service/user/cmd/api/internal/svc"
	"uapply-micro/service/user/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetDepsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDepsLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetDepsLogic {
	return GetDepsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDepsLogic) GetDeps(req types.GetDepsReq) (*types.GetDepsRsp, error) {
	// 拿到组织id
	orgId, _ := strconv.Atoi(req.Id)
	// 调用部门微服务拿到组织下全部的部门信息
	deps, err := l.svcCtx.DepRpc.GetDeps(l.ctx, &department.DepReq{OrgId: int32(orgId)})
	if err != nil {
		return nil, errorx.HandleGrpcErrorToHttp(err)
	}
	var rsp = make([]types.DepInfo, 0)
	for _, dep := range deps.Deps {
		rsp = append(rsp, types.DepInfo{
			DepID:   int(dep.Id),
			DepName: dep.Name,
		})
	}
	return &types.GetDepsRsp{
		Deps: rsp,
	}, nil
}
