package handler

import (
	"net/http"
	echo "simpleapi-serverless/server/handler/echo"
	"simpleapi-serverless/server/svc"
	"time"

	"github.com/zeromicro/go-zero/rest"
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
					Path:    "/echo",
					Handler: echo.Echo(serverCtx),
				},
			},
			rest.WithPrefix("/api/simpleapi-serverless"),
		)
	}

}
