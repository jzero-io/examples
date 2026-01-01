package plugins

import (
	"github.com/zeromicro/go-zero/zrpc"

	helloworldHelloworld "simplegateway/zrpcclient-go/plugins/helloworld/typed/helloworld"
)

type Helloworld interface {
	Helloworld() helloworldHelloworld.Helloworld
}

type helloworldClient struct {
	conn       zrpc.Client
	helloworld helloworldHelloworld.Helloworld
}

func (x *helloworldClient) Helloworld() helloworldHelloworld.Helloworld {
	return x.helloworld
}

func NewHelloworld(conn zrpc.Client) Helloworld {
	return &helloworldClient{
		conn:       conn,
		helloworld: helloworldHelloworld.NewHelloworld(conn),
	}
}
