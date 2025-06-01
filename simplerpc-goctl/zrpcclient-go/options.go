package zrpcclient_go

import (
	"github.com/zeromicro/go-zero/zrpc"

	"simplerpc-goctl/zrpcclient-go/typed/simplerpc_goctl"
)

type Opt func(client *Clientset)

func WithSimplerpcGoctlClient(cli zrpc.Client) Opt {
	return func(client *Clientset) {
		client.simplerpcGoctl = simplerpc_goctl.New(cli)
	}
}
