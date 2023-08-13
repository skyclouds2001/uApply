package logic

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"uapply-micro/service/user/model"

	"uapply-micro/service/user/cmd/rpc/internal/svc"
	"uapply-micro/service/user/cmd/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersLogic {
	return &GetUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUsers 获取用户信息
func (l *GetUsersLogic) GetUsers(in *user.GetSexReq) (*user.GetUsersRsp, error) {
	var ids []int
	for _, id := range in.Ids {
		ids = append(ids, int(id))
	}
	all, err := l.svcCtx.UserModel.FindAll(ids, l.svcCtx.Mysqlx)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "没有uid对应信息")
		}
		return nil, status.Error(codes.Internal, "系统繁忙")
	}
	var users []*user.UserInfo
	for _, info := range all {
		users = append(users, &user.UserInfo{
			Uid:        info.Uid,
			Name:       info.Name,
			Sex:        info.Sex,
			StudentNum: info.StudentNum,
			AddressNum: info.AddressNum.String,
			Major:      info.Major.String,
			PhoneNum:   info.PhoneNum,
			Email:      info.Email,
			Intro:      info.Intro,
		})
	}
	return &user.GetUsersRsp{
		Users: users,
	}, nil
}
