package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"uapply-micro/common/errorx"
	"uapply-micro/service/user/cmd/api/internal/logic"
	"uapply-micro/service/user/model"

	"uapply-micro/service/user/cmd/api/internal/svc"
	"uapply-micro/service/user/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CheckResumeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckResumeLogic(ctx context.Context, svcCtx *svc.ServiceContext) CheckResumeLogic {
	return CheckResumeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckResumeLogic) CheckResume() (*types.IfExist, error) {
	// 拿到 uid
	uidNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value(logic.UIDKey)))
	uid, err := uidNumber.Int64()
	if err != nil {
		logx.Error(err)
		return nil, errorx.NewSysError()
	}
	if uid <= 0 {
		return nil, errorx.NewError("登录信息错误", errorx.Fail)
	}

	one, err := l.svcCtx.ResumeModel.FindOne(uid)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return &types.IfExist{Exist: false}, nil
		}
		return nil, errorx.NewSysError()
	}

	return &types.IfExist{
		Exist: true,
		Resume: types.AddResumeReq{
			Name:       one.Name,
			Sex:        one.Sex,
			StudentNum: one.StudentNum,
			AddressNum: one.AddressNum.String,
			Major:      one.Major.String,
			PhoneNum:   one.PhoneNum,
			Email:      one.Email,
			Intro:      one.Intro,
		},
	}, nil
}
