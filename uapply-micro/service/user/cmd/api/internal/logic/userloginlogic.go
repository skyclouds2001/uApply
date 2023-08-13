package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"io/ioutil"
	"net/http"
	"time"
	"uapply-micro/common/errorx"
	"uapply-micro/service/user/model"

	"uapply-micro/service/user/cmd/api/internal/svc"
	"uapply-micro/service/user/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

const (
	UIDKey     = "uid"
	OpenIDKey  = "openid"
	WXLoginUrl = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserLoginLogic {
	return UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req types.LoginReq) (*types.LoginRsp, error) {
	// 合成url, 这里的appId和secret是在微信公众平台上获取的
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

	// 先查找是否有uid
	var uid int64
	id, err := l.svcCtx.LoginModel.FindOneByOpenId(sql.NullString{String: ws.OpenID, Valid: true})
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			// 没有记录就要插入
			insert, err := l.svcCtx.LoginModel.Insert(&model.LoginInfo{OpenId: sql.NullString{String: ws.OpenID, Valid: true}})
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

	// 生成 token
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	token, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, ws.OpenID, int(uid))
	if err != nil {
		return nil, errorx.NewSysError()
	}

	return &types.LoginRsp{
		UID:          int(uid),
		AccessToken:  token,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2, // 建议客户端刷新token的绝对时间
	}, nil
}

// 登录获取token
func (l *UserLoginLogic) getJwtToken(secretKey string, iat, seconds int64, openId string, uid int) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[UIDKey] = uid
	claims[OpenIDKey] = openId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
