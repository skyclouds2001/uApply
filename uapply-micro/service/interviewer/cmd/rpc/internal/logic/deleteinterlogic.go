package logic

import (
	"context"
	"errors"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"uapply-micro/service/interviewer/cmd/rpc/inter"
	"uapply-micro/service/interviewer/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteInterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteInterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteInterLogic {
	return &DeleteInterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  删除面试官
func (l *DeleteInterLogic) DeleteInter(in *inter.InterReq) (*inter.Empty, error) {
	uid := in.Uid
	depid := in.Depid
	err := l.svcCtx.InterModel.DeleteByUidAndDepId(uid, depid)
	if err != nil {
		log.Println(err)
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "没有面试官uid")
		}
		return nil, status.Error(codes.Internal, "系统繁忙")
	}
	return &inter.Empty{}, nil
}
