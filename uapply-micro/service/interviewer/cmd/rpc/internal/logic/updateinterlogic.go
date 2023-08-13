package logic

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"uapply-micro/service/interviewer/model"

	"uapply-micro/service/interviewer/cmd/rpc/inter"
	"uapply-micro/service/interviewer/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateInterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateInterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInterLogic {
	return &UpdateInterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改面试官
func (l *UpdateInterLogic) UpdateInter(in *inter.UpdateInterReq) (*inter.Empty, error) {
	interInfo, err := l.svcCtx.InterModel.FindOneByUidAndDepId(int64(in.Uid), int64(in.DepId))
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "用户不存在")
		}
		return nil, status.Error(codes.Internal, "系统繁忙")
	}

	interInfo.Name = in.Name
	err = l.svcCtx.InterModel.Update(interInfo)
	if err != nil {
		return nil, status.Error(codes.Internal, "系统繁忙")
	}
	return &inter.Empty{}, nil
}
