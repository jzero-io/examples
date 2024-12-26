package cmd

import (
	"os"
	"simplerpc/internal/config"
	"simplerpc/internal/middleware"
	"simplerpc/internal/server"
	"simplerpc/internal/svc"

	"github.com/common-nighthawk/go-figure"
	"github.com/jzero-io/jzero-contrib/dynamic_conf"
	"github.com/spf13/cobra"
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "simplerpc server",
	Long:  "simplerpc server",
	Run: func(cmd *cobra.Command, args []string) {
		ss, err := dynamic_conf.NewFsNotify(cfgFile)
		logx.Must(err)
		cc := configurator.MustNewConfigCenter[config.Config](configurator.Config{
			Type: "yaml",
		}, ss)
		c, err := cc.GetConfig()
		logx.Must(err)

		// set up logger
		if err := logx.SetUp(c.Log.LogConf); err != nil {
			logx.Must(err)
		}
		if c.Log.LogConf.Mode != "console" {
			logx.AddWriter(logx.NewWriter(os.Stdout))
		}

		svcCtx := svc.NewServiceContext(cc)
		run(svcCtx)
	},
}

func run(svcCtx *svc.ServiceContext) {
	c := svcCtx.MustGetConfig()

	zrpc := server.RegisterZrpc(c, svcCtx)
	middleware.Register(zrpc)

	group := service.NewServiceGroup()
	group.Add(zrpc)
	group.Add(svcCtx.Custom)

	printBanner(c)
	logx.Infof("Starting rpc server at %s...", c.Zrpc.ListenOn)
	group.Start()
}

func printBanner(c config.Config) {
	figure.NewColorFigure(c.Banner.Text, c.Banner.FontName, c.Banner.Color, true).Print()
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
