// Code generated by goctl. DO NOT EDIT.
// Source: hello.proto

package server

import (
	"context"

	"simplegateway/internal/logic/hello"
	"simplegateway/internal/pb/hellopb"
	"simplegateway/internal/svc"
)

type Hello struct {
	svcCtx *svc.ServiceContext
	hellopb.UnimplementedHelloServer
}

func NewHello(svcCtx *svc.ServiceContext) *Hello {
	return &Hello{
		svcCtx: svcCtx,
	}
}

func (s *Hello) SayHello(ctx context.Context, in *hellopb.SayHelloRequest) (*hellopb.SayHelloResponse, error) {
	l := hellologic.NewSayHello(ctx, s.svcCtx)
	return l.SayHello(in)
}
