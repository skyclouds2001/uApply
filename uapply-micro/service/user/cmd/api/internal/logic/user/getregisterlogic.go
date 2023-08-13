package user

import (
	"context"
	"encoding/json"
	"fmt"
	"uapply-micro/common/errorx"
	"uapply-micro/service/department/cmd/rpc/department"
	"uapply-micro/service/user/cmd/api/internal/logic"
	"uapply-micro/service/user/cmd/api/internal/svc"
	"uapply-micro/service/user/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetRegisterLogic {
	return GetRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRegisterLogic) GetRegister() (*types.GetRegRsp, error) {
	// 获取uid
	uidNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value(logic.UIDKey)))
	uid, err := uidNumber.Int64()
	if err != nil {
		logx.Error(err)
		return nil, errorx.NewSysError()
	}
	if uid <= 0 {
		return nil, errorx.NewError("登录信息错误", errorx.Fail)
	}

	res, err := l.svcCtx.DepRpc.GetRegister(l.ctx, &department.RegReq{
		Uid: int32(uid),
	})
	if err != nil {
		return nil, errorx.HandleGrpcErrorToHttp(err)
	}
	var rsp []types.ODInfo
	for _, od := range res.Rsp {
		rsp = append(rsp, types.ODInfo{
			DepID:   int(od.DepID),
			DepName: od.DepName,
			OrgID:   int(od.OrgID),
			OrgName: od.OrgName,
			First:   int(od.First),
			Second:  int(od.Second),
			Final:   int(od.Final),
		})
	}
	return &types.GetRegRsp{
		Rsp: rsp,
	}, nil
}
