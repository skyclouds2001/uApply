package logic

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
	"uapply-micro/service/user/model"

	"uapply-micro/service/user/cmd/rpc/internal/svc"
	"uapply-micro/service/user/cmd/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type SaveResumeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveResumeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveResumeLogic {
	return &SaveResumeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 保存简历，该接口会做参数校验
func (l *SaveResumeLogic) SaveResume(in *user.UserInfo) (*user.Empty, error) {
	// 查看登录信息
	_, err := l.svcCtx.LoginModel.FindOne(in.Uid)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "用户不存在")
		}
		return nil, status.Error(codes.Internal, "系统繁忙")
	}

	// 检查参数
	if !checkInParam(in) {
		return nil, status.Error(codes.InvalidArgument, "参数错误")
	}

	_, err = l.svcCtx.UserModel.FindOne(in.Uid)
	if err != nil {
		// 未找到，插入
		if errors.Is(err, model.ErrNotFound) {
			_, err := l.svcCtx.UserModel.Insert(&model.UserInfo{
				Uid:        in.Uid,
				Name:       in.Name,
				Sex:        in.Sex,
				StudentNum: in.StudentNum,
				AddressNum: sql.NullString{String: in.AddressNum, Valid: true},
				Major:      sql.NullString{String: in.Major, Valid: true},
				PhoneNum:   in.PhoneNum,
				Email:      in.Email,
				Intro:      in.Intro,
			})
			// 添加失败
			if err != nil {
				if errMysql, ok := err.(*mysql.MySQLError); ok {
					switch errMysql.Number {
					case 1062:
						return nil, status.Error(codes.AlreadyExists, "该用户已存在")
					default:
						return nil, status.Error(codes.Internal, "系统繁忙")
					}
				}
				return nil, status.Error(codes.Internal, "系统繁忙")
			}
			return &user.Empty{}, nil
		}
		return nil, status.Error(codes.Internal, "系统繁忙")
	}

	// 没出错，则更新
	err = l.svcCtx.UserModel.Update(&model.UserInfo{
		Uid:        in.Uid,
		Name:       in.Name,
		Sex:        in.Sex,
		StudentNum: in.StudentNum,
		AddressNum: sql.NullString{String: in.AddressNum, Valid: true},
		Major:      sql.NullString{String: in.Major, Valid: true},
		PhoneNum:   in.PhoneNum,
		Email:      in.Email,
		Intro:      in.Intro,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "系统繁忙")
	}

	return &user.Empty{}, nil
}

func checkInParam(in *user.UserInfo) bool {
	if in.Uid <= 0 || in.Name == "" || (in.Sex != "男" && in.Sex != "女") || in.StudentNum == "" || in.Intro == "" {
		return false
	}
	return validateMobile(in.PhoneNum) && validateEmail(in.Email)
}

const (
	EMAIL = `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	PHONE = `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
)

// ValidateMobile 自定义手机验证
func validateMobile(mobile string) bool {
	// 手机号码正则表达式验证
	return validateFunc(mobile, PHONE)
}

func validateEmail(email string) bool {
	// 邮箱正则验证
	return validateFunc(email, EMAIL)
}

func validateFunc(val string, rule string) bool {
	ok, _ := regexp.MatchString(rule, val)
	if ok {
		return true
	}
	return false
}
