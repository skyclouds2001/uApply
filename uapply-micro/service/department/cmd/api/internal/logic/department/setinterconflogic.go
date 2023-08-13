package department

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"uapply-micro/common/errorx"
	"uapply-micro/service/department/cmd/api/internal/logic"
	"uapply-micro/service/department/model"

	"uapply-micro/service/department/cmd/api/internal/svc"
	"uapply-micro/service/department/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SetInterConfLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetInterConfLogic(ctx context.Context, svcCtx *svc.ServiceContext) SetInterConfLogic {
	return SetInterConfLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetInterConfLogic) SetInterConf(req types.SetInterConfReq) (resp *types.Msg, err error) {
	// 检查参数
	if err := checkSetInterConfParam(&req); err != nil {
		return nil, err
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
		// 未找到，插入
		if errors.Is(err, model.ErrNotFound) {
			_, err := l.svcCtx.InterConfModel.Insert(&model.InterConf{
				DepId:        depId,
				Turn:         int64(req.Turn),
				Start:        req.Start,
				End:          req.End,
				InterAddress: req.InterAddress,
				ContactPhone: req.ContactPhone,
			})
			if err != nil {
				return nil, errorx.NewSysError()
			}
			return &types.Msg{Msg: "设置成功"}, nil
		}
		return nil, errorx.NewSysError()
	}

	interConf.Start = req.Start
	interConf.End = req.End
	interConf.InterAddress = req.InterAddress
	interConf.ContactPhone = req.ContactPhone
	err = l.svcCtx.InterConfModel.Update(interConf)
	if err != nil {
		return nil, errorx.NewSysError()
	}

	resp = &types.Msg{Msg: "设置成功"}
	return
}

func checkSetInterConfParam(req *types.SetInterConfReq) error {
	if req.Turn < 1 || req.Turn > 2 {
		return errorx.NewCodeError(errorx.ParamInvalid, "轮次错误")
	}

	if req.Start < 0 || req.End < 0 {
		return errorx.NewCodeError(errorx.ParamInvalid, "时间戳错误")
	}

	if req.InterAddress == "" || req.ContactPhone == "" {
		return errorx.NewCodeError(errorx.ParamInvalid, "未填必要参数")
	}

	if !validatePhone(req.ContactPhone) {
		return errorx.NewCodeError(errorx.ParamInvalid, "电话格式不正确")
	}
	return nil
}

const PhoneReg = `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`

func validatePhone(phone string) bool {
	matched, _ := regexp.MatchString(PhoneReg, phone)
	return matched
}
