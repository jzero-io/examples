package cmd

import (
	"fmt"
	"simpleapi/internal/config"
	"simpleapi/internal/handler"
	"simpleapi/internal/svc"

	"github.com/jzero-io/jzero-contrib/swaggerv2"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "simpleapi server",
	Long:  "simpleapi server",
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
	server := rest.MustNewServer(ctx.Config.Rest.RestConf)

	handler.RegisterHandlers(server, ctx)

	swaggerv2.RegisterRoutes(server)

	group := service.NewServiceGroup()
	group.Add(server)

	fmt.Printf("Starting rest server at %s:%d...\n", ctx.Config.Rest.Host, ctx.Config.Rest.Port)
	group.Start()
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
