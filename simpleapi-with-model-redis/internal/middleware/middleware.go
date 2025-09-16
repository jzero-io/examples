package middleware

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"

	"simpleapi-with-model-redis/internal/svc"
)

func Register(server *rest.Server) {
	httpx.SetOkHandler(ResponseMiddleware)
	httpx.SetErrorHandlerCtx(ErrorMiddleware)
	httpx.SetValidator(NewValidator())
}

func NewMiddleware() svc.Middleware {
	return svc.Middleware{}
}
