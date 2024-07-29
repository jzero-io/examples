// Code generated by jzero. DO NOT EDIT.
// type: clientset

package quickstart_go

import (
	"github.com/jzero-io/restc"
	"quickstart/quickstart-go/typed"

	"quickstart/quickstart-go/typed/quickstart"
)

type Interface interface {
	Direct() typed.DirectInterface

	Quickstart() quickstart.QuickstartInterface
}

type Clientset struct {
	// direct client to request
	direct *typed.DirectClient

	quickstart *quickstart.QuickstartClient
}

func (x *Clientset) Direct() typed.DirectInterface {
	return x.direct
}

func (x *Clientset) Quickstart() quickstart.QuickstartInterface {
	return x.quickstart
}

func NewClientWithOptions(ops ...restc.Opt) (quickstart.QuickstartInterface, error) {
	c := &restc.RESTClient{}
	for _, op := range ops {
		if err := op(c); err != nil {
			return nil, err
		}
	}
	configShallowCopy := *c
	var cs Clientset
	var err error
	cs.direct, err = typed.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.quickstart, err = quickstart.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	
	return cs.Quickstart(), nil
}