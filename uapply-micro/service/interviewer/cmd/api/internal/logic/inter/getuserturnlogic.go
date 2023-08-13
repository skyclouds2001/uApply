package inter

import (
	"context"
	"uapply-micro/common/errorx"
	"uapply-micro/common/filed"
	"uapply-micro/service/department/cmd/rpc/department"

	"uapply-micro/service/interviewer/cmd/api/internal/svc"
	"uapply-micro/service/interviewer/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

const (
	finalTurn  = iota // 0
	firstTurn         // 1
	secondTurn        // 2
)

type GetUserTurnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserTurnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserTurnLogic {
	return &GetUserTurnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserTurnLogic) GetUserTurn(req *types.TurnReq) (resp *types.TurnResp, err error) {
	regRsp, err := l.svcCtx.DepRpc.GetRegister(l.ctx, &department.RegReq{
		Uid: int32(req.Uid),
	})
	if err != nil {
		return nil, errorx.HandleGrpcErrorToHttp(err)
	}

	var regInfo *department.RegInfo
	for _, reg := range regRsp.Rsp {
		if reg.DepID != int32(req.DepId) {
			continue
		}
		regInfo = reg
		break
	}
	if regInfo == nil {
		return nil, errorx.NewCodeError(errorx.NotFound, "用户不存在")
	}

	resp = &types.TurnResp{}
	if regInfo.Final != 0 {
		resp.Turn = finalTurn
		return
	}
	if (regInfo.First == filed.NEW || regInfo.First == filed.ARRANGED || regInfo.First == filed.INTERVIEWED) && regInfo.Second == filed.NEW {
		resp.Turn = firstTurn
		return
	}
	if (regInfo.First == filed.PASS || regInfo.First == filed.OUT) && (regInfo.Second == filed.NEW || regInfo.Second == filed.ARRANGED || regInfo.Second == filed.INTERVIEWED) {
		resp.Turn = secondTurn
	}
	if regInfo.Second == filed.PASS || regInfo.Second == filed.OUT {
		resp.Turn = finalTurn
	}
	return
}
