// Code generated by jzero. DO NOT EDIT.

package server

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"gateway-sqlx/internal/config"
	"gateway-sqlx/internal/svc"

	usersvr "gateway-sqlx/internal/server/user"

	"gateway-sqlx/internal/pb/userpb"
)

func RegisterZrpc(c config.Config, ctx *svc.ServiceContext) *zrpc.RpcServer {
	s := zrpc.MustNewServer(c.Zrpc.RpcServerConf, func(grpcServer *grpc.Server) {
	    
		userpb.RegisterUserServer(grpcServer, usersvr.NewUser(ctx))

		if c.Zrpc.Mode == service.DevMode || c.Zrpc.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	return s
}