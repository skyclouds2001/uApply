package department

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"uapply-micro/service/department/cmd/api/internal/logic/department"
	"uapply-micro/service/department/cmd/api/internal/svc"
)

func GetOutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := department.NewGetOutLogic(r.Context(), svcCtx)
		resp, err := l.GetOut()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
