package department

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"uapply-micro/common/errorx"
	"uapply-micro/common/filed"
	"uapply-micro/service/department/cmd/api/internal/logic"
	"uapply-micro/service/department/model"
	"uapply-micro/service/user/cmd/rpc/userclient"

	"uapply-micro/service/department/cmd/api/internal/svc"
	"uapply-micro/service/department/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetOutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetOutLogic {
	return GetOutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOutLogic) GetOut() (*types.UsersInfo, error) {
	/// 拿部门id
	depIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value(logic.DepIdKey)))
	depId, err := depIdNumber.Int64()
	if err != nil {
		logx.Error(err)
		return nil, errorx.NewSysError()
	}
	if depId <= 0 {
		return nil, errorx.NewError("登录信息错误", errorx.Fail)
	}

	users, err := l.svcCtx.RegModel.FindFinalStatus(filed.OUT, depId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, nil
		}
		return nil, errorx.NewSysError()
	}

	var uid []int64
	for _, u := range users {
		uid = append(uid, u.Uid)
	}
	// 如果没有数据就没必要rpc调用了
	if len(uid) == 0 {
		return &types.UsersInfo{Users: make([]types.UserInfo, 0)}, nil
	}
	sum := len(uid)

	usersRsp, err := l.svcCtx.UserRpc.GetUsers(l.ctx, &userclient.GetSexReq{
		Ids: uid,
	})
	if err != nil {
		return nil, errorx.HandleGrpcErrorToHttp(err)
	}

	var rsp []types.UserInfo
	for i, u := range usersRsp.Users {
		if users[i].Uid == u.Uid {
			m := types.UserInfo{
				Uid:        int(u.Uid),
				Name:       u.Name,
				Sex:        u.Sex,
				Phone:      u.PhoneNum,
				Email:      u.Email,
				StudentNum: u.StudentNum,
				Intro:      u.Intro,
				Status:     filed.OUT,
			}
			out := users[i]
			if out.SecondStatus == filed.OUT {
				m.Eva = out.SecondEva.String
				m.Mark = int(out.SecondMark)
			}
			if out.FinalStatus == filed.OUT {
				m.Eva = out.FirstEva.String
				m.Mark = int(out.FirstMark)
			}
			rsp = append(rsp, m)
		}
	}
	return &types.UsersInfo{
		Sum:   sum,
		Users: rsp,
	}, nil
}
