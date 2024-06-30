package hellologic

import (
	"context"

	"simplegateway/internal/pb/hellopb"
	"simplegateway/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SayHello struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSayHello(ctx context.Context, svcCtx *svc.ServiceContext) *SayHello {
	return &SayHello{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SayHello) SayHello(in *hellopb.Empty) (*hellopb.SayHelloResponse, error) {
	// todo: add your logic here and delete this line

	return &hellopb.SayHelloResponse{}, nil
}
