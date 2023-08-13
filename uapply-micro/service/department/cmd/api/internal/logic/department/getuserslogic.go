package department

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"uapply-micro/common/errorx"
	"uapply-micro/service/department/cmd/api/internal/logic"
	"uapply-micro/service/department/model"
	"uapply-micro/service/user/cmd/rpc/userclient"

	"uapply-micro/service/department/cmd/api/internal/svc"
	"uapply-micro/service/department/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUsersLogic {
	return GetUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetUsers 获取一面所以人的状态
func (l *GetUsersLogic) GetUsers() (*types.UsersInfo, error) {
	// 拿部门id
	depIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value(logic.DepIdKey)))
	depId, err := depIdNumber.Int64()
	log.Println("-------------------")
	if err != nil {
		logx.Error(err)
		return nil, errorx.NewSysError()
	}
	if depId <= 0 {
		return nil, errorx.NewError("登录信息错误", errorx.Fail)
	}
	// 查出uid和对应的状态
	all, err := l.svcCtx.RegModel.FindDepAll(depId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, nil
		}
		return nil, errorx.NewSysError()
	}
	var ids []int64
	for _, user := range all {
		ids = append(ids, user.Uid)
	}
	sum := len(ids)
	if sum == 0 {
		return &types.UsersInfo{Users: make([]types.UserInfo, 0)}, nil
	}

	users, err := l.svcCtx.UserRpc.GetUsers(l.ctx, &userclient.GetSexReq{
		Ids: ids,
	})
	if err != nil {
		return nil, errorx.HandleGrpcErrorToHttp(err)
	}

	if sum != len(users.Users) {
		return nil, errorx.NewError("有非法数据，请联系后台人员", errorx.Fail)
	}

	var rsp []types.UserInfo
	for i, user := range users.Users {
		if all[i].Uid == user.Uid {
			rsp = append(rsp, types.UserInfo{
				Uid:        int(all[i].Uid),
				Name:       user.Name,
				Sex:        user.Sex,
				Phone:      user.PhoneNum,
				Email:      user.Email,
				StudentNum: user.StudentNum,
				Intro:      user.Intro,
				Status:     int(all[i].FirstStatus),
				Eva:        all[i].FirstEva.String,
				Mark:       int(all[i].FirstMark),
			})
		}
	}
	return &types.UsersInfo{
		Sum:   sum,
		Users: rsp,
	}, nil
}
