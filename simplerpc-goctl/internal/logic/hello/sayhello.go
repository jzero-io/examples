package hellologic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"simplerpc-goctl/internal/pb/hellopb"
	"simplerpc-goctl/internal/svc"
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

func (l *SayHello) SayHello(in *hellopb.SayHelloRequest) (*hellopb.SayHelloResponse, error) {
	// todo: add your logic here and delete this line

	return &hellopb.SayHelloResponse{}, nil
}
