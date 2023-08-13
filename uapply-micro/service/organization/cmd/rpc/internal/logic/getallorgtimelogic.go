package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
	"uapply-micro/service/organization/cmd/rpc/internal/svc"
	"uapply-micro/service/organization/cmd/rpc/organization"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAllOrgTimeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllOrgTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllOrgTimeLogic {
	return &GetAllOrgTimeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllOrgTimeLogic) GetAllOrgTime(in *organization.Empty) (*organization.TimeRsp, error) {
	all, err := l.svcCtx.TimeModel.FindAll()
	if err != nil {
		return nil, status.Error(codes.Internal, "系统繁忙")
	}
	timeInfo := make([]*organization.Time, 0)
	for _, one := range all {
		t := time.Now().Unix()
		if one.EndTime > t && one.StartTime <= t {
			timeInfo = append(timeInfo, &organization.Time{
				Orgid:   int32(one.Id),
				OrgName: one.Name,
				Start:   one.StartTime,
				End:     one.EndTime,
			})
		}
	}
	return &organization.TimeRsp{TimeInfo: timeInfo}, nil
}
