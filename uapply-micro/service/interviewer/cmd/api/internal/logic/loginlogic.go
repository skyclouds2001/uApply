package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"uapply-micro/common/errorx"
	"uapply-micro/service/interviewer/cmd/api/internal/svc"
	"uapply-micro/service/interviewer/cmd/api/internal/types"
	"uapply-micro/service/interviewer/model"

	"github.com/tal-tech/go-zero/core/logx"
)

const (
	UIDKey     = "uid"
	OpenIDKey  = "openid"
	WXLoginUrl = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
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

func (l *LoginLogic) Login(req types.LoginReq) (resp *types.LoginRsp, err error) {
	url := fmt.Sprintf(WXLoginUrl,
		l.svcCtx.Config.App.Appid,
		l.svcCtx.Config.App.Secret,
		req.Code,
	)
	client := &http.Client{}

	request, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return nil, errorx.NewSysError()
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, errorx.NewSysError()
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errorx.NewSysError()
	}

	var ws types.WxLoginForm
	if err := json.Unmarshal(body, &ws); err != nil {
		return nil, errorx.NewSysError()
	}

	// 如果code失效或者错误
	if ws.OpenID == "" {
		return nil, errorx.NewCodeError(errorx.ParamInvalid, "登录code不合法")
	}

	if err != nil {
		return nil, errorx.NewSysError()
	}

	var uid int64
	id, err := l.svcCtx.LoginInfo.FindOneByOpenId(ws.OpenID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			// 没有记录就要插入
			insert, err := l.svcCtx.LoginInfo.Insert(&model.LoginInfo{OpenId: ws.OpenID})
			if err != nil {
				return nil, errorx.NewSysError()
			}
			// 拿到uid
			uid, err = insert.LastInsertId()
			if err != nil {
				return nil, errorx.NewSysError()
			}
		} else {
			return nil, errorx.NewSysError()
		}
	} else {
		uid = id.Uid
	}

	log.Println(uid)
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	token, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, ws.OpenID, int(uid))
	if err != nil {
		return nil, errorx.NewSysError()
	}

	resp = new(types.LoginRsp)
	resp.AccessToken = token
	resp.AccessExpire = now + accessExpire
	resp.RefreshAfter = now + accessExpire/2
	resp.UID = int(uid)

	return
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds int64, openId string, uid int) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[UIDKey] = uid
	claims[OpenIDKey] = openId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
