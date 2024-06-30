package hellov2logic

import (
	"context"
	"simplegateway/internal/pb/hellov2pb"
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

func (l *SayHello) SayHello(in *hellov2pb.Empty) (*hellov2pb.SayHelloResponse, error) {
	return &hellov2pb.SayHelloResponse{}, nil
}
