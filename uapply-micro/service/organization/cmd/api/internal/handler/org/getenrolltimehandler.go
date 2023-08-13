package org

import (
	"net/http"

	"uapply-micro/service/organization/cmd/api/internal/logic/org"
	"uapply-micro/service/organization/cmd/api/internal/svc"
	"uapply-micro/service/organization/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetEnrollTimeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := org.NewGetEnrollTimeLogic(r.Context(), svcCtx)
		resp, err := l.GetEnrollTime(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
