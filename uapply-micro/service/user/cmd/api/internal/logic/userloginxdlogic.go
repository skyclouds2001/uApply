package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt"
	"io"
	"net/http"
	"net/url"
	"time"
	"uapply-micro/common/errorx"
	"uapply-micro/service/user/cmd/api/internal/svc"
	"uapply-micro/service/user/cmd/api/internal/types"
	"uapply-micro/service/user/model"

	"github.com/tal-tech/go-zero/core/logx"
)

const (
	XdStuNumKey     = "xd-stu-num"
	XdLoginUrl      = "https://xxcapp.xidian.edu.cn/uc/wap/login/check"
	XdLoginSuccessE = 0
)

type UserLoginXdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginXdLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserLoginXdLogic {
	return UserLoginXdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginXdLogic) UserLoginXd(req types.XdLoginReq) (resp *types.LoginRsp, err error) {
	form := url.Values{}
	form.Set("username", req.Username)
	form.Set("password", req.Password)

	loginResp, err := http.PostForm(XdLoginUrl, form)
	if err != nil {
		return nil, errorx.NewSysError()
	}
	body, err := io.ReadAll(loginResp.Body)
	if err != nil {
		return nil, errorx.NewSysError()
	}

	var loginResult types.XdLoginForm
	if err := json.Unmarshal(body, &loginResult); err != nil {
		return nil, errorx.NewSysError()
	}
	if loginResult.E != XdLoginSuccessE {
		return nil, errorx.NewCodeError(errorx.ParamInvalid, "登录失败，请检查账号密码")
	}

	var uid int64
	loginInfo, err := l.svcCtx.LoginModel.FindOneByXdStuNum(sql.NullString{String: req.Username, Valid: true})
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			// 插入
			result, err := l.svcCtx.LoginModel.Insert(&model.LoginInfo{XdStuNum: sql.NullString{String: req.Username, Valid: true}})
			if err != nil {
				return nil, errorx.NewSysError()
			}
			// 取出 uid
			uid, err = result.LastInsertId()
			if err != nil {
				return nil, errorx.NewSysError()
			}
		} else {
			return nil, errorx.NewSysError()
		}
	} else {
		uid = loginInfo.Uid
	}

	// 生成 token
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	token, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, req.Username, int(uid))
	if err != nil {
		return nil, errorx.NewSysError()
	}

	resp = &types.LoginRsp{
		AccessToken:  token,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
		UID:          int(uid),
	}
	return
}

func (l *UserLoginXdLogic) getJwtToken(secretKey string, iat, seconds int64, xdStuNum string, uid int) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[UIDKey] = uid
	claims[XdStuNumKey] = xdStuNum
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
