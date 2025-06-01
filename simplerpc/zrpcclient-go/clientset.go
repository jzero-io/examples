package zrpcclient_go

import (
	"simplerpc/zrpcclient-go/typed/simplerpc"
)

type Interface interface {
	Simplerpc() simplerpc.Interface
}

type Clientset struct {
	simplerpc *simplerpc.Client
}

func (x *Clientset) Simplerpc() simplerpc.Interface {
	return x.simplerpc
}

func NewClientWithOptions(ops ...Opt) *Clientset {
	cs := &Clientset{}

	for _, op := range ops {
		op(cs)
	}

	return cs
}
