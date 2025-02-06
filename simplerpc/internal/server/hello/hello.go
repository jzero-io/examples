// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.6
// Source: hello.proto

package server

import (
	"context"

	"simplerpc/internal/logic/hello"
	"simplerpc/internal/pb/hellopb"
	"simplerpc/internal/svc"
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

func (s *Hello) Say(ctx context.Context, in *hellopb.SayRequest) (*hellopb.SayResponse, error) {
	l := hellologic.NewSay(ctx, s.svcCtx)
	return l.Say(in)
}
