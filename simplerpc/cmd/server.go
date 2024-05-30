package cmd

import (
	"fmt"
	"simplerpc/internal/config"
	"simplerpc/internal/server"
	"simplerpc/internal/svc"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "simplerpc server",
	Long:  "simplerpc server",
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
	s := server.RegisterZrpc(ctx.Config, ctx)

	group := service.NewServiceGroup()
	group.Add(s)

	fmt.Printf("Starting rpc server at %s...\n", ctx.Config.Zrpc.ListenOn)
	group.Start()
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
