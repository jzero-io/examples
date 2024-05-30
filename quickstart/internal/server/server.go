package server

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"quickstart/internal/config"
	"quickstart/internal/pb/hellopb"
	helloserver "quickstart/internal/server/hello"
	"quickstart/internal/svc"
)

func GetZrpcServer(c config.Config, ctx *svc.ServiceContext) *zrpc.RpcServer {
	s := zrpc.MustNewServer(c.Zrpc.RpcServerConf, func(grpcServer *grpc.Server) {
		hellopb.RegisterHelloServer(grpcServer, helloserver.NewHelloServer(ctx))
		if c.Zrpc.Mode == service.DevMode || c.Zrpc.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	return s
}
