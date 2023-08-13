package logic

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"uapply-micro/service/department/model"

	"uapply-micro/service/department/cmd/rpc/department"
	"uapply-micro/service/department/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetRegisterMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRegisterMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRegisterMsgLogic {
	return &GetRegisterMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRegisterMsgLogic) GetRegisterMsg(in *department.UserUid) (*department.RegInfo, error) {
	userMsg, err := l.svcCtx.RegModel.FindByUid(in.Uid)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "无报名数据")
		}
		return nil, status.Error(codes.Internal, "系统繁忙")
	}

	u := &department.RegInfo{
		DepID:   int32(userMsg.DepId),
		DepName: userMsg.DepName,
		OrgID:   int32(userMsg.OrgId),
		OrgName: userMsg.OrgName,
		First:   int32(userMsg.FirstStatus),
		Second:  int32(userMsg.SecondStatus),
		Final:   int32(userMsg.FirstStatus),
	}
	return u, nil
}
