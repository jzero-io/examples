package cmd

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/jzero-io/jzero/core/configcenter/subscriber"
	"github.com/spf13/cobra"
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"simplerpc/internal/config"
	"simplerpc/internal/custom"
	"simplerpc/internal/global"
	"simplerpc/internal/middleware"
	"simplerpc/internal/server"
	"simplerpc/internal/svc"
	"simplerpc/plugins"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "simplerpc server",
	Long:  "simplerpc server",
	Run: func(cmd *cobra.Command, args []string) {
		cc := configurator.MustNewConfigCenter[config.Config](configurator.Config{
			Type: "yaml",
		}, subscriber.MustNewFsnotifySubscriber(cfgFile, subscriber.WithUseEnv(true)))
		c, err := cc.GetConfig()
		logx.Must(err)

		// set up logger
		logx.Must(logx.SetUp(c.Log.LogConf))

		printBanner(c)
		printVersion()

		svcCtx := svc.NewServiceContext(cc)
		global.ServiceContext = *svcCtx
		run(svcCtx)
	},
}

func run(svcCtx *svc.ServiceContext) {
	zrpcServer := zrpc.MustNewServer(svcCtx.MustGetConfig().Zrpc.RpcServerConf, func(grpcServer *grpc.Server) {
		server.RegisterZrpcServer(grpcServer, svcCtx)
		// register plugins
		plugins.LoadPlugins(grpcServer, svcCtx)
		if svcCtx.MustGetConfig().Zrpc.Mode == service.DevMode || svcCtx.MustGetConfig().Zrpc.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	ctm := custom.New(zrpcServer)
	ctm.Init()

	middleware.Register(zrpcServer)

	group := service.NewServiceGroup()
	group.Add(zrpcServer)
	group.Add(ctm)

	logx.Infof("Starting rpc server at %s...", svcCtx.MustGetConfig().Zrpc.ListenOn)
	group.Start()
}

func printBanner(c config.Config) {
	figure.NewColorFigure(c.Banner.Text, c.Banner.FontName, c.Banner.Color, true).Print()
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
