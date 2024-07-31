package middleware

import (
	"github.com/zeromicro/go-zero/gateway"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/zrpc"
)

func RegisterZrpc(z *zrpc.RpcServer) {
	z.AddUnaryInterceptors(ServerValidationUnaryInterceptor)
}

func RegisterGateway(g *gateway.Server) {
	httpx.SetErrorHandler(ErrorHandler)
	g.Use(ResponseHandler)
}
