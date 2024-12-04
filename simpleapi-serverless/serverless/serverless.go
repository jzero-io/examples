package serverless

import (
	"path/filepath"
	"simpleapi-serverless/server/config"
	"simpleapi-serverless/server/handler"
	"simpleapi-serverless/server/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

// Serverless please replace coreSvcCtx any type to real core svcCtx
func Serverless(server *rest.Server, coreSvcCtx any) {
	var c config.Config

	if err := conf.Load(filepath.Join("plugins", "simpleapi-serverless", "etc", "etc.yaml"), &c); err != nil {
		panic(err)
	}
	config.C = c

	svcCtx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, svcCtx)
}
