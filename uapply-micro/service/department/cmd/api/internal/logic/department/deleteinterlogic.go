package department

import (
	"context"
	"encoding/json"
	"fmt"
	"uapply-micro/common/errorx"
	"uapply-micro/service/department/cmd/api/internal/logic"
	"uapply-micro/service/department/cmd/api/internal/svc"
	"uapply-micro/service/department/cmd/api/internal/types"
	"uapply-micro/service/interviewer/cmd/rpc/inter"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteInterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteInterLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteInterLogic {
	return DeleteInterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteInterLogic) DeleteInter(req types.InterdelReq) (resp *types.Msg, err error) {
	// 拿部门id
	depIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value(logic.DepIdKey)))
	depId, err := depIdNumber.Int64()
	if err != nil {
		logx.Error(err)
		return nil, errorx.NewSysError()
	}
	if depId <= 0 {
		return nil, errorx.NewError("登录信息错误", errorx.Fail)
	}

	in := &inter.InterReq{
		Uid:   int32(req.Uid),
		Depid: depId,
	}
	_, err = l.svcCtx.InterRpc.DeleteInter(l.ctx, in)
	if err != nil {
		return nil, errorx.HandleGrpcErrorToHttp(err)
	}

	return &types.Msg{Msg: "删除成功"}, nil
}
