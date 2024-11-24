package echo

import (
	"context"
	"simpleapi-serverless/server/svc"
	types "simpleapi-serverless/server/types/echo"

	"github.com/zeromicro/go-zero/core/logx"
)

type Echo struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEcho(ctx context.Context, svcCtx *svc.ServiceContext) *Echo {
	return &Echo{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Echo) Echo(req *types.EchoRequest) (resp *types.EchoResponse, err error) {

	return
}
