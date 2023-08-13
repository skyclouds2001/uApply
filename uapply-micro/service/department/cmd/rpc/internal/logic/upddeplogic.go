package logic

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"uapply-micro/common/password"
	"uapply-micro/service/department/model"

	"uapply-micro/service/department/cmd/rpc/department"
	"uapply-micro/service/department/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdDepLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdDepLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdDepLogic {
	return &UpdDepLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdDepLogic) UpdDep(in *department.UdpReq) (*department.DepInfoRsp, error) {
	if in.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "id必须大于0")
	}
	dep, err := l.svcCtx.DepModel.FindOne(int64(in.Id))
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.InvalidArgument, "无此部门")
		}
		return nil, status.Error(codes.Internal, "系统繁忙")
	}
	if dep.OrgId != int64(in.OrgId) {
		return nil, status.Error(codes.InvalidArgument, "组织下无此部门")
	}
	if in.Name != "" {
		// 检查名字是否合法
		exist, err := l.svcCtx.DepModel.FindOneByName(in.Name)
		if err != nil {
			if !errors.Is(err, model.ErrNotFound) {
				return nil, status.Error(codes.Internal, "系统繁忙")
			}
		} else {
			if exist != nil {
				return nil, status.Error(codes.AlreadyExists, "名字非法，请更换")
			}
		}
		dep.Name = in.Name
	}
	if in.Account != "" {
		dep.Account = in.Account
	}
	if in.Password != "" {
		ok := password.VerifyPwd(in.Password, 0)
		if !ok {
			return nil, status.Error(codes.InvalidArgument, "密码格式错误")
		}
		dep.Password = in.Password
	}
	err = l.svcCtx.DepModel.Update(dep)
	if err != nil {
		if errMySQL, ok := err.(*mysql.MySQLError); ok {
			switch errMySQL.Number {
			case 1062:
				return nil, status.Error(codes.AlreadyExists, "账号重复，请更换")
			default:
				return nil, status.Error(codes.Internal, "系统繁忙")
			}
		}
		return nil, status.Error(codes.Internal, "更新失败")
	}
	return &department.DepInfoRsp{
		Id:       int32(dep.Id),
		Name:     dep.Name,
		Account:  dep.Account,
		Password: dep.Password,
	}, nil
}
