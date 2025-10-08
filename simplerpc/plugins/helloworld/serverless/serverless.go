package serverless

import (
	"path/filepath"

	"github.com/jzero-io/jzero/core/configcenter"
	"github.com/jzero-io/jzero/core/configcenter/subscriber"
	"google.golang.org/grpc"

	"helloworld/internal/config"
	"helloworld/internal/global"
	"helloworld/internal/server"
	"helloworld/internal/svc"
)

type Serverless struct {
	SvcCtx             *svc.ServiceContext // 服务上下文
	RegisterZrpcServer func(grpcServer *grpc.Server, ctx *svc.ServiceContext)
}

func New() *Serverless {
	cc := configcenter.MustNewConfigCenter[config.Config](configcenter.Config{
		Type: "yaml",
	}, subscriber.MustNewFsnotifySubscriber(filepath.Join("plugins", "helloworld", "etc", "etc.yaml"), subscriber.WithUseEnv(true)))

	svcCtx := svc.NewServiceContext(cc)
	global.ServiceContext = *svcCtx

	return &Serverless{
		SvcCtx:             svcCtx,
		RegisterZrpcServer: server.RegisterZrpcServer,
	}
}
