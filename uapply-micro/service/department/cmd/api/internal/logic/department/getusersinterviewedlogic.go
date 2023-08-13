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

type GetUsersInterviewedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUsersInterviewedLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUsersInterviewedLogic {
	return GetUsersInterviewedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetUsersInterviewed 获取第一面已经面试的同学的状态
// 返回第二面状态和最终状态
func (l *GetUsersInterviewedLogic) GetUsersInterviewed() (*types.UsersInfo2, error) {
	// 拿部门id
	depIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value(logic.DepIdKey)))
	depId, err := depIdNumber.Int64()
	log.Println("------------------------")
	if err != nil {
		logx.Error(err)
		return nil, errorx.NewSysError()
	}
	if depId <= 0 {
		return nil, errorx.NewError("登录信息错误", errorx.Fail)
	}

	// 查库
	users, err := l.svcCtx.RegModel.FindUsersInterviewedFirstTurn(depId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, nil
		}
		return nil, errorx.NewSysError()
	}
	// 拿出uid
	var ids []int64
	for _, user := range users {
		ids = append(ids, user.Uid)
	}

	sum := len(ids)
	if sum == 0 {
		return &types.UsersInfo2{Users: make([]types.UserInfo2, 0)}, nil
	}

	// 查人信息
	all, err := l.svcCtx.UserRpc.GetUsers(l.ctx, &userclient.GetSexReq{
		Ids: ids,
	})
	if err != nil {
		return nil, errorx.HandleGrpcErrorToHttp(err)
	}

	if sum != len(all.Users) {
		return nil, errorx.NewError("有非法数据，请联系后台人员", errorx.Fail)
	}
	// 封装返回格式
	var rsp []types.UserInfo2
	for i, user := range all.Users {
		if users[i].Uid == user.Uid {
			m := types.UserInfo2{
				Uid:          int(users[i].Uid),
				Name:         user.Name,
				Sex:          user.Sex,
				Phone:        user.PhoneNum,
				Email:        user.Email,
				StudentNum:   user.StudentNum,
				Intro:        user.Intro,
				SecondStatus: int(users[i].SecondStatus),
			}
			u := users[i]
			if u.SecondStatus != 0 {
				m.Eva = u.SecondEva.String
				m.Mark = int(u.SecondMark)
			} else {
				m.Eva = u.FirstEva.String
				m.Mark = int(u.FirstMark)
			}
			rsp = append(rsp, m)
		}
	}
	return &types.UsersInfo2{
		Sum:   sum,
		Users: rsp,
	}, nil
}
