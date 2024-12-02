package serverless

import (
	"path/filepath"
	"simpleapi-serverless/server/config"
	"simpleapi-serverless/server/handler"
	"simpleapi-serverless/server/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var (
	server *rest.Server
)

func serverless() {
	var c config.Config

	if err := conf.Load(filepath.Join("plugins", "simpleapi-serverless", "etc", "etc.yaml"), &c); err != nil {
		panic(err)
	}
	config.C = c

	server = rest.MustNewServer(c.Rest.RestConf)

	svcCtx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, svcCtx)
}

func Routes() []rest.Route {
	serverless()
	return server.Routes()
}
