package simplegateway

import (
	"github.com/zeromicro/go-zero/zrpc"
	
	"simplegateway/zrpcclient-go/typed/simplegateway/version"
	
)

type Interface interface {
	
	Version() version.Version
	
}

type Client struct {
	client zrpc.Client
}

func New(c zrpc.Client) *Client {
	return &Client{client: c}
}


func (x *Client) Version() version.Version {
	return version.NewVersion(x.client)
}
