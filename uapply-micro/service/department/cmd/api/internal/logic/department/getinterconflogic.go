package department

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"uapply-micro/common/errorx"
	"uapply-micro/service/department/cmd/api/internal/logic"
	"uapply-micro/service/department/model"

	"uapply-micro/service/department/cmd/api/internal/svc"
	"uapply-micro/service/department/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetInterConfLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInterConfLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetInterConfLogic {
	return GetInterConfLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInterConfLogic) GetInterConf(req types.GetInterConfReq) (resp *types.GetInterConfResp, err error) {
	if req.Turn < 1 || req.Turn > 2 {
		return nil, errorx.NewCodeError(errorx.ParamInvalid, "轮次错误")
	}

	// 部门 id
	depIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value(logic.DepIdKey)))
	depId, err := depIdNumber.Int64()
	if err != nil {
		logx.Error(err)
		return nil, errorx.NewSysError()
	}
	if depId <= 0 {
		return nil, errorx.NewError("登录信息错误", errorx.Fail)
	}

	interConf, err := l.svcCtx.InterConfModel.FindOneByDepIdTurn(depId, int64(req.Turn))
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errorx.NewCodeError(errorx.NotFound, "未找到面试配置信息")
		}
		return nil, errorx.NewSysError()
	}

	resp = &types.GetInterConfResp{
		DepId:        depId,
		Turn:         req.Turn,
		Start:        interConf.Start,
		End:          interConf.End,
		InterAddress: interConf.InterAddress,
		ContactPhone: interConf.ContactPhone,
	}
	return
}
