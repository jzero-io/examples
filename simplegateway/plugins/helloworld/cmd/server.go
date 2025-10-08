package cmd

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/jzero-io/jzero/core/configcenter"
	"github.com/jzero-io/jzero/core/configcenter/subscriber"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/gateway"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"helloworld/desc/pb"
	"helloworld/internal/config"
	"helloworld/internal/custom"
	"helloworld/internal/global"
	"helloworld/internal/middleware"
	"helloworld/internal/server"
	"helloworld/internal/svc"
)

// serverCmd represents the server command
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

		svcCtx := svc.NewServiceContext(cc)
		global.ServiceContext = *svcCtx

		var err error
		cc.MustGetConfig().Gateway.Upstreams[0].ProtoSets, err = pb.WriteToLocal(pb.Embed)
		logx.Must(err)

		zrpcServer := zrpc.MustNewServer(cc.MustGetConfig().Zrpc.RpcServerConf, func(grpcServer *grpc.Server) {
			server.RegisterZrpcServer(grpcServer, svcCtx)

			if cc.MustGetConfig().Zrpc.Mode == service.DevMode || cc.MustGetConfig().Zrpc.Mode == service.TestMode {
				reflection.Register(grpcServer)
			}
		})
		gatewayServer := gateway.MustNewServer(cc.MustGetConfig().Gateway.GatewayConf, middleware.WithHeaderProcessor())

		ctm := custom.New(zrpcServer, gatewayServer)
		ctm.Init()
		// register middleware
		middleware.Register(zrpcServer, gatewayServer)
		group := service.NewServiceGroup()
		group.Add(zrpcServer)
		group.Add(gatewayServer)
		group.Add(ctm)
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
