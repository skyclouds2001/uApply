package inter

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"uapply-micro/service/interviewer/cmd/api/internal/logic/inter"
	"uapply-micro/service/interviewer/cmd/api/internal/svc"
)

func GetOrgDepHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := inter.NewGetOrgDepLogic(r.Context(), svcCtx)
		resp, err := l.GetOrgDep()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
