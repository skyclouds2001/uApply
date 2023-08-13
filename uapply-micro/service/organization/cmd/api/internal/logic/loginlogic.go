package logic

import (
	"context"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
	"uapply-micro/common/errorx"
	"uapply-micro/service/organization/model"

	"uapply-micro/service/organization/cmd/api/internal/svc"
	"uapply-micro/service/organization/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

var KeyID = "orgid"
var KeyName = "orgname"

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Login 组织登录
func (l *LoginLogic) Login(req types.LoginReq) (*types.CommonReply, error) {
	if len(strings.TrimSpace(req.Account)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errorx.NewError("参数错误", errorx.ParamInvalid)
	}
	orgInfo, err := l.svcCtx.OrgModel.FindOneByAccount(req.Account)
	logx.Error(err)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.NewError("账号不存在", errorx.NotFound)
	default:
		logx.Error(err)
		return nil, errorx.NewSysError()
	}

	if orgInfo.Password != req.Password {
		return nil, errorx.NewError("密码错误", errorx.PasswordInvalid)
	}

	// jwt鉴权
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, orgInfo.Id, orgInfo.Name)
	if err != nil {
		return nil, errorx.NewSysError()
	}

	return &types.CommonReply{
		Id:           orgInfo.Id,
		Name:         orgInfo.Name,
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

// 登录获取token
func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, orgId int64, orgName string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["orgid"] = orgId
	claims["orgname"] = orgName
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
