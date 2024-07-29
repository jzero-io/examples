package zrpcclient_go

import (
	"quickstart/zrpcclient-go/typed/quickstart"
)

type Interface interface {
	Quickstart() quickstart.Interface
}

type Clientset struct {
	quickstart *quickstart.Client
}

func (x *Clientset) Quickstart() quickstart.Interface {
	return x.quickstart
}


func NewClientWithOptions(ops ...Opt) *Clientset {
	cs := &Clientset{}

	for _, op := range ops {
		op(cs)
	}

	return cs
}
