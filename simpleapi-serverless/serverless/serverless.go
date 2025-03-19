package serverless

import (
	"path/filepath"
	"simpleapi-serverless/internal/config"
	"simpleapi-serverless/internal/handler"
	"simpleapi-serverless/internal/svc"

	"github.com/jzero-io/jzero-contrib/dynamic_conf"
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

type Serverless struct {
	SvcCtx      *svc.ServiceContext                                   // 服务上下文
	HandlerFunc func(server *rest.Server, svcCtx *svc.ServiceContext) // 服务路由
}

// Serverless please replace coreSvcCtx any type to real CoreSvcCtx
func New(coreSvcCtx any) *Serverless {
	ss, err := dynamic_conf.NewFsNotify(filepath.Join("plugins", "simpleapi-serverless", "etc", "etc.yaml"), dynamic_conf.WithUseEnv(true))
	logx.Must(err)
	cc := configurator.MustNewConfigCenter[config.Config](configurator.Config{
		Type: "yaml",
	}, ss)

	svcCtx := svc.NewServiceContext(cc)
	return &Serverless{
		SvcCtx:      svcCtx,
		HandlerFunc: handler.RegisterHandlers,
	}
}
