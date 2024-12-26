package cmd

import (
	"os"
	"simpleapi/internal/config"
	"simpleapi/internal/handler"
	"simpleapi/internal/middleware"
	"simpleapi/internal/svc"

	"github.com/common-nighthawk/go-figure"
	"github.com/jzero-io/jzero-contrib/dynamic_conf"
	"github.com/spf13/cobra"
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "simpleapi server",
	Long:  "simpleapi server",
	Run: func(cmd *cobra.Command, args []string) {
		ss, err := dynamic_conf.NewFsNotify(cfgFile, dynamic_conf.WithUseEnv(true))
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

	server := rest.MustNewServer(c.Rest.RestConf)
	middleware.Register(server)

	// server add api handlers
	handler.RegisterHandlers(server, svcCtx)

	// server add custom routes
	svcCtx.Custom.AddRoutes(server)

	group := service.NewServiceGroup()
	group.Add(server)
	group.Add(svcCtx.Custom)

	printBanner(c)
	logx.Infof("Starting rest server at %s:%d...", c.Rest.Host, c.Rest.Port)
	group.Start()
}

func printBanner(c config.Config) {
	figure.NewColorFigure(c.Banner.Text, c.Banner.FontName, c.Banner.Color, true).Print()
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
