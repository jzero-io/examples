package cmd

import (
	"os"
	"quickstart/desc/pb"
	"quickstart/internal/config"
	"quickstart/internal/handler"
	"quickstart/internal/middleware"
	"quickstart/internal/server"
	"quickstart/internal/svc"

	"github.com/common-nighthawk/go-figure"
	"github.com/jzero-io/jzero-contrib/gwx"
	"github.com/jzero-io/jzero-contrib/swaggerv2"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/proc"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/gateway"
	"golang.org/x/sync/errgroup"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "quickstart server",
	Long:  "quickstart server",
	Run: func(cmd *cobra.Command, args []string) {
		Start(cfgFile)
	},
}

func Start(cfgFile string) {
	var c config.Config
	conf.MustLoad(cfgFile, &c)
	config.C = c

	// write pb to local
	var err error
	c.Gateway.Upstreams[0].ProtoSets, err = gwx.WritePbToLocal(pb.Embed)
	if err != nil {
		logx.Must(err)
	}

	// set up logger
	if err = logx.SetUp(c.Log.LogConf); err != nil {
		logx.Must(err)
	}
	if c.Log.LogConf.Mode != "console" {
		logx.AddWriter(logx.NewWriter(os.Stdout))
	}

	ctx := svc.NewServiceContext(c)
	start(ctx)
}

func start(svcCtx *svc.ServiceContext) {
	zrpc := server.RegisterZrpc(svcCtx.Config, svcCtx)
	middleware.RegisterZrpc(zrpc)

	gw := gateway.MustNewServer(svcCtx.Config.Gateway.GatewayConf)
	middleware.RegisterGateway(gw)

	// gw add api routes
	handler.RegisterHandlers(gw.Server, svcCtx)

	// gw add swagger routes. If you do not want it, you can delete this line
	swaggerv2.RegisterRoutes(gw.Server)
	// gw add routes
	// You can use gw.Server.AddRoutes()

	group := service.NewServiceGroup()
	group.Add(zrpc)
	group.Add(gw)

	// shutdown listener
	waitExit := proc.AddShutdownListener(svcCtx.Custom.Stop)

	eg := errgroup.Group{}
	eg.Go(func() error {
		printBanner(svcCtx.Config)
		logx.Infof("Starting rpc server at %s...", svcCtx.Config.Zrpc.ListenOn)
		logx.Infof("Starting gateway server at %s:%d...", svcCtx.Config.Gateway.Host, svcCtx.Config.Gateway.Port)
		group.Start()
		return nil
	})

	// add custom start logic
	eg.Go(func() error {
		svcCtx.Custom.Start()
		return nil
	})

	if err := eg.Wait(); err != nil {
		panic(err)
	}

	waitExit()
}

func printBanner(c config.Config) {
	figure.NewColorFigure(c.Banner.Text, c.Banner.FontName, c.Banner.Color, true).Print()
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
