package helloworld

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"

	"helloworld/internal/svc"
	types "helloworld/internal/types/helloworld"
)

type Create struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewCreate(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *Create {
	return &Create{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *Create) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
