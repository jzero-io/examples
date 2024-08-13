package cmd

import (
	"gateway-sqlx/desc/pb"
	"gateway-sqlx/internal/config"
	"gateway-sqlx/internal/middleware"
	"gateway-sqlx/internal/server"
	"gateway-sqlx/internal/svc"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/jzero-io/jzero-contrib/gwx"
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
	Short: "gateway-sqlx server",
	Long:  "gateway-sqlx server",
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

	gw := gateway.MustNewServer(svcCtx.Config.Gateway.GatewayConf, middleware.WithHeaderProcessor())
	middleware.RegisterGateway(gw)

	// gw add custom routes
	svcCtx.Custom.AddRoutes(gw)

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
