package cmd

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/jzero-io/jzero/core/configcenter"
	"github.com/jzero-io/jzero/core/configcenter/subscriber"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"

	"simpleapi-with-model-redis/internal/config"
	"simpleapi-with-model-redis/internal/custom"
	"simpleapi-with-model-redis/internal/handler"
	"simpleapi-with-model-redis/internal/middleware"
	"simpleapi-with-model-redis/internal/svc"
	"simpleapi-with-model-redis/plugins"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "simpleapi-with-model-redis server",
	Long:  "simpleapi-with-model-redis server",
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

		svcCtx := svc.NewServiceContext(cc)

		restServer := rest.MustNewServer(svcCtx.ConfigCenter.MustGetConfig().Rest.RestConf)
		customServer := custom.New()

		handler.RegisterHandlers(restServer, svcCtx)
		middleware.Register(restServer)

		// load plugins
		plugins.LoadPlugins(restServer, svcCtx)

		group := service.NewServiceGroup()
		group.Add(restServer)
		group.Add(customServer)

		group.Start()
	},
}

func printBanner(c config.BannerConf) {
	figure.NewColorFigure(c.Text, c.FontName, c.Color, true).Print()
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
