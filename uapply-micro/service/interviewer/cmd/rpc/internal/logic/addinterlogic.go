package logic

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"uapply-micro/service/interviewer/cmd/rpc/inter"
	"uapply-micro/service/interviewer/model"

	"uapply-micro/service/interviewer/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddInterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddInterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddInterLogic {
	return &AddInterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddInter 添加面试官接口
func (l *AddInterLogic) AddInter(in *inter.AddReq) (*inter.Empty, error) {
	// 检查uid是否存在
	_, err := l.svcCtx.LoginModel.FindOne(int64(in.Uid))
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "用户不存在")
		}
	}
	// 添加
	_, err = l.svcCtx.InterModel.Insert(&model.InterInfo{
		Uid:              int64(in.Uid),
		Name:             in.Name,
		OrganizationId:   int64(in.OrgId),
		OrganizationName: in.OrgName,
		DepartmentId:     int64(in.DepId),
		DepartmentName:   in.DepName,
	})
	if err != nil {
		if errMysql, ok := err.(*mysql.MySQLError); ok {
			switch errMysql.Number {
			case 1062:

				return nil, status.Error(codes.AlreadyExists, "面试官已存在")
			default:
				return nil, status.Error(codes.Internal, "系统繁忙")
			}
		}
		return nil, status.Error(codes.Internal, "系统繁忙")
	}
	return &inter.Empty{}, nil
}
