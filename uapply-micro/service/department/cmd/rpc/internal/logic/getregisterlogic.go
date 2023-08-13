package logic

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"uapply-micro/service/department/model"

	"uapply-micro/service/department/cmd/rpc/department"
	"uapply-micro/service/department/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRegisterLogic {
	return &GetRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRegisterLogic) GetRegister(in *department.RegReq) (*department.GetRegRsp, error) {
	log.Println(in.Uid)
	all, err := l.svcCtx.RegModel.FindAll(int64(in.Uid))
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "无报名数据")
		}
		return nil, status.Error(codes.Internal, "系统繁忙")
	}
	if len(all) == 0 {
		return nil, status.Error(codes.NotFound, "无报名数据")
	}

	var rsp []*department.RegInfo
	for _, register := range all {
		rsp = append(rsp, &department.RegInfo{
			DepID:      int32(register.DepId),
			DepName:    register.DepName,
			OrgID:      int32(register.OrgId),
			OrgName:    register.OrgName,
			First:      int32(register.FirstStatus),
			Second:     int32(register.SecondStatus),
			Final:      int32(register.FinalStatus),
			FirstMark:  int32(register.FirstMark),
			FirstEva:   register.FirstEva.String,
			SecondEva:  register.SecondEva.String,
			SecondMark: int32(register.SecondMark),
		})
	}
	return &department.GetRegRsp{
		Rsp: rsp,
	}, nil
}
