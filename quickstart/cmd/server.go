package cmd

import (
	"fmt"
	"quickstart/internal/config"
	"quickstart/internal/handler"
	"quickstart/internal/server"
	"quickstart/internal/svc"

	"github.com/jzero-io/jzero-contrib/swaggerv2"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/gateway"
)

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

	if err := logx.SetUp(c.Log.LogConf); err != nil {
		logx.Must(err)
	}

	ctx := svc.NewServiceContext(c)
	start(ctx)
}

func start(ctx *svc.ServiceContext) {
	s := server.GetZrpcServer(ctx.Config, ctx)
	gw := gateway.MustNewServer(ctx.Config.Gateway.GatewayConf)

	handler.RegisterHandlers(gw.Server, ctx)

	swaggerv2.RegisterRoutes(gw.Server)

	group := service.NewServiceGroup()
	group.Add(s)
	group.Add(gw)

	fmt.Printf("Starting rpc server at %s...\n", ctx.Config.Zrpc.ListenOn)
	fmt.Printf("Starting gateway server at %s:%d...\n", ctx.Config.Gateway.Host, ctx.Config.Gateway.Port)
	group.Start()

}

func init() {
	rootCmd.AddCommand(serverCmd)
}
