package org

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"time"
	"uapply-micro/common/errorx"
	"uapply-micro/service/organization/cmd/api/internal/logic"
	"uapply-micro/service/organization/model"

	"uapply-micro/service/organization/cmd/api/internal/svc"
	"uapply-micro/service/organization/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddTimeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddTimeLogic {
	return AddTimeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTimeLogic) AddTime(req types.ATimeReq) (*types.MsgRsp, error) {
	// 拿组织id
	orgIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value(logic.KeyID)))
	orgId, err := orgIdNumber.Int64()
	if err != nil {
		logx.Error(err)
		return nil, errorx.NewSysError()
	}
	if orgId <= 0 {
		return nil, errorx.NewError("登录信息错误", errorx.Fail)
	}

	orgNameNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value(logic.KeyName)))
	orgName := orgNameNumber.String()
	// 判断参数
	t := time.Now().Unix()
	if req.Start >= req.End {
		return nil, errorx.NewError("开始时间不能大于结束时间", errorx.ParamInvalid)
	}
	if t >= req.End {
		return nil, errorx.NewError("结束时间不能小于当前时间", errorx.ParamInvalid)
	}
	// 设置报名时间
	_, err = l.svcCtx.TimeModel.Insert(&model.EnrollTime{
		Id:        orgId,
		Name:      orgName,
		StartTime: req.Start,
		EndTime:   req.End,
	})
	if err != nil {
		logx.Error(err)
		if errMySQL, ok := err.(*mysql.MySQLError); ok {
			switch errMySQL.Number {
			case 1062:
				err = l.svcCtx.TimeModel.Update(&model.EnrollTime{
					Id:        orgId,
					Name:      orgName,
					StartTime: req.Start,
					EndTime:   req.End,
				})
				if err != nil {
					return nil, errorx.NewSysError()
				}
				return &types.MsgRsp{
					Msg: "修改招新时间成功",
				}, nil
			default:
				return nil, errorx.NewSysError()
			}
		}
		return nil, errorx.NewSysError()
	}
	return &types.MsgRsp{
		Msg: "添加招新时间成功",
	}, nil
}
