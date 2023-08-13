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

type GetDepsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDepsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDepsLogic {
	return &GetDepsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDepsLogic) GetDeps(in *department.DepReq) (*department.DepsInfoRsp, error) {
	all, err := l.svcCtx.DepModel.FindAll(int64(in.OrgId))
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "该组织下无部门")
		}
		return nil, status.Error(codes.Internal, "系统繁忙")
	}
	var rsp []*department.DepInfoRsp
	for _, d := range all {
		rsp = append(rsp, &department.DepInfoRsp{
			Id:       int32(d.Id),
			Name:     d.Name,
			Account:  d.Account,
			Password: d.Password,
		})
	}
	return &department.DepsInfoRsp{
		Deps: rsp,
	}, nil
}
