package logic

import (
	"context"
	"uapply-micro/common/errorx"
	"uapply-micro/service/organization/cmd/rpc/organization"

	"uapply-micro/service/user/cmd/api/internal/svc"
	"uapply-micro/service/user/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetOrgEnrollingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrgEnrollingLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetOrgEnrollingLogic {
	return GetOrgEnrollingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrgEnrollingLogic) GetOrgEnrolling() (*types.EnrollingRsp, error) {
	all, err := l.svcCtx.OrgRpc.GetAllOrgTime(l.ctx, &organization.Empty{})
	if err != nil {
		err = errorx.HandleGrpcErrorToHttp(err)
		if err != nil {
			return nil, err
		}
	}
	var rsp = make([]types.Enrolling, 0)
	for _, time := range all.TimeInfo {
		rsp = append(rsp, types.Enrolling{
			OrgId:     int(time.Orgid),
			OrgName:   time.OrgName,
			StartTime: time.Start,
			EndTime:   time.End,
		})
	}
	return &types.EnrollingRsp{
		Organizations: rsp,
	}, nil
}
