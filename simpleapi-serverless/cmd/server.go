package cmd

import (
	"os"
	"simpleapi-serverless/server/config"
	"simpleapi-serverless/server/handler"
	"simpleapi-serverless/server/middleware"
	"simpleapi-serverless/server/svc"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "simpleapi-serverless server",
	Long:  "simpleapi-serverless server",
	Run: func(cmd *cobra.Command, args []string) {
		var c config.Config
		conf.MustLoad(CfgFile, &c)
		config.C = c

		if err := logx.SetUp(c.Log.LogConf); err != nil {
			logx.Must(err)
		}
		if c.Log.LogConf.Mode != "console" {
			logx.AddWriter(logx.NewWriter(os.Stdout))
		}

		ctx := svc.NewServiceContext(c)
		run(ctx)
	},
}

func run(svcCtx *svc.ServiceContext) {
	server := rest.MustNewServer(svcCtx.Config.Rest.RestConf)
	middleware.Register(server)

	handler.RegisterHandlers(server, svcCtx)

	svcCtx.Custom.AddRoutes(server)

	group := service.NewServiceGroup()
	group.Add(server)
	group.Add(svcCtx.Custom)

	printBanner(svcCtx.Config)
	logx.Infof("Starting rest server at %s:%d...", svcCtx.Config.Rest.Host, svcCtx.Config.Rest.Port)
	group.Start()
}

func printBanner(c config.Config) {
	figure.NewColorFigure(c.Banner.Text, c.Banner.FontName, c.Banner.Color, true).Print()
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
