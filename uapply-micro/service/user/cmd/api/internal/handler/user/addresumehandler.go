package user

import (
	"net/http"
	"uapply-micro/common/errorx"

	"github.com/tal-tech/go-zero/rest/httpx"
	"uapply-micro/service/user/cmd/api/internal/logic/user"
	"uapply-micro/service/user/cmd/api/internal/svc"
	"uapply-micro/service/user/cmd/api/internal/types"
)

func AddResumeHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddResumeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewError(err.Error(), errorx.ParamInvalid))
			return
		}

		l := user.NewAddResumeLogic(r.Context(), ctx)
		resp, err := l.AddResume(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
