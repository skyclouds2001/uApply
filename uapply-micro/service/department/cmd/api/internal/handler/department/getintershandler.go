package department

import (
	"net/http"

	"uapply-micro/service/department/cmd/api/internal/logic/department"
	"uapply-micro/service/department/cmd/api/internal/svc"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetIntersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := department.NewGetIntersLogic(r.Context(), svcCtx)
		resp, err := l.GetInters()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
