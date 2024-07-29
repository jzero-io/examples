package zrpcclient_go

import (
	"github.com/zeromicro/go-zero/zrpc"

	"simplegateway/zrpcclient-go/typed/simplegateway"
)

type Opt func(client *Clientset)


func WithSimplegatewayClient(cli zrpc.Client) Opt {
	return func(client *Clientset) {
		client.simplegateway = simplegateway.New(cli)
	}
}


