package org

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"uapply-micro/service/organization/cmd/api/internal/logic/org"
	"uapply-micro/service/organization/cmd/api/internal/svc"
	"uapply-micro/service/organization/cmd/api/internal/types"
)

func GetDepHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := org.NewGetDepLogic(r.Context(), ctx)
		resp, err := l.GetDep(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
