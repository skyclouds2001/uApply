package department

import (
	"context"
	"encoding/json"
	"fmt"
	"uapply-micro/common/errorx"
	"uapply-micro/service/department/cmd/api/internal/logic"
	"uapply-micro/service/interviewer/cmd/rpc/inter"

	"uapply-micro/service/department/cmd/api/internal/svc"
	"uapply-micro/service/department/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetIntersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetIntersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIntersLogic {
	return &GetIntersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetIntersLogic) GetInters() (*types.InterInfosResp, error) {
	depIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value(logic.DepIdKey)))
	depId, err := depIdNumber.Int64()
	if err != nil {
		logx.Error(err)
		return nil, errorx.NewSysError()
	}
	if depId <= 0 {
		return nil, errorx.NewError("登录信息错误", errorx.Fail)
	}

	intersResp, err := l.svcCtx.InterRpc.GetInters(l.ctx, &inter.IntersReq{
		DepId: int32(depId),
	})
	if err != nil {
		return nil, errorx.HandleGrpcErrorToHttp(err)
	}

	retInfos := make([]types.InterInfo, 0, len(intersResp.Infos))
	for _, interInfo := range intersResp.Infos {
		retInfos = append(retInfos, types.InterInfo{
			Id:   interInfo.Id,
			Uid:  interInfo.Uid,
			Name: interInfo.Name,
		})
	}

	return &types.InterInfosResp{
		Infos: retInfos,
	}, nil
}
