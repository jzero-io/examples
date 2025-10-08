package serverless

import (
	"path/filepath"

	"github.com/jzero-io/jzero/core/configcenter"
	"github.com/jzero-io/jzero/core/configcenter/subscriber"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"

	"helloworld/desc/pb"
	"helloworld/internal/config"
	"helloworld/internal/server"
	"helloworld/internal/svc"
)

type Serverless struct {
	SvcCtx             *svc.ServiceContext // 服务上下文
	RegisterZrpcServer func(grpcServer *grpc.Server, ctx *svc.ServiceContext)
	ProtoSets          []string
}

func New() *Serverless {
	cc := configcenter.MustNewConfigCenter[config.Config](configcenter.Config{
		Type: "yaml",
	}, subscriber.MustNewFsnotifySubscriber(filepath.Join("plugins", "helloworld", "etc", "etc.yaml"), subscriber.WithUseEnv(true)))

	svcCtx := svc.NewServiceContext(cc)

	// get protoSets
	protoSets, err := pb.WriteToLocal(pb.Embed)
	logx.Must(err)

	return &Serverless{
		SvcCtx:             svcCtx,
		RegisterZrpcServer: server.RegisterZrpcServer,
		ProtoSets:          protoSets,
	}
}
