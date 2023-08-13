package department

import (
	"context"

	"uapply-micro/service/department/cmd/api/internal/svc"
	"uapply-micro/service/department/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type Announce2InterviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAnnounce2InterviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) Announce2InterviewLogic {
	return Announce2InterviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Announce2InterviewLogic) Announce2Interview(req types.AInterReq) (*types.AinterRsp, error) {
	// todo: add your logic here and delete this line

	return &types.AinterRsp{}, nil
}
