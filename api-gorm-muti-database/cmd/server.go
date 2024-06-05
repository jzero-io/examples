package cmd

import (
	"fmt"

	"api-gorm-muti-database/internal/config"
	"api-gorm-muti-database/internal/engine"
	"api-gorm-muti-database/internal/handler"
	"api-gorm-muti-database/internal/middlewares"
	"api-gorm-muti-database/internal/svc"
	"github.com/jzero-io/jzero-contrib/logtoconsole"
	"github.com/jzero-io/jzero-contrib/swaggerv2"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "api-gorm-muti-database server",
	Long:  "api-gorm-muti-database server",
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

	sqlxConn := sqlx.NewSqlConn(c.DataBaseType, engine.BuildDataSource(c))
	gormConn, err := engine.NewGormConn(c)
	if err != nil {
		panic(err)
	}
	err = engine.Migrate(gormConn)
	if err != nil {
		panic(err)
	}

	ctx := svc.NewServiceContext(c, sqlxConn, gormConn)
	start(ctx)
}

func start(ctx *svc.ServiceContext) {
	// print log to console if Log.Mode is file or volume
	if ctx.Config.Log.LogToConsole {
		logtoconsole.MustLogToConsole(ctx.Config.Log.LogConf)
	}

	// verify sqlx conn
	_, err := ctx.SqlxConn.Exec("select 1")
	if err != nil {
		panic(err)
	}

	// verify gorm conn
	err = ctx.GormConn.Exec("select 1").Error
	if err != nil {
		panic(err)
	}

	server := rest.MustNewServer(ctx.Config.Rest.RestConf, rest.WithCors("*"))

	httpx.SetOkHandler(middlewares.OKHandler)
	httpx.SetErrorHandler(middlewares.ErrorHandler)

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
