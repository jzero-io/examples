// Code generated by jzero. DO NOT EDIT.
package handler

import (
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/rest"
	"simpleapi/internal/svc"

	version "simpleapi/internal/handler/version"
)

var (
	_ = http.StatusOK
	_ = time.Now()
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	{
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

}
