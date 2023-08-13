package user

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"uapply-micro/service/user/cmd/api/internal/logic/user"
	"uapply-micro/service/user/cmd/api/internal/svc"
)

func GetRegisterHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetRegisterLogic(r.Context(), ctx)
		resp, err := l.GetRegister()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
