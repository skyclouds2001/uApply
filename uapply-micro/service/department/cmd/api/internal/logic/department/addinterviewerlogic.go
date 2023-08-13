package department

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"uapply-micro/common/errorx"
	"uapply-micro/service/department/cmd/api/internal/logic"
	"uapply-micro/service/department/model"
	"uapply-micro/service/interviewer/cmd/rpc/interviewer"

	"uapply-micro/service/department/cmd/api/internal/svc"
	"uapply-micro/service/department/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ADDInterviewerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewADDInterviewerLogic(ctx context.Context, svcCtx *svc.ServiceContext) ADDInterviewerLogic {
	return ADDInterviewerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ADDInterviewer 添加面试官
func (l *ADDInterviewerLogic) ADDInterviewer(req types.AddInterReq) (*types.AddInterRsp, error) {
	// 拿到部门的id和信息
	// 拿部门id
	depIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value(logic.DepIdKey)))
	depId, err := depIdNumber.Int64()
	if err != nil {
		logx.Error(err)
		return nil, errorx.NewSysError()
	}
	if depId <= 0 {
		return nil, errorx.NewError("登录信息错误", errorx.Fail)
	}
	// 拿部门信息
	dep, err := l.svcCtx.DepModel.FindOne(depId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, nil
		}
		return nil, errorx.NewSysError()
	}
	// rpc请求添加面试官
	_, err = l.svcCtx.InterRpc.AddInter(l.ctx, &interviewer.AddReq{
		Uid:     int32(req.Uid),
		Name:    req.Name,
		DepId:   int32(dep.Id),
		OrgId:   int32(dep.OrgId),
		DepName: dep.Name,
		OrgName: dep.OrgName,
	})
	if err != nil {
		return nil, errorx.HandleGrpcErrorToHttp(err)
	}

	return &types.AddInterRsp{
		Msg: "添加成功",
	}, nil
}
