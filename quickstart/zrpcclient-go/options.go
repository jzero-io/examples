package zrpcclient_go

import (
	"github.com/zeromicro/go-zero/zrpc"

	"quickstart/zrpcclient-go/typed/quickstart"
)

type Opt func(client *Clientset)


func WithQuickstartClient(cli zrpc.Client) Opt {
	return func(client *Clientset) {
		client.quickstart = quickstart.New(cli)
	}
}


