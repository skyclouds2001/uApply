package user

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"uapply-micro/service/user/cmd/api/internal/logic/user"
	"uapply-micro/service/user/cmd/api/internal/svc"
	"uapply-micro/service/user/cmd/api/internal/types"
)

func GetDepsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetDepsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewGetDepsLogic(r.Context(), ctx)
		resp, err := l.GetDeps(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
