package middleware

import (
	"github.com/zeromicro/go-zero/zrpc"
)

func RegisterZrpc(z *zrpc.RpcServer) {
	z.AddUnaryInterceptors(ServerValidationUnaryInterceptor)
}
