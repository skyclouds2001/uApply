package department

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"uapply-micro/service/department/cmd/api/internal/logic/department"
	"uapply-micro/service/department/cmd/api/internal/svc"
	"uapply-micro/service/department/cmd/api/internal/types"
)

func Announce2InterviewHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AInterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := department.NewAnnounce2InterviewLogic(r.Context(), ctx)
		resp, err := l.Announce2Interview(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
