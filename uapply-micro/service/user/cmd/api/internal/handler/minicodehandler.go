package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"uapply-micro/service/user/cmd/api/internal/logic"
	"uapply-micro/service/user/cmd/api/internal/svc"
)

func MiniCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewMiniCodeLogic(r.Context(), svcCtx)
		resp, err := l.MiniCode()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
