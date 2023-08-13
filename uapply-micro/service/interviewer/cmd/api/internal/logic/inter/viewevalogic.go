package inter

import (
	"context"
	"uapply-micro/common/errorx"
	"uapply-micro/service/department/cmd/rpc/department"
	"uapply-micro/service/interviewer/cmd/api/internal/svc"
	"uapply-micro/service/interviewer/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ViewEvaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewViewEvaLogic(ctx context.Context, svcCtx *svc.ServiceContext) ViewEvaLogic {
	return ViewEvaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ViewEvaLogic) ViewEva(req types.GetReq) (resp *types.EvaRsq, err error) {
	num := req.Num
	if num < 1 || num > 2 {
		return nil, errorx.NewError("轮次无效", errorx.ParamInvalid)
	}

	uid := req.UID
	in := &department.RegReq{Uid: int32(uid)}
	reg, err := l.svcCtx.DepRpc.GetRegister(l.ctx, in)
	if err != nil {
		return nil, errorx.HandleGrpcErrorToHttp(err)
	}
	if reg == nil {
		return nil, errorx.NewError("用户没有注册", errorx.NotFound)
	}

	for _, info := range reg.Rsp {
		if info.DepID != int32(req.DepId) {
			continue
		}

		var u types.EvaRsq
		switch num {
		case 1:
			if info.FirstMark == 0 {
				return nil, errorx.NewError("用户还没有参加第一轮面试", errorx.ParamInvalid)
			}
			u.Eva = info.FirstEva
			u.Mark = int(info.FirstMark)
		case 2:
			if info.SecondMark == 0 {
				return nil, errorx.NewError("用户还没有参加第二轮面试", errorx.ParamInvalid)
			}
			u.Eva = info.SecondEva
			u.Mark = int(info.SecondMark)
		}

		return &u, nil
	}

	return nil, errorx.NewError("面试官没有权限", errorx.NotAuth)
}
