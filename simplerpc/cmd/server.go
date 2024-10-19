package cmd

import (
	"os"
	"simplerpc/internal/config"
	"simplerpc/internal/middleware"
	"simplerpc/internal/server"
	"simplerpc/internal/svc"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "simplerpc server",
	Long:  "simplerpc server",
	Run: func(cmd *cobra.Command, args []string) {
		var c config.Config
		conf.MustLoad(cfgFile, &c)
		config.C = c

		// set up logger
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
	zrpc := server.RegisterZrpc(svcCtx.Config, svcCtx)
	middleware.Register(zrpc)

	group := service.NewServiceGroup()
	group.Add(zrpc)
	group.Add(svcCtx.Custom)

	printBanner(svcCtx.Config)
	logx.Infof("Starting rpc server at %s...", svcCtx.Config.Zrpc.ListenOn)
	group.Start()
}

func printBanner(c config.Config) {
	figure.NewColorFigure(c.Banner.Text, c.Banner.FontName, c.Banner.Color, true).Print()
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
