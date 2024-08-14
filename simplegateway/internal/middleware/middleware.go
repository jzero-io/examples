package middleware

import (
	"github.com/zeromicro/go-zero/gateway"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/zrpc"
)

func Register(z *zrpc.RpcServer, gw *gateway.Server) {
	z.AddUnaryInterceptors(ValidatorMiddleware)
	z.AddUnaryInterceptors(WithValueMiddleware)

	httpx.SetErrorHandler(ErrorMiddleware)
	gw.Use(ResponseMiddleware)
}
