package zrpcclient_go

import (
	"github.com/zeromicro/go-zero/zrpc"

	"simplerpc/zrpcclient-go/typed/simplerpc"
)

type Opt func(client *Clientset)

func WithSimplerpcClient(cli zrpc.Client) Opt {
	return func(client *Clientset) {
		client.simplerpc = simplerpc.New(cli)
	}
}
