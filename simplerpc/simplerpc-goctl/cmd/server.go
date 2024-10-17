package cmd

import (
	"os"
	"simplerpc-goctl/internal/config"
	"simplerpc-goctl/internal/middleware"
	"simplerpc-goctl/internal/server"
	"simplerpc-goctl/internal/svc"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/proc"
	"github.com/zeromicro/go-zero/core/service"
	"golang.org/x/sync/errgroup"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "simplerpc-goctl server",
	Long:  "simplerpc-goctl server",
	Run: func(cmd *cobra.Command, args []string) {
		Start(cfgFile)
	},
}

func Start(cfgFile string) {
	var c config.Config
	conf.MustLoad(cfgFile, &c)
	config.C = c

	// set up logger
	if err := logx.SetUp(c.Log.LogConf); err != nil {
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
	middleware.Register(zrpc)

	group := service.NewServiceGroup()
	group.Add(zrpc)

	// shutdown listener
	waitExit := proc.AddShutdownListener(svcCtx.Custom.Stop)

	eg := errgroup.Group{}
	eg.Go(func() error {
		printBanner(svcCtx.Config)
		logx.Infof("Starting rpc server at %s...", svcCtx.Config.Zrpc.ListenOn)
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
