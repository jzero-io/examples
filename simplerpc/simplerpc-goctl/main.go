package main

import (
	"flag"
	"simplerpc-goctl/internal/config"
	"simplerpc-goctl/internal/server"
	"simplerpc-goctl/internal/svc"

	"github.com/common-nighthawk/go-figure"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/etc.yaml", "set config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := server.RegisterZrpc(c, ctx)
	defer s.Stop()

	printBanner(c)
	logx.Infof("Starting rpc server at %s...", c.ListenOn)
	s.Start()
}

func printBanner(c config.Config) {
	figure.NewColorFigure(c.Banner.Text, c.Banner.FontName, c.Banner.Color, true).Print()
}
