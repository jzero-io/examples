package logic

import (
	"net/http"

	"context"

	"simpleapi-serverless/internal/svc"
	"simpleapi-serverless/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Get struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGet(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *Get {
	return &Get{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *Get) Get(req *types.GetRequest) (resp *types.GetResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
