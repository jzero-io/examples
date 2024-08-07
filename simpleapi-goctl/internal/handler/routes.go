// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	version "simpleapi-goctl/internal/handler/version"
	"simpleapi-goctl/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/version",
				Handler: version.GetVersion(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)
}
