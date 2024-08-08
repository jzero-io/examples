package main

import (
	"flag"
	"simpleapi-goctl/internal/config"
	"simpleapi-goctl/internal/handler"
	"simpleapi-goctl/internal/svc"

	"github.com/common-nighthawk/go-figure"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/etc.yaml", "set config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	printBanner(c)
	logx.Infof("Starting rest server at %s:%d...", c.Host, c.Port)
	server.Start()
}

func printBanner(c config.Config) {
	figure.NewColorFigure(c.Banner.Text, c.Banner.FontName, c.Banner.Color, true).Print()
}
