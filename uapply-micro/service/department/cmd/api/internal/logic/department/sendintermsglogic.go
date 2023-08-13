package department

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"strconv"
	"time"
	"uapply-micro/common/errorx"
	"uapply-micro/common/filed"
	"uapply-micro/service/department/cmd/api/internal/logic"
	"uapply-micro/service/department/model"
	"uapply-micro/service/user/cmd/rpc/user"
	"unicode/utf8"

	"uapply-micro/service/department/cmd/api/internal/svc"
	"uapply-micro/service/department/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SendInterMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendInterMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) SendInterMsgLogic {
	return SendInterMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type OrgDepData struct {
	orgName string
	depId   int
	depName string
}

func (l *SendInterMsgLogic) SendInterMsg(req types.SendInterNoticeReq) (resp *types.Msg, err error) {
	// 部门id
	depIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value(logic.DepIdKey)))
	depId, err := depIdNumber.Int64()
	if err != nil {
		logx.Error(err)
		return nil, errorx.NewSysError()
	}
	if depId <= 0 {
		return nil, errorx.NewError("登录信息错误", errorx.Fail)
	}

	depName := fmt.Sprintf("%v", l.ctx.Value(logic.DepNameKey))
	orgName := fmt.Sprintf("%v", l.ctx.Value(logic.OrgNameKey))

	if depName == "" || orgName == "" {
		return nil, errorx.NewError("登录信息错误", errorx.Fail)
	}

	// 检验面试轮次
	if req.Num != 1 && req.Num != 2 {
		return nil, errorx.NewError("面试只有一轮和二轮", errorx.ParamInvalid)
	}

	// 获取面试配置信息
	interConf, err := l.svcCtx.InterConfModel.FindOneByDepIdTurn(depId, int64(req.Num))
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errorx.NewCodeError(errorx.NotFound, "未找到面试配置信息")
		}
		return nil, errorx.NewSysError()
	}
	// 腾讯云个人认证模板变量字符数不超过12
	if utf8.RuneCountInString(interConf.InterAddress) > 12 {
		return nil, errorx.NewCodeError(errorx.Fail, "面试地点不能超过12个字符")
	}

	// 获取用户信息
	uids := make([]int64, 0, len(req.Uids))
	for _, uid := range req.Uids {
		uids = append(uids, int64(uid))
	}
	usersResp, err := l.svcCtx.UserRpc.GetUsers(l.ctx, &user.GetSexReq{Ids: uids})
	if err != nil {
		return nil, errorx.HandleGrpcErrorToHttp(err)
	}

	orgDepData := &OrgDepData{
		orgName: orgName,
		depId:   int(depId),
		depName: depName,
	}

	go l.sendInterMsgSms(usersResp.Users, orgDepData, req.Num, interConf)

	resp = &types.Msg{
		Msg: "已发送",
	}

	return
}

func (l *SendInterMsgLogic) sendInterMsgSms(users []*user.UserInfo, orgDepData *OrgDepData, turn int, interConf *model.InterConf) {
	smsConf := l.svcCtx.Config.Sms
	credential := common.NewCredential(smsConf.SecretId, smsConf.SecretKey)

	// 客户端配置
	cpf := profile.NewClientProfile()
	client, err := sms.NewClient(credential, smsConf.Region, cpf)
	if err != nil {
		panic(err)
	}

	var successUids []int

	var turnStr string
	var templateId *string
	// 根据轮次选择模板
	if turn == 1 {
		templateId = common.StringPtr(smsConf.Templates.First)
		turnStr = filed.First
	} else if turn == 2 {
		templateId = common.StringPtr(smsConf.Templates.Second)
		turnStr = filed.Second
	}
	startTime := formatTime(interConf.Start)
	endTime := formatTime(interConf.End)

	for _, userInfo := range users {
		// request
		request := sms.NewSendSmsRequest()
		request.PhoneNumberSet = common.StringPtrs([]string{userInfo.PhoneNum})
		request.SmsSdkAppId = common.StringPtr(smsConf.AppId)
		request.SignName = common.StringPtr(smsConf.SignName)
		request.TemplateId = templateId

		// 组装模板参数
		request.TemplateParamSet = common.StringPtrs([]string{
			userInfo.Name,                       // 1
			orgDepData.orgName,                  // 2
			orgDepData.depName,                  // 3
			strconv.FormatInt(userInfo.Uid, 10), // 4
			startTime,                           // 5
			endTime,                             // 6
			interConf.InterAddress,              // 7
			interConf.ContactPhone,              // 8
		})

		// 忽略发送的错误，直接下一个，因为有可能当前用户错填了电话，导致一直不成功
		// 因此在这里跳过该用户，记录日志，发送下一个
		resp, err := client.SendSms(request)
		if err != nil {
			logx.Errorf(
				"send pass [%d] sms to uid [%d] error, phone number [%s], reason [%v]\n",
				turn,
				userInfo.Uid,
				userInfo.PhoneNum,
				err,
			)
			continue
		}
		respCode := *resp.Response.SendStatusSet[0].Code
		if *resp.Response.SendStatusSet[0].Code != "Ok" {
			logx.Errorf(
				"send to pass [%d] sms uid [%d] failed, phone number [%s], response code [%s]\n",
				turn,
				userInfo.Uid,
				userInfo.PhoneNum,
				respCode,
			)
			continue
		}
		successUids = append(successUids, int(userInfo.Uid))
	}
	// 设置状态
	if len(successUids) > 0 {
		err = l.svcCtx.RegModel.UpdateMany(successUids, orgDepData.depId, l.svcCtx.Mysqlx, turnStr, filed.ARRANGED)
		if err != nil && !errors.Is(err, model.ErrNotFound) {
			logx.Errorf(
				"update [%s] of dep [%d] failed, caused by [%s]",
				turnStr,
				orgDepData.depId,
				err.Error(),
			)
		}
	}
}

func formatTime(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("01月02日15:04")
}
