package org

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"uapply-micro/common/errorx"
	"uapply-micro/common/filed"
	"uapply-micro/service/department/cmd/rpc/department"
	"uapply-micro/service/department/model"
	"uapply-micro/service/organization/cmd/api/internal/logic"
	"uapply-micro/service/user/cmd/rpc/user"

	"uapply-micro/service/organization/cmd/api/internal/svc"
	"uapply-micro/service/organization/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetDataLogic {
	return GetDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDataLogic) GetData() (*types.GetOrgDataRsp, error) {
	// 拿到组织的各部门数据
	// 先获取 orgId
	orgIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value(logic.KeyID)))
	orgId, err := orgIdNumber.Int64()
	if err != nil {
		logx.Error(err)
		return nil, errorx.NewSysError()
	}
	if orgId <= 0 {
		return nil, errorx.NewError("登录信息错误", errorx.Fail)
	}

	// 获取组织名称
	orgInfo, err := l.svcCtx.OrgModel.FindOne(orgId)
	if err != nil {
		logx.Error(err)
		if errors.Is(err, model.ErrNotFound) {
			return nil, errorx.NewCodeError(errorx.NotFound, "未找到组织信息")
		}
		return nil, errorx.NewSysError()
	}

	// 拿到组织下的部门id
	deps, err := l.svcCtx.DepRpc.GetDeps(l.ctx, &department.DepReq{OrgId: int32(orgId)})
	if err != nil {
		errorx.HandleGrpcErrorToHttp(err)
		return nil, err
	}
	var depids []int32
	var depDatas []types.Data
	for _, dep := range deps.Deps {
		depids = append(depids, dep.Id)
	}
	ids, _ := l.svcCtx.DepRpc.GetRegByDepIds(l.ctx, &department.RbdReq{Ids: depids})
	logx.Info(ids.Data, len(ids.Data), len(depids))
	if len(ids.Data) != len(depids) {
		return nil, errorx.NewError("系统数据出错，请联系后台人员", errorx.Fail)
	}
	for idx, datum := range ids.Data {
		// 通过uid查性别
		var uids []int64
		for _, data := range datum.Dep {
			uids = append(uids, int64(data.Uid))
		}
		var rsp types.Data
		rsp.Sum = len(datum.Dep)
		logx.Info(depids, datum.Dep, idx)
		rsp.ID = int(depids[idx])
		rsp.Name = deps.Deps[idx].Name
		if len(uids) != 0 {
			sex, err := l.svcCtx.UserRpc.GetUserSex(l.ctx, &user.GetSexReq{Ids: uids})
			logx.Info("sex:%v", sex)
			if len(sex.Sexs) != len(datum.Dep) {
				return nil, errorx.NewError("系统数据出错，请联系后台人员", errorx.Fail)
			}
			if err != nil {
				return nil, errorx.HandleGrpcErrorToHttp(err)
			}
			for i, user := range datum.Dep {
				// 男生
				if sex.Sexs[i] == filed.MALE && user.Final == filed.PASS {
					rsp.Male++
				}
				// 女生
				if sex.Sexs[i] == filed.FEMALE && user.Final == filed.PASS {
					rsp.Female++
				}
			}
			rsp.Pass = rsp.Male + rsp.Female
		} else {
			rsp.Sum = 0
		}
		depDatas = append(depDatas, rsp)
	}
	var orgdata types.Data
	for _, data := range depDatas {
		orgdata.Male += data.Male
		orgdata.Female += data.Female
		orgdata.Pass += data.Pass
		orgdata.Sum += data.Sum
	}
	orgdata.ID = int(orgId)
	orgdata.Name = orgInfo.Name
	return &types.GetOrgDataRsp{
		Org: orgdata,
		Dep: depDatas,
	}, nil
}
