package inter

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"uapply-micro/common/errorx"
	"uapply-micro/service/interviewer/cmd/api/internal/logic"

	"uapply-micro/service/interviewer/cmd/api/internal/svc"
	"uapply-micro/service/interviewer/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetOrgDepLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrgDepLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetOrgDepLogic {
	return GetOrgDepLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrgDepLogic) GetOrgDep() (resp *types.InterDepOrg, err error) {
	uidnum := json.Number(fmt.Sprintf("%v", l.ctx.Value(logic.UIDKey)))
	uid, err := uidnum.Int64()
	if err != nil {
		log.Println(err)
		return nil, errorx.NewSysError()
	}

	interMsg, err := l.svcCtx.InterInfo.FindByUid(uid)
	if err != nil {
		return nil, errorx.NewSysError()
	}
	if interMsg == nil {
		return nil, errorx.NewError("没有面试官权限", errorx.NotAuth)
	}

	resp = new(types.InterDepOrg)
	resp.Dep = make(map[int]string, len(interMsg))
	resp.Org = make(map[int]string, len(interMsg))

	for _, info := range interMsg {
		resp.Org[int(info.OrganizationId)] = info.OrganizationName
		resp.Dep[int(info.DepartmentId)] = info.DepartmentName
	}

	return
}
