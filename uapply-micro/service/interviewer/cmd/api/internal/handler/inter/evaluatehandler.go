package inter

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"uapply-micro/service/interviewer/cmd/api/internal/logic/inter"
	"uapply-micro/service/interviewer/cmd/api/internal/svc"
	"uapply-micro/service/interviewer/cmd/api/internal/types"
)

func EvaluateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MarkReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := inter.NewEvaluateLogic(r.Context(), svcCtx)
		resp, err := l.Evaluate(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
