package department

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"uapply-micro/service/department/cmd/api/internal/logic/department"
	"uapply-micro/service/department/cmd/api/internal/svc"
	"uapply-micro/service/department/cmd/api/internal/types"
)

func ADDInterviewerHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddInterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := department.NewADDInterviewerLogic(r.Context(), ctx)
		resp, err := l.ADDInterviewer(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
