package logic

import (
	"context"

	"uapply-micro/service/department/cmd/rpc/department"
	"uapply-micro/service/department/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetRegByDepIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRegByDepIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRegByDepIdsLogic {
	return &GetRegByDepIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRegByDepIdsLogic) GetRegByDepIds(in *department.RbdReq) (*department.RbdRsp, error) {
	rsp := make([]*department.EachDep, 0)
	for _, id := range in.Ids {
		all, _ := l.svcCtx.RegModel.FindDepAll(int64(id))
		reg := make([]*department.RegData, 0)
		for _, register := range all {
			reg = append(reg, &department.RegData{
				Uid:    int32(register.Uid),
				First:  int32(register.FirstStatus),
				Second: int32(register.SecondStatus),
				Final:  int32(register.FinalStatus),
			})
		}
		rsp = append(rsp, &department.EachDep{Dep: reg})
	}
	return &department.RbdRsp{
		Data: rsp,
	}, nil
}
