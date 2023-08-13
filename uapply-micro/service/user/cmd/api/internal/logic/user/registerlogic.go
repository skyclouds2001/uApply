package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"uapply-micro/common/errorx"
	"uapply-micro/service/department/cmd/rpc/departmentclient"
	"uapply-micro/service/user/cmd/api/internal/logic"
	"uapply-micro/service/user/model"

	"uapply-micro/service/user/cmd/api/internal/svc"
	"uapply-micro/service/user/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegReq) (*types.RegRsp, error) {
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

	// 先看用户有没有填简历
	_, err = l.svcCtx.ResumeModel.FindOne(uid)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errorx.NewCodeError(errorx.Fail, "未填写简历")
		}
		return nil, errorx.NewSysError()
	}

	_, err = l.svcCtx.DepRpc.Register(l.ctx, &departmentclient.RegReq{
		Uid:   int32(uid),
		OrgId: int32(req.OrgId),
		DepId: int32(req.DepId),
	})

	if err != nil {
		err := errorx.HandleGrpcErrorToHttp(err)
		return nil, err
	}

	return &types.RegRsp{
		Msg: "报名成功",
	}, nil
}
