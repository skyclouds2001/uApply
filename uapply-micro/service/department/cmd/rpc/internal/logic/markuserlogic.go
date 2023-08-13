package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"uapply-micro/service/department/cmd/rpc/department"
	"uapply-micro/service/department/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type MarkUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMarkUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarkUserLogic {
	return &MarkUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MarkUserLogic) MarkUser(in *department.MarkInfo) (*department.Empty, error) {
	uid := in.Uid
	err := l.svcCtx.RegModel.UpdateMark(int64(uid), in.DepId, in.Mark, in.MarkTag, in.EvaTag, in.Eva)
	if err != nil {
		return nil, status.Error(codes.Internal, "系统繁忙")
	}
	return &department.Empty{}, nil
}
