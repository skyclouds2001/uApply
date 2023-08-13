package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"uapply-micro/service/interviewer/cmd/api/internal/logic"
	"uapply-micro/service/interviewer/cmd/api/internal/svc"
	"uapply-micro/service/interviewer/cmd/api/internal/types"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
