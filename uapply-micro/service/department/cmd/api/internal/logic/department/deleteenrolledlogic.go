package department

import (
	"context"
	"encoding/json"
	"fmt"
	"uapply-micro/common/errorx"
	"uapply-micro/service/department/cmd/api/internal/logic"

	"uapply-micro/service/department/cmd/api/internal/svc"
	"uapply-micro/service/department/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteEnrolledLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteEnrolledLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteEnrolledLogic {
	return DeleteEnrolledLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteEnrolledLogic) DeleteEnrolled(req types.DeleteEnrolledReq) (resp *types.Msg, err error) {
	depIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value(logic.DepIdKey)))
	depId, err := depIdNumber.Int64()
	if err != nil {
		logx.Error(err)
		return nil, errorx.NewSysError()
	}
	if depId <= 0 {
		return nil, errorx.NewError("登录信息错误", errorx.Fail)
	}

	err = l.svcCtx.RegModel.DeleteByDepAndUids(depId, req.Uids, l.svcCtx.Mysqlx)
	if err != nil {
		return nil, errorx.NewSysError()
	}

	resp = &types.Msg{Msg: "删除成功"}
	return
}
