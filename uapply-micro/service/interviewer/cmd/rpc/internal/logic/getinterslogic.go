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

type GetIntersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIntersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIntersLogic {
	return &GetIntersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取所有面试官
func (l *GetIntersLogic) GetInters(in *inter.IntersReq) (*inter.IntersResp, error) {
	interModels, err := l.svcCtx.InterModel.FindByDepId(int64(in.DepId))
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "未找到面试官信息")
		}
		return nil, status.Error(codes.Internal, "系统繁忙")
	}

	retInfos := make([]*inter.InterInfo, 0, len(interModels))
	for _, interModel := range interModels {
		retInfos = append(retInfos, &inter.InterInfo{
			Id:   interModel.Id,
			Uid:  interModel.Uid,
			Name: interModel.Name,
		})
	}

	return &inter.IntersResp{
		Infos: retInfos,
	}, nil
}
