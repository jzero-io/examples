package serverless

import (
	"os"
	"path/filepath"
	"simpleapi-serverless/server/config"
	"simpleapi-serverless/server/handler"
	"simpleapi-serverless/server/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var (
	server *rest.Server
)

func init() {
	var c config.Config

	if err := conf.Load(filepath.Join("plugins", "simpleapi-serverless", "etc", "etc.yaml"), &c); err != nil {
		panic(err)
	}
	config.C = c

	if err := logx.SetUp(c.Log.LogConf); err != nil {
		logx.Must(err)
	}
	if c.Log.LogConf.Mode != "console" {
		logx.AddWriter(logx.NewWriter(os.Stdout))
	}

	server = rest.MustNewServer(c.Rest.RestConf)

	svcCtx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, svcCtx)
}

func Routes() []rest.Route {
	return server.Routes()
}
