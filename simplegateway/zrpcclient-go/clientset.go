package zrpcclient_go

import (
	"simplegateway/zrpcclient-go/typed/simplegateway"
)

type Interface interface {
	Simplegateway() simplegateway.Interface
}

type Clientset struct {
	simplegateway *simplegateway.Client
}

func (x *Clientset) Simplegateway() simplegateway.Interface {
	return x.simplegateway
}

func NewClientWithOptions(ops ...Opt) *Clientset {
	cs := &Clientset{}

	for _, op := range ops {
		op(cs)
	}

	return cs
}
