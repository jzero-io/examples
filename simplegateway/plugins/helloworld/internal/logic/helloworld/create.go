package helloworldlogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"simplegateway/plugins/helloworld/internal/svc"
	"simplegateway/plugins/helloworld/internal/types/helloworld"
)

type Create struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreate(ctx context.Context, svcCtx *svc.ServiceContext) *Create {
	return &Create{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Create) Create(in *helloworld.CreateRequest) (*helloworld.CreateResponse, error) {
	// todo: add your logic here and delete this line

	return &helloworld.CreateResponse{}, nil
}
