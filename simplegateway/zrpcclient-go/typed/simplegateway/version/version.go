// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.3
// Source: version.proto

package version

import (
	"context"

	"simplegateway/zrpcclient-go/model/simplegateway/pb/versionpb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetRequest  = versionpb.GetRequest
	GetResponse = versionpb.GetResponse

	Version interface {
		Say(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	}

	defaultVersion struct {
		cli zrpc.Client
	}
)

func NewVersion(cli zrpc.Client) Version {
	return &defaultVersion{
		cli: cli,
	}
}

func (m *defaultVersion) Say(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	client := versionpb.NewVersionClient(m.cli.Conn())
	return client.Say(ctx, in, opts...)
}
