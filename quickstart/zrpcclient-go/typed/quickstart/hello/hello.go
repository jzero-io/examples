// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.1
// Source: hello.proto

package hello

import (
	"context"

	"quickstart/zrpcclient-go/model/quickstart/pb/hellopb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	SayHelloRequest  = hellopb.SayHelloRequest
	SayHelloResponse = hellopb.SayHelloResponse

	Hello interface {
		SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error)
	}

	defaultHello struct {
		cli zrpc.Client
	}
)

func NewHello(cli zrpc.Client) Hello {
	return &defaultHello{
		cli: cli,
	}
}

func (m *defaultHello) SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error) {
	client := hellopb.NewHelloClient(m.cli.Conn())
	return client.SayHello(ctx, in, opts...)
}
