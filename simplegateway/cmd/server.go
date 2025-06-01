package cmd

import (
	"path/filepath"

	"github.com/common-nighthawk/go-figure"
	"github.com/jzero-io/jzero/core/configcenter/subscriber"
	"github.com/jzero-io/jzero/core/embedx"
	"github.com/spf13/cobra"
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/gateway"

	"simplegateway/desc/pb"
	"simplegateway/internal/config"
	"simplegateway/internal/middleware"
	"simplegateway/internal/server"
	"simplegateway/internal/svc"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "simplegateway server",
	Long:  "simplegateway server",
	Run: func(cmd *cobra.Command, args []string) {
		cc := configurator.MustNewConfigCenter[config.Config](configurator.Config{
			Type: "yaml",
		}, subscriber.MustNewFsnotifySubscriber(cfgFile, subscriber.WithUseEnv(true)))
		c, err := cc.GetConfig()
		logx.Must(err)

		// set up logger
		logx.Must(logx.SetUp(c.Log.LogConf))

		// write pb to local
		c.Gateway.Upstreams[0].ProtoSets, err = embedx.WriteToLocal(pb.Embed, embedx.WithFileMatchFunc(func(path string) bool {
			return filepath.Ext(path) == ".pb"
		}))
		logx.Must(err)

		svcCtx := svc.NewServiceContext(cc)
		run(svcCtx)
	},
}

func run(svcCtx *svc.ServiceContext) {
	c := svcCtx.MustGetConfig()

	zrpc := server.RegisterZrpc(c, svcCtx)
	gw := gateway.MustNewServer(c.Gateway.GatewayConf, middleware.WithHeaderProcessor())

	// register middleware
	middleware.Register(zrpc, gw)

	// gw add custom routes
	svcCtx.Custom.AddRoutes(gw)

	group := service.NewServiceGroup()
	group.Add(zrpc)
	group.Add(gw)
	group.Add(svcCtx.Custom)

	printBanner(c)
	printVersion()

	logx.Infof("Starting rpc server at %s...", c.Zrpc.ListenOn)
	logx.Infof("Starting gateway server at %s:%d...", c.Gateway.Host, c.Gateway.Port)

	group.Start()
}

func printBanner(c config.Config) {
	figure.NewColorFigure(c.Banner.Text, c.Banner.FontName, c.Banner.Color, true).Print()
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
