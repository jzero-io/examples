package simplerpc

import (
	"github.com/zeromicro/go-zero/zrpc"
	
	"simplerpc/zrpcclient-go/typed/simplerpc/hello"
	
)

type Interface interface {
	
	Hello() hello.Hello
	
}

type Client struct {
	client zrpc.Client
}

func New(c zrpc.Client) *Client {
	return &Client{client: c}
}


func (x *Client) Hello() hello.Hello {
	return hello.NewHello(x.client)
}
