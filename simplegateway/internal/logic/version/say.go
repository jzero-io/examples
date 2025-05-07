package versionlogic

import (
	"context"

	"simplegateway/internal/pb/versionpb"
	"simplegateway/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type Say struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSay(ctx context.Context, svcCtx *svc.ServiceContext) *Say {
	return &Say{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Say) Say(in *versionpb.GetRequest) (*versionpb.GetResponse, error) {
	// todo: add your logic here and delete this line

	return &versionpb.GetResponse{}, nil
}
