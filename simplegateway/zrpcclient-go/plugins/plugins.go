package plugins

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Plugins interface {
	Helloworld() Helloworld
}

type plugins struct {
	helloworld Helloworld
}

func (p *plugins) Helloworld() Helloworld {
	return p.helloworld
}

func NewPlugins(conn zrpc.Client) Plugins {
	return &plugins{
		helloworld: NewHelloworld(conn),
	}
}
