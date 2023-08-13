package logic

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"uapply-micro/service/department/model"

	"uapply-micro/service/department/cmd/rpc/department"
	"uapply-micro/service/department/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateDepLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDepLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDepLogic {
	return &CreateDepLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateDepLogic) CreateDep(in *department.CreReq) (*department.CreInfoRsp, error) {
	if len(strings.TrimSpace(in.Account)) == 0 || len(strings.TrimSpace(in.Password)) == 0 || len(strings.TrimSpace(in.Name)) == 0 {
		return nil, status.Error(codes.InvalidArgument, "参数错误")
	}
	// 检查名字是否合法
	exist, err := l.svcCtx.DepModel.FindOneByName(in.Name)
	if err != nil {
		if !errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.Internal, "系统繁忙")
		}
	} else {
		if exist != nil {
			return nil, status.Error(codes.AlreadyExists, "部门名字非法，请更换")
		}
	}
	// 插入数据
	_, err = l.svcCtx.DepModel.Insert(&model.Department{
		Name:     in.Name,
		Account:  in.Account,
		Password: in.Password,
		OrgId:    int64(in.OrgId),
		OrgName:  in.OrgName,
	})
	if err != nil {
		if errMySQL, ok := err.(*mysql.MySQLError); ok {
			switch errMySQL.Number {
			case 1062:
				return nil, status.Error(codes.AlreadyExists, "账号重复，请更换")
			default:
				return nil, status.Error(codes.Internal, "系统繁忙")
			}
		}
		return nil, status.Error(codes.Internal, "系统繁忙")
	}
	return &department.CreInfoRsp{
		Name:  in.Name,
		OrgId: in.OrgId,
	}, nil
}
