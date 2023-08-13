package org

import (
	"context"
	"errors"
	"strconv"
	"uapply-micro/service/organization/model"

	"uapply-micro/common/errorx"
	"uapply-micro/service/organization/cmd/api/internal/svc"
	"uapply-micro/service/organization/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetEnrollTimeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEnrollTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEnrollTimeLogic {
	return &GetEnrollTimeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEnrollTimeLogic) GetEnrollTime(req *types.IdReq) (*types.TimeResp, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil || id < 0 {
		logx.Error(err)
		return nil, errorx.NewError("参数错误", errorx.ParamInvalid)
	}

	enrollTime, err := l.svcCtx.TimeModel.FindOne(int64(id))
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errorx.NewError("参数错误", errorx.NotFound)
		}
		return nil, errorx.NewSysError()
	}

	return &types.TimeResp{
		Start: enrollTime.StartTime,
		End:   enrollTime.EndTime,
	}, nil
}
