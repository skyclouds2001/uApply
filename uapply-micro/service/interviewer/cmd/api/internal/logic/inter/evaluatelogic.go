package inter

import (
	"context"
	"log"
	"uapply-micro/common/errorx"
	"uapply-micro/common/filed"
	"uapply-micro/service/department/cmd/rpc/department"

	"uapply-micro/service/interviewer/cmd/api/internal/svc"
	"uapply-micro/service/interviewer/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type EvaluateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEvaluateLogic(ctx context.Context, svcCtx *svc.ServiceContext) EvaluateLogic {
	return EvaluateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EvaluateLogic) Evaluate(req types.MarkReq) (resp *types.MsgRsq, err error) {
	num := req.Num
	if num > 2 || num < 1 {
		return nil, errorx.NewError("轮次错误", errorx.ParamInvalid)
	}

	uid := req.UID
	in := &department.RegReq{Uid: int32(uid)}
	log.Println(in.Uid)
	regMsg, err := l.svcCtx.DepRpc.GetRegister(l.ctx, in)
	if err != nil {
		return nil, errorx.HandleGrpcErrorToHttp(err)
	}
	if regMsg == nil {
		return nil, errorx.NewError("用户没有注册", errorx.NotFound)
	}

	for _, info := range regMsg.Rsp {
		if info.DepID != int32(req.DepId) {
			continue
		}
		if info.Final == filed.OUT {
			return nil, errorx.NewError("用户已经被淘汰了", errorx.ParamInvalid)
		}
		u := &department.MarkInfo{
			DepId: int64(info.DepID),
			Mark:  int32(req.Mark),
			Eva:   req.Eva,
			Uid:   int32(uid),
		}
		switch num {
		case 1:
			if info.FirstMark != 0 || info.SecondMark != 0 || info.Final == filed.PASS {
				return nil, errorx.NewError("用户无需参加此轮面试", errorx.ParamInvalid)
			}
			u.MarkTag = filed.First_Mark
			u.EvaTag = filed.First_Eva
		case 2:
			if info.First != filed.PASS {
				return nil, errorx.NewError("用户没有参加第一次面试", errorx.ParamInvalid)
			}
			if info.SecondMark != 0 || info.Final == filed.PASS {
				return nil, errorx.NewError("用户无需参加此轮面试", errorx.ParamInvalid)
			}
			u.MarkTag = filed.Second_Mark
			u.EvaTag = filed.Second_Eva
		}

		_, err = l.svcCtx.DepRpc.MarkUser(l.ctx, u)
		if err != nil {
			return nil, errorx.HandleGrpcErrorToHttp(err)
		}
		return &types.MsgRsq{
			Msg: "评价成功",
		}, nil
	}

	return nil, errorx.NewError("面试官没有权限", errorx.NotAuth)
}
