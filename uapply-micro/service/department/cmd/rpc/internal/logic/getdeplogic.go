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

type GetDepLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDepLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDepLogic {
	return &GetDepLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDepLogic) GetDep(in *department.DepReq) (*department.DepInfoRsp, error) {
	// 检查部门id
	if len(in.Id) != 1 {
		return nil, status.Error(codes.InvalidArgument, "请检查id参数接口")
	}
	// 数据库搜索
	dep, err := l.svcCtx.DepModel.FindOne(int64(in.Id[0]))
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.InvalidArgument, "组织下无此部门")
		}
		return nil, status.Error(codes.Internal, "系统繁忙")
	}
	return &department.DepInfoRsp{
		Name:     dep.Name,
		Account:  dep.Account,
		Password: dep.Password,
		Id:       int32(dep.Id),
	}, nil
}
