package org

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"uapply-micro/service/organization/cmd/api/internal/logic/org"
	"uapply-micro/service/organization/cmd/api/internal/svc"
)

func GetAllDepHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := org.NewGetAllDepLogic(r.Context(), ctx)
		resp, err := l.GetAllDep()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
