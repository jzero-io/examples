package cmd

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/jzero-io/jzero/core/configcenter"
	"github.com/jzero-io/jzero/core/configcenter/subscriber"
	"github.com/jzero-io/jzero/core/swaggerv2"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"

	"simpleapi/plugins/helloworld/internal/config"
	"simpleapi/plugins/helloworld/internal/custom"
	"simpleapi/plugins/helloworld/internal/handler"
	"simpleapi/plugins/helloworld/internal/middleware"
	"simpleapi/plugins/helloworld/internal/svc"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "helloworld server",
	Long:  "helloworld server",
	Run: func(cmd *cobra.Command, args []string) {
		cc := configcenter.MustNewConfigCenter[config.Config](configcenter.Config{
			Type: "yaml",
		}, subscriber.MustNewFsnotifySubscriber(cmd.Flag("config").Value.String(), subscriber.WithUseEnv(true)))

		// set up logger
		logx.Must(logx.SetUp(cc.MustGetConfig().Log.LogConf))

		// print banner
		printBanner(cc.MustGetConfig().Banner)
		// print version
		printVersion()

		// create service context
		svcCtx := svc.NewServiceContext(cc)

		// create rest server
		restServer := rest.MustNewServer(svcCtx.ConfigCenter.MustGetConfig().Rest.RestConf)
		// create custom server
		customServer := custom.New()

		// register auto generated routes
		handler.RegisterHandlers(restServer, svcCtx)
		// register swagger routes
		swaggerv2.RegisterRoutes(restServer)
		// register middleware
		middleware.Register(restServer)

		group := service.NewServiceGroup()
		group.Add(restServer)
		group.Add(customServer)

		logx.Infof("Starting rest server at %s:%d...", cc.MustGetConfig().Rest.Host, cc.MustGetConfig().Rest.Port)
		group.Start()
	},
}

func printBanner(c config.BannerConf) {
	figure.NewColorFigure(c.Text, c.FontName, c.Color, true).Print()
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
