package logic

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"uapply-micro/service/department/model"

	"uapply-micro/service/department/cmd/rpc/department"
	"uapply-micro/service/department/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *department.RegReq) (*department.Empty, error) {
	// 检查部门是否存在
	dep, err := l.svcCtx.DepModel.FindOne(int64(in.DepId))
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.InvalidArgument, "无此部门")
		}
		return nil, status.Error(codes.Internal, "系统繁忙")
	}
	if int32(dep.OrgId) != in.OrgId {
		return nil, status.Error(codes.InvalidArgument, "该组织下无此部门")
	}

	// 需求是一个组织最多两个申报的志愿
	regCnt, err := l.svcCtx.RegModel.CountByOrgAndUid(int64(in.OrgId), int64(in.Uid), l.svcCtx.Mysqlx)
	if err != nil {
		return nil, status.Error(codes.Internal, "系统繁忙")
	}
	if regCnt >= 2 {
		return nil, status.Error(codes.AlreadyExists, "一个组织最多两个志愿")
	}

	// 插入
	_, err = l.svcCtx.RegModel.Insert(&model.Register{
		Uid:     int64(in.Uid),
		OrgId:   int64(in.OrgId),
		OrgName: dep.OrgName,
		DepId:   int64(in.DepId),
		DepName: dep.Name,
	})
	if err != nil {
		if errMysql, ok := err.(*mysql.MySQLError); ok {
			switch errMysql.Number {
			case 1062:
				return nil, status.Error(codes.AlreadyExists, "已申报过该部门")
			default:
				return nil, status.Error(codes.Internal, "系统繁忙")
			}
		}
		return nil, status.Error(codes.Internal, "系统繁忙")
	}
	return &department.Empty{}, nil
}
