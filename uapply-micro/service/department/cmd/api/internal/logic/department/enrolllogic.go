package department

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"uapply-micro/common/errorx"
	"uapply-micro/common/filed"
	"uapply-micro/service/department/cmd/api/internal/logic"
	"uapply-micro/service/department/model"
	"uapply-micro/service/user/cmd/rpc/user"

	"uapply-micro/service/department/cmd/api/internal/svc"
	"uapply-micro/service/department/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type EnrollLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEnrollLogic(ctx context.Context, svcCtx *svc.ServiceContext) EnrollLogic {
	return EnrollLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Enroll 录取用户
func (l *EnrollLogic) Enroll(req types.EnrollReq) (*types.Msg, error) {
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

	// 部门名称
	depName := fmt.Sprintf("%v", l.ctx.Value(logic.DepNameKey))
	// 组织名称
	orgName := fmt.Sprintf("%v", l.ctx.Value(logic.OrgNameKey))

	// 获取面试配置信息
	interConf, err := l.svcCtx.InterConfModel.FindOneByDepIdTurn(depId, 2)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errorx.NewCodeError(errorx.NotFound, "未找到面试配置信息")
		}
		return nil, errorx.NewSysError()
	}

	// 获取用户信息
	uids := make([]int64, 0, len(req.Uid))
	for _, uid := range req.Uid {
		uids = append(uids, int64(uid))
	}
	usersResp, err := l.svcCtx.UserRpc.GetUsers(l.ctx, &user.GetSexReq{Ids: uids})
	if err != nil {
		return nil, errorx.HandleGrpcErrorToHttp(err)
	}

	err = l.svcCtx.RegModel.UpdateMany(req.Uid, int(depId), l.svcCtx.Mysqlx, filed.Second, filed.PASS)
	err = l.svcCtx.RegModel.UpdateMany(req.Uid, int(depId), l.svcCtx.Mysqlx, filed.Final, filed.PASS)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errorx.NewSysError()
	}

	go l.sendEnrollSms(usersResp.Users, depName, orgName, interConf)

	return &types.Msg{
		Msg: "录取成功",
	}, nil
}

func (l *EnrollLogic) sendEnrollSms(users []*user.UserInfo, depName, orgName string, interConf *model.InterConf) {
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
		request.TemplateId = common.StringPtr(smsConf.Templates.Enroll)
		request.SignName = common.StringPtr(smsConf.SignName)
		// 模板参数
		request.TemplateParamSet = common.StringPtrs([]string{
			userInfo.Name,
			orgName + depName,
			interConf.ContactPhone,
		})

		// 忽略发送的错误，直接下一个，因为有可能当前用户错填了电话，导致一直不成功
		// 因此在这里跳过该用户，记录日志，发送下一个
		resp, err := client.SendSms(request)
		if err != nil {
			logx.Infof(
				"send enroll sms to uid [%d] error, phone number [%s], reason [%v]\n",
				userInfo.Uid,
				userInfo.PhoneNum,
				err,
			)
			continue
		}
		respCode := *resp.Response.SendStatusSet[0].Code
		if *resp.Response.SendStatusSet[0].Code != "Ok" {
			logx.Infof(
				"send enroll sms to uid [%d] failed, phone number [%s], response code [%s]\n",
				userInfo.Uid,
				userInfo.PhoneNum,
				respCode,
			)
		}
	}

}
