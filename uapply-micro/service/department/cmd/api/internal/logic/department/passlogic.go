package department

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"uapply-micro/common/errorx"
	"uapply-micro/common/filed"
	"uapply-micro/service/department/cmd/api/internal/logic"
	"uapply-micro/service/department/cmd/api/internal/svc"
	"uapply-micro/service/department/cmd/api/internal/types"
	"uapply-micro/service/department/model"

	"github.com/tal-tech/go-zero/core/logx"
)

type PassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPassLogic(ctx context.Context, svcCtx *svc.ServiceContext) PassLogic {
	return PassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PassLogic) Pass(req types.UidInfo) (*types.Msg, error) {
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

	// 检验面试轮次
	var num string
	if req.Num == 1 {
		num = filed.First
	} else if req.Num == 2 {
		num = filed.Second
	} else {
		return nil, errorx.NewError("面试只有一轮和二轮", errorx.ParamInvalid)
	}

	err = l.svcCtx.RegModel.UpdateMany(req.UID, int(depId), l.svcCtx.Mysqlx, num, filed.PASS)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errorx.NewSysError()
	}

	return &types.Msg{
		Msg: "提交成功",
	}, nil
}
