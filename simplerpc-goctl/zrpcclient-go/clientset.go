package zrpcclient_go

import (
	"simplerpc-goctl/zrpcclient-go/typed/simplerpc_goctl"
)

type Interface interface {
	SimplerpcGoctl() simplerpc_goctl.Interface
}

type Clientset struct {
	simplerpcGoctl *simplerpc_goctl.Client
}

func (x *Clientset) SimplerpcGoctl() simplerpc_goctl.Interface {
	return x.simplerpcGoctl
}

func NewClientWithOptions(ops ...Opt) *Clientset {
	cs := &Clientset{}

	for _, op := range ops {
		op(cs)
	}

	return cs
}
