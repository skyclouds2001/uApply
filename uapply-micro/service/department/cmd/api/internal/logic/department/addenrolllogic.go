package department

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"uapply-micro/common/errorx"
	"uapply-micro/common/filed"
	"uapply-micro/service/department/cmd/api/internal/logic"
	"uapply-micro/service/department/model"
	"uapply-micro/service/user/cmd/rpc/user"

	"uapply-micro/service/department/cmd/api/internal/svc"
	"uapply-micro/service/department/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddEnrollLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddEnrollLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddEnrollLogic {
	return AddEnrollLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddEnrollLogic) AddEnroll(req types.AddEnrollReq) (resp *types.Msg, err error) {
	// 先保存简历
	_, err = l.svcCtx.UserRpc.SaveResume(l.ctx, &user.UserInfo{
		Uid:        req.Uid,
		Name:       req.Name,
		Sex:        req.Sex,
		StudentNum: req.StudentNum,
		AddressNum: req.AddressNum,
		Major:      req.Major,
		PhoneNum:   req.PhoneNum,
		Email:      req.Email,
		Intro:      req.Intro,
	})
	if err != nil {
		fmt.Printf("%v", err)
		return nil, errorx.HandleGrpcErrorToHttp(err)
	}

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
	depInfo, err := l.svcCtx.DepModel.FindOne(depId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, nil
		}
		return nil, errorx.NewSysError()
	}
	// 再添加申请
	_, err = l.svcCtx.RegModel.Insert(&model.Register{
		Uid:         req.Uid,
		OrgId:       depInfo.OrgId,
		OrgName:     depInfo.OrgName,
		DepId:       depInfo.Id,
		DepName:     depInfo.Name,
		FinalStatus: filed.PASS,
	})
	if err != nil {
		if errMysql, ok := err.(*mysql.MySQLError); ok {
			switch errMysql.Number {
			case 1062:
				return nil, errorx.NewCodeError(errorx.Exist, "该用户已存在")
			default:
				return nil, errorx.NewSysError()
			}
		}
		return nil, errorx.NewSysError()
	}

	resp = &types.Msg{Msg: "添加成功"}
	return
}
