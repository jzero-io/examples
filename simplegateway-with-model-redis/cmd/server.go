package cmd

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/jzero-io/jzero/core/configcenter"
	"github.com/jzero-io/jzero/core/configcenter/subscriber"
	"github.com/jzero-io/jzero/core/stores/migrate"
	"github.com/jzero-io/jzero/core/swaggerv2"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/gateway"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"simplegateway-with-model-redis/desc/pb"
	"simplegateway-with-model-redis/internal/config"
	"simplegateway-with-model-redis/internal/custom"
	"simplegateway-with-model-redis/internal/middleware"
	"simplegateway-with-model-redis/internal/server"
	"simplegateway-with-model-redis/internal/svc"
	"simplegateway-with-model-redis/plugins"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "simplegateway-with-model-redis server",
	Long:  "simplegateway-with-model-redis server",
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

		m, err := migrate.NewMigrate(cc.MustGetConfig().Sqlx.SqlConf)
		logx.Must(err)
		defer m.Close()
		logx.Must(m.Up())

		// create service context
		svcCtx := svc.NewServiceContext(cc)

		// write protoSets to local
		protoSets, err := pb.WriteToLocal(pb.Embed)
		logx.Must(err)
		cc.MustGetConfig().Gateway.Upstreams[0].ProtoSets = protoSets

		// create zrpc server
		zrpcServer := zrpc.MustNewServer(cc.MustGetConfig().Zrpc.RpcServerConf, func(grpcServer *grpc.Server) {
			server.RegisterZrpcServer(grpcServer, svcCtx)
			// register plugins
			plugins.LoadPlugins(grpcServer, svcCtx)
			if cc.MustGetConfig().Zrpc.Mode == service.DevMode || cc.MustGetConfig().Zrpc.Mode == service.TestMode {
				reflection.Register(grpcServer)
			}
		})
		// create gateway server
		gatewayServer := gateway.MustNewServer(cc.MustGetConfig().Gateway.GatewayConf, middleware.WithHeaderProcessor())
		// register swagger routes
		swaggerv2.RegisterRoutes(gatewayServer.Server)
		// // create custom server
		customServer := custom.New()

		// register middleware
		middleware.Register(zrpcServer, gatewayServer)

		group := service.NewServiceGroup()
		group.Add(zrpcServer)
		group.Add(gatewayServer)
		group.Add(customServer)

		logx.Infof("Starting rpc server at %s...", cc.MustGetConfig().Zrpc.ListenOn)
		logx.Infof("Starting gateway server at %s:%d...", cc.MustGetConfig().Gateway.Host, cc.MustGetConfig().Gateway.Port)
		group.Start()
	},
}

func printBanner(c config.BannerConf) {
	figure.NewColorFigure(c.Text, c.FontName, c.Color, true).Print()
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
