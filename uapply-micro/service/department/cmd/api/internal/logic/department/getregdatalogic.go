package department

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"uapply-micro/common/errorx"
	"uapply-micro/common/filed"
	"uapply-micro/service/department/cmd/api/internal/logic"
	"uapply-micro/service/user/cmd/rpc/user"

	"uapply-micro/service/department/cmd/api/internal/svc"
	"uapply-micro/service/department/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
	//myredis "github.com/go-redis/redis"
)

type GetRegDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRegDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetRegDataLogic {
	return GetRegDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRegDataLogic) GetRegData() (*types.DataRsp, error) {
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

	deps, err := l.svcCtx.RegModel.FindDepAll(depId)
	var rsp = &types.DataRsp{
		RegSum:     len(deps),
		RegMale:    0,
		RegFemale:  0,
		PassSum:    0,
		PassMale:   0,
		PassFemale: 0,
	}
	// 通过uid查性别
	var uids []int64
	for _, dep := range deps {
		uids = append(uids, dep.Uid)
	}
	if len(uids) != 0 {
		sex, err := l.svcCtx.UserRpc.GetUserSex(l.ctx, &user.GetSexReq{Ids: uids})
		if err != nil {
			return nil, errorx.HandleGrpcErrorToHttp(err)
		}
		// 加锁
		//redisConn := redis.NewRedis(l.svcCtx.Config.Redis.Host, l.svcCtx.Config.Redis.Type, l.svcCtx.Config.Redis.Pass)
		redisConn := redis.New(l.svcCtx.Config.Redis.Host, redis.WithPass(l.svcCtx.Config.Redis.Pass))
		redisLockKey := fmt.Sprintf("%s%d", "dep.reg.data", depId)
		redisLock := redis.NewRedisLock(redisConn, redisLockKey)
		// 2. 可选操作，设置 redislock 过期时间
		redisLock.SetExpire(10)
		if ok, err := redisLock.Acquire(); !ok || err != nil {
			logx.Info("redis锁失败", err)
			return nil, errorx.NewError("当前有其他用户正在进行操作，请稍后重试", errorx.Fail)
		}
		for i, user := range deps {
			// 男生
			if sex.Sexs[i] == filed.MALE {
				rsp.RegMale++
				if user.FinalStatus == filed.PASS {
					rsp.PassMale++
				}
			}
			// 女生
			if sex.Sexs[i] == filed.FEMALE {
				rsp.RegFemale++
				if user.FinalStatus == filed.PASS {
					rsp.PassFemale++
				}
			}
		}
		rsp.PassSum = rsp.PassMale + rsp.PassFemale
		redisLock.Release()
	} else {
		rsp.RegSum = 0
	}
	return rsp, nil
}
