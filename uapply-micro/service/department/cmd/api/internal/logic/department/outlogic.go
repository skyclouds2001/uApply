package department

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"uapply-micro/common/errorx"
	"uapply-micro/common/filed"
	"uapply-micro/service/department/cmd/api/internal/logic"
	"uapply-micro/service/department/cmd/api/internal/svc"
	"uapply-micro/service/department/cmd/api/internal/types"
	"uapply-micro/service/department/model"
	"uapply-micro/service/user/cmd/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111" // 引入sms
)

type OutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) OutLogic {
	return OutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Out 淘汰用户
func (l *OutLogic) Out(req types.OutReq) (*types.Msg, error) {
	if len(req.Uid) == 0 {
		return nil, errorx.NewError("参数不能为空", errorx.ParamInvalid)
	}
	// 拿部门id
	depIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value(logic.DepIdKey)))
	depId, err := depIdNumber.Int64()
	if err != nil {
		logx.Error(err)
		return nil, errorx.NewSysError()
	}
	if depId <= 0 {
		return nil, errorx.NewError("登录信息错误", errorx.Fail)
	}

	// 检验面试轮次
	var num string
	if req.Num == 1 {
		num = filed.First
	} else if req.Num == 2 {
		num = filed.Second
	} else {
		return nil, errorx.NewError("面试只有一轮和二轮", errorx.ParamInvalid)
	}

	err = l.svcCtx.RegModel.UpdateMany(req.Uid, int(depId), l.svcCtx.Mysqlx, num, filed.OUT)
	err = l.svcCtx.RegModel.UpdateMany(req.Uid, int(depId), l.svcCtx.Mysqlx, filed.Final, filed.OUT)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errorx.NewSysError()
	}

	// 转换
	var ids []int64
	for _, i := range req.Uid {
		ids = append(ids, int64(i))
	}

	return &types.Msg{
		Msg: "淘汰成功",
	}, nil
}

func (l *OutLogic) sendOutSms(users []*user.UserInfo, depName, orgName string) {
	smsConf := l.svcCtx.Config.Sms
	credential := common.NewCredential(smsConf.SecretId, smsConf.SecretKey)

	// 客户端配置
	cpf := profile.NewClientProfile()
	client, err := sms.NewClient(credential, smsConf.Region, cpf)
	if err != nil {
		panic(err)
	}

	for _, userInfo := range users {
		// request
		request := sms.NewSendSmsRequest()
		request.PhoneNumberSet = common.StringPtrs([]string{userInfo.PhoneNum})
		request.SmsSdkAppId = common.StringPtr(smsConf.AppId)
		request.TemplateId = common.StringPtr(smsConf.Templates.Out)
		request.SignName = common.StringPtr(smsConf.SignName)
		// 模板参数
		request.TemplateParamSet = common.StringPtrs([]string{
			userInfo.Name,
			orgName + depName,
		})

		// 忽略发送的错误，直接下一个，因为有可能当前用户错填了电话，导致一直不成功
		// 因此在这里跳过该用户，记录日志，发送下一个
		resp, err := client.SendSms(request)
		if err != nil {
			logx.Infof(
				"send out sms to uid [%d] error, phone number [%s], reason [%v]\n",
				userInfo.Uid,
				userInfo.PhoneNum,
				err,
			)
			continue
		}
		respCode := *resp.Response.SendStatusSet[0].Code
		if *resp.Response.SendStatusSet[0].Code != "Ok" {
			logx.Infof(
				"send out sms to uid [%d] failed, phone number [%s], response code [%s]\n",
				userInfo.Uid,
				userInfo.PhoneNum,
				respCode,
			)
		}
	}
}
