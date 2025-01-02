// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.4
// Source: hello.proto

package server

import (
	"context"

	"simplerpc-goctl/internal/logic/hello"
	"simplerpc-goctl/internal/pb/hellopb"
	"simplerpc-goctl/internal/svc"
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
