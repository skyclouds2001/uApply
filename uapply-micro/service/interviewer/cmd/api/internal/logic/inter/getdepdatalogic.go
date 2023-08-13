package inter

import (
	"context"
	"uapply-micro/common/errorx"
	"uapply-micro/common/filed"
	"uapply-micro/service/department/cmd/rpc/department"
	"uapply-micro/service/user/cmd/rpc/user"

	"uapply-micro/service/interviewer/cmd/api/internal/svc"
	"uapply-micro/service/interviewer/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetDepDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDepDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetDepDataLogic {
	return GetDepDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDepDataLogic) GetDepData(req types.DepDataReq) (resp *types.DepDataResp, err error) {
	regs, err := l.svcCtx.DepRpc.GetRegByDepIds(l.ctx, &department.RbdReq{Ids: []int32{int32(req.DepId)}})
	if err != nil {
		return nil, errorx.HandleGrpcErrorToHttp(err)
	}
	if len(regs.Data) < 1 {
		return nil, errorx.NewCodeError(errorx.NotFound, "未找到部门招新数据")
	}

	// register 信息
	regData := regs.Data[0].Dep
	length := len(regData)
	uids := make([]int64, 0, length)
	for _, reg := range regData {
		uids = append(uids, int64(reg.Uid))
	}

	// 没有数据，直接返回
	if len(uids) == 0 {
		return &types.DepDataResp{
			DepId: req.DepId,
		}, nil
	}
	// 性别信息
	sexs, err := l.svcCtx.UserRpc.GetUserSex(l.ctx, &user.GetSexReq{Ids: uids})
	if err != nil {
		return nil, errorx.HandleGrpcErrorToHttp(err)
	}
	if len(sexs.Sexs) != length {
		return nil, errorx.NewCodeError(errorx.Fail, "系统数据出错，请联系后台人员")
	}

	sexData := sexs.Sexs
	resp = &types.DepDataResp{
		DepId: req.DepId,
		Sum:   length,
	}
	// 局部性优化
	var (
		sumMale    = 0
		sumFemale  = 0
		pass       = 0
		passMale   = 0
		passFemale = 0
	)

	for i := 0; i < length; i++ {
		if sexData[i] == filed.MALE {
			sumMale++
			if regData[i].Final == filed.PASS {
				pass++
				passMale++
			}
		}
		if sexData[i] == filed.FEMALE {
			sumFemale++
			if regData[i].Final == filed.PASS {
				pass++
				passFemale++
			}
		}
	}
	resp.SumMale = sumMale
	resp.SumFemale = sumFemale
	resp.Pass = pass
	resp.PassMale = passMale
	resp.PassFemale = passFemale

	return
}
