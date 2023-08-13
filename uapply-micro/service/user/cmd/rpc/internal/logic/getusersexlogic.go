package logic

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"uapply-micro/service/user/cmd/rpc/internal/svc"
	"uapply-micro/service/user/cmd/rpc/user"
	"uapply-micro/service/user/model"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserSexLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserSexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserSexLogic {
	return &GetUserSexLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserSexLogic) GetUserSex(in *user.GetSexReq) (*user.GetSexRsp, error) {
	var ids []int
	for _, id := range in.Ids {
		ids = append(ids, int(id))
	}
	info, err := l.svcCtx.UserModel.FindSexInfo(ids, l.svcCtx.Mysqlx)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "无uid信息")
		}
		logx.Error(err)
		return nil, status.Error(codes.Internal, "系统繁忙"+err.Error())
	}
	return &user.GetSexRsp{
		Sexs: info,
	}, nil
}
