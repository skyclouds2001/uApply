package logic

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
	"uapply-micro/common/errorx"
	"uapply-micro/service/department/model"

	"uapply-micro/service/department/cmd/api/internal/svc"
	"uapply-micro/service/department/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

var (
	OrgIdKey   = "orgid"
	OrgNameKey = "orgname"
	DepIdKey   = "depid"
	DepNameKey = "depname"
)

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

func (l *LoginLogic) Login(req types.LoginReq) (*types.CommonReply, error) {
	dep, err := l.svcCtx.DepModel.FindOneByAccount(req.Account)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errorx.NewError("账号不存在", errorx.NotFound)
		}
		return nil, errorx.NewSysError()
	}
	// 检查密码
	if dep.Password != req.Password {
		return nil, errorx.NewError("密码错误", errorx.PasswordInvalid)
	}
	// 获取token
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, dep)
	if err != nil {
		return nil, errorx.NewSysError()
	}
	return &types.CommonReply{
		Id:           int(dep.Id),
		Name:         dep.Name,
		OrgId:        int(dep.OrgId),
		OrgName:      dep.OrgName,
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: accessExpire/2 + now,
	}, nil
}

// 登录获取token
func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds int64, req *model.Department) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[OrgIdKey] = req.OrgId
	claims[OrgNameKey] = req.OrgName
	claims[DepIdKey] = req.Id
	claims[DepNameKey] = req.Name
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
