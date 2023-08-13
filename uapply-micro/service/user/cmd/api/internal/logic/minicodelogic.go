package logic

import (
	"context"

	"uapply-micro/service/user/cmd/api/internal/svc"
	"uapply-micro/service/user/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MiniCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMiniCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) MiniCodeLogic {
	return MiniCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MiniCodeLogic) MiniCode() (resp *types.MiniCodeResp, err error) {
	resp = &types.MiniCodeResp{
		Code: 0,
	}

	return
}
