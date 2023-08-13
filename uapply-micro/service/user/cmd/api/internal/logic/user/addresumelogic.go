package user

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"regexp"
	"uapply-micro/common/errorx"
	"uapply-micro/service/user/cmd/api/internal/logic"
	"uapply-micro/service/user/model"

	"uapply-micro/service/user/cmd/api/internal/svc"
	"uapply-micro/service/user/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddResumeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddResumeLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddResumeLogic {
	return AddResumeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddResumeLogic) AddResume(req types.AddResumeReq) (*types.AddResumeRsp, error) {
	// 拿到 uid
	uidNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value(logic.UIDKey)))
	uid, err := uidNumber.Int64()
	if err != nil {
		logx.Error(err)
		return nil, errorx.NewSysError()
	}
	if uid <= 0 {
		return nil, errorx.NewError("登录信息错误", errorx.Fail)
	}
	// 验证参数
	if req.Name == "" || req.Intro == "" || req.StudentNum == "" {
		return nil, errorx.NewCodeError(errorx.ParamInvalid, "有必填参数为空")
	}
	// 验证手机号和邮箱
	if !validateMobile(req.PhoneNum) {
		return nil, errorx.NewCodeError(errorx.ParamInvalid, "手机号格式错误")
	}
	if !validateEmail(req.Email) {
		return nil, errorx.NewCodeError(errorx.ParamInvalid, "邮箱格式错误")
	}
	// 插入
	_, err = l.svcCtx.ResumeModel.Insert(&model.UserInfo{
		Uid:        uid,
		Name:       req.Name,
		Sex:        req.Sex,
		StudentNum: req.StudentNum,
		AddressNum: sql.NullString{
			String: req.AddressNum,
			Valid:  true,
		},
		Major: sql.NullString{
			String: req.Major,
			Valid:  true,
		},
		PhoneNum: req.PhoneNum,
		Email:    req.Email,
		Intro:    req.Intro,
	})
	if err != nil {
		logx.Error(err)
		if errMySQL, ok := err.(*mysql.MySQLError); ok {
			switch errMySQL.Number {
			case 1062:
				err = l.svcCtx.ResumeModel.Update(&model.UserInfo{
					Uid:        uid,
					Name:       req.Name,
					Sex:        req.Sex,
					StudentNum: req.StudentNum,
					AddressNum: sql.NullString{
						String: req.AddressNum,
						Valid:  true,
					},
					Major: sql.NullString{
						String: req.Major,
						Valid:  true,
					},
					PhoneNum: req.PhoneNum,
					Email:    req.Email,
					Intro:    req.Intro,
				})
				if err != nil {
					return nil, errorx.NewSysError()
				}
				return &types.AddResumeRsp{
					Msg: "修改简历成功",
				}, nil
			default:
				return nil, errorx.NewSysError()
			}
		}
		return nil, errorx.NewSysError()
	}
	return &types.AddResumeRsp{
		Msg: "保存简历成功",
	}, nil
}

const (
	EMAIL = `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	PHONE = `^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`
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
