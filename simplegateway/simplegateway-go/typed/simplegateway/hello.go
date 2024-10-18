// Code generated by jzero. DO NOT EDIT.
// type: Hello

package simplegateway

import (
	"context"
	"simplegateway/simplegateway-go/model/simplegateway/pb/hellopb"

	"github.com/jzero-io/restc"
)

var (
	_ = context.Background()
)

type HelloGetter interface {
	Hello() HelloInterface
}

type HelloInterface interface {
	// API /api/v1/hello/say
	Say(ctx context.Context, param *hellopb.SayRequest) (*hellopb.SayResponse, error)

	HelloExpansion
}

type HelloClient struct {
	client restc.Interface
}

func newHelloClient(c *SimplegatewayClient) *HelloClient {
	return &HelloClient{
		client: c.RESTClient(),
	}
}

func (x *HelloClient) Say(ctx context.Context, param *hellopb.SayRequest) (*hellopb.SayResponse, error) {
	var resp *hellopb.SayResponse
	err := x.client.Verb("GET").
		SubPath(
			"/api/v1/hello/say",
		).
		Params(
			restc.QueryParam{Name: "message", Value: param.Message},
		).
		Do(ctx).
		Into(&resp, true)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
