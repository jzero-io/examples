package version

import (
	"context"

	"simpleapi-serverless/server/svc"
	types "simpleapi-serverless/server/types/version"

	"github.com/zeromicro/go-zero/core/logx"
)

type Get struct {
	logx.Logger
	ctx	context.Context
	svcCtx	*svc.ServiceContext
}

func NewGet(ctx context.Context, svcCtx *svc.ServiceContext) *Get {
	return &Get{
		Logger:	logx.WithContext(ctx),
		ctx:	ctx,
		svcCtx:	svcCtx,
	}
}

func (l *Get) Get(req *types.GetRequest) (resp *types.GetResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
