// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	inter "uapply-micro/service/interviewer/cmd/api/internal/handler/inter"
	"uapply-micro/service/interviewer/cmd/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/inter/login/:code",
				Handler: LoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/inter/gerorgdep",
				Handler: inter.GetOrgDepHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/inter/getresume/:num",
				Handler: inter.GetUserResumeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/inter/evaluate/:num",
				Handler: inter.EvaluateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/inter/vieweva/:num",
				Handler: inter.ViewEvaHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/inter/turn",
				Handler: inter.GetUserTurnHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/inter/get-data/:depId",
				Handler: inter.GetDepDataHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
