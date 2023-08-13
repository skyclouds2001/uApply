package logic

import (
	"context"
	"errors"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"uapply-micro/service/user/cmd/rpc/internal/svc"
	"uapply-micro/service/user/cmd/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserByUidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByUidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByUidLogic {
	return &GetUserByUidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  根据uid获取用户信息
func (l *GetUserByUidLogic) GetUserByUid(in *user.UserUid) (*user.UserInfo, error) {
	userMsg, err := l.svcCtx.UserModel.FindOne(in.Uid)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "用户不存在")
		}
		return nil, status.Error(codes.Internal, "系统错误")
	}
	rsq := &user.UserInfo{
		Uid:        in.Uid,
		Name:       userMsg.Name,
		Sex:        userMsg.Sex,
		StudentNum: userMsg.StudentNum,
		AddressNum: userMsg.AddressNum.String,
		Major:      userMsg.Major.String,
		PhoneNum:   userMsg.PhoneNum,
		Email:      userMsg.Email,
		Intro:      userMsg.Intro,
	}
	return rsq, nil
}
