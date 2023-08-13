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

type DelDepLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelDepLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelDepLogic {
	return &DelDepLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelDepLogic) DelDep(in *department.DelReq) (*department.DepInfoRsp, error) {
	// 查询orgid是否对应此depid
	one, err := l.svcCtx.DepModel.FindOne(int64(in.DepId))
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "无此部门信息")
		}
		return nil, status.Error(codes.Internal, "系统繁忙")
	}
	if one.OrgId != int64(in.OrgId) {
		return nil, status.Error(codes.InvalidArgument, "组织下无此部门")
	}
	// 调用删除
	err = l.svcCtx.DepModel.Delete(one, int64(in.DepId))
	if err != nil {
		return nil, status.Error(codes.Internal, "系统繁忙")
	}
	return &department.DepInfoRsp{
		Id:   int32(one.Id),
		Name: one.Name,
	}, nil
}
