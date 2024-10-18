// Code generated by jzero. DO NOT EDIT.

package server

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"simplerpc-goctl/internal/config"
	"simplerpc-goctl/internal/svc"

	hellosvr "simplerpc-goctl/internal/server/hello"

	"simplerpc-goctl/internal/pb/hellopb"
)

func RegisterZrpc(c config.Config, ctx *svc.ServiceContext) *zrpc.RpcServer {
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
	    
		hellopb.RegisterHelloServer(grpcServer, hellosvr.NewHello(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	return s
}