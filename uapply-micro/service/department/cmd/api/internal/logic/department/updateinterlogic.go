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

type UpdateInterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateInterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInterLogic {
	return &UpdateInterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateInterLogic) UpdateInter(req *types.UpdateInterReq) (*types.UpdateInterRsp, error) {
	// 获取 depId
	depIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value(logic.DepIdKey)))
	depId, err := depIdNumber.Int64()
	if err != nil {
		logx.Error(err)
		return nil, errorx.NewSysError()
	}
	if depId <= 0 {
		return nil, errorx.NewError("登录信息错误", errorx.Fail)
	}

	_, err = l.svcCtx.InterRpc.UpdateInter(l.ctx, &inter.UpdateInterReq{
		Uid:   int32(req.Uid),
		DepId: int32(depId),
		Name:  req.Name,
	})
	if err != nil {
		return nil, errorx.HandleGrpcErrorToHttp(err)
	}

	return &types.UpdateInterRsp{
		Msg: "修改成功",
	}, nil
}
