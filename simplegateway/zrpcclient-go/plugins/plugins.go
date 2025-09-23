package plugins

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Plugins any

type plugins struct {
}

func NewPlugins(conn zrpc.Client) Plugins {
	return &plugins{}
}
