package cmd

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/jzero-io/jzero/core/configcenter/subscriber"
	"github.com/spf13/cobra"
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"

	"simpleapi/internal/config"
	"simpleapi/internal/handler"
	"simpleapi/internal/middleware"
	"simpleapi/internal/svc"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "simpleapi server",
	Long:  "simpleapi server",
	Run: func(cmd *cobra.Command, args []string) {
		cc := configurator.MustNewConfigCenter[config.Config](configurator.Config{
			Type: "yaml",
		}, subscriber.MustNewFsnotifySubscriber(cfgFile, subscriber.WithUseEnv(true)))
		c, err := cc.GetConfig()
		logx.Must(err)

		// set up logger
		logx.Must(logx.SetUp(c.Log.LogConf))

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
	printVersion()

	logx.Infof("Starting rest server at %s:%d...", c.Rest.Host, c.Rest.Port)
	group.Start()
}

func printBanner(c config.Config) {
	figure.NewColorFigure(c.Banner.Text, c.Banner.FontName, c.Banner.Color, true).Print()
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
