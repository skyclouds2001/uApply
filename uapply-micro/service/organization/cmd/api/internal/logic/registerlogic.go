package logic

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"strings"
	"uapply-micro/common/errorx"
	"uapply-micro/service/organization/model"

	"uapply-micro/service/organization/cmd/api/internal/svc"
	"uapply-micro/service/organization/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Register 组织注册
func (l *RegisterLogic) Register(req types.RegReq) (*types.CommonReply, error) {
	if len(strings.TrimSpace(req.Account)) == 0 || len(strings.TrimSpace(req.Password)) == 0 || len(strings.TrimSpace(req.Name)) == 0 {
		return nil, errorx.NewError("参数错误", errorx.ParamInvalid)
	}
	// 检查名字是否合法
	exist, err := l.svcCtx.OrgModel.FindOneByName(req.Name)
	if err != nil {
		if !errors.Is(err, model.ErrNotFound) {
			return nil, errorx.NewSysError()
		}
	} else {
		if exist != nil {
			return nil, errorx.NewError("组织名字非法，请更换", errorx.Exist)
		}
	}

	// 插入组织表
	org := &model.Organization{
		Name:     req.Name,
		Account:  req.Account,
		Password: req.Password,
	}
	res, err := l.svcCtx.OrgModel.Insert(org)
	if err != nil {
		logx.Error(err)
		if errMySQL, ok := err.(*mysql.MySQLError); ok {
			switch errMySQL.Number {
			case 1062:
				return nil, errorx.NewError("账号重复，请更换", errorx.Exist)
			default:
				return nil, errorx.NewSysError()
			}
		}
		return nil, errorx.NewSysError()
	}
	id, _ := res.LastInsertId()
	return &types.CommonReply{
		Id:   id,
		Name: org.Name,
	}, nil
}
