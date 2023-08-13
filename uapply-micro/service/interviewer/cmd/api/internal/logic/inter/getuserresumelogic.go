package inter

import (
	"context"
	"log"
	"uapply-micro/common/errorx"
	"uapply-micro/common/filed"
	"uapply-micro/service/department/cmd/rpc/department"
	"uapply-micro/service/interviewer/cmd/api/internal/svc"
	"uapply-micro/service/interviewer/cmd/api/internal/types"
	"uapply-micro/service/user/cmd/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserResumeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserResumeLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserResumeLogic {
	return GetUserResumeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserResumeLogic) GetUserResume(req types.GetReq) (resp *types.CommonRsq, err error) {
	num := req.Num
	if num > 2 || num < 1 {
		return nil, errorx.NewError("轮次错误", errorx.ParamInvalid)
	}

	uid := req.UID
	userdep := req.DepId
	log.Println(uid)
	in := &department.RegReq{Uid: int32(uid)}

	userMsg, err := l.svcCtx.DepRpc.GetRegister(l.ctx, in)
	if err != nil {
		log.Println(err)
		return nil, errorx.HandleGrpcErrorToHttp(err)
	}
	if userMsg == nil {
		return nil, errorx.NewError("用户不存在", errorx.NotFound)
	}

	// 设置标志位 如果为0 就是没有评价过的用户 如果是1 就是已经评价过的
	var tag int = 0
	// 可能他报名了多个组织 就会有多个登录信息
	for _, info := range userMsg.Rsp {
		if userdep != int(info.DepID) {
			continue
		}
		if info.Final == filed.OUT {
			return nil, errorx.NewError("用户已经被淘汰了", errorx.ParamInvalid)
		}
		switch num {
		case 1:
			if info.FirstMark != 0 || info.SecondMark != 0 || info.Final == filed.PASS {
				tag = 1
			}
		case 2:
			if info.First != filed.PASS {
				return nil, errorx.NewError("用户没有参加第一轮面试", errorx.ParamInvalid)
			}
			if info.SecondMark != 0 || info.Final == filed.PASS {
				tag = 1
			}
		}
		u := &user.UserUid{Uid: int64(uid)}
		userInfo, err := l.svcCtx.UserRpc.GetUserByUid(l.ctx, u)
		if err != nil {
			return nil, errorx.HandleGrpcErrorToHttp(err)
		}

		resp = new(types.CommonRsq)
		resp.PhoneNum = userInfo.PhoneNum
		resp.Email = userInfo.Email
		resp.Intro = userInfo.Intro
		resp.Major = userInfo.Major
		resp.Sex = userInfo.Sex
		resp.Name = userInfo.Name
		resp.StudentNum = userInfo.StudentNum
		resp.AddressNum = userInfo.AddressNum
		resp.Tag = tag
		return resp, nil
	}

	return nil, errorx.NewError("面试官没有权限", errorx.NotAuth)
}
