// Code generated by jzero. DO NOT EDIT.
// type: clientset

package simpleapi_goctl_go

import (
	"github.com/jzero-io/restc"
	"simpleapi-goctl/simpleapi-goctl-go/typed"

	"simpleapi-goctl/simpleapi-goctl-go/typed/simpleapi_goctl"
)

type Interface interface {
	Direct() typed.DirectInterface

	SimpleapiGoctl() simpleapi_goctl.SimpleapiGoctlInterface
}

type Clientset struct {
	// direct client to request
	direct *typed.DirectClient

	simpleapiGoctl *simpleapi_goctl.SimpleapiGoctlClient
}

func (x *Clientset) Direct() typed.DirectInterface {
	return x.direct
}

func (x *Clientset) SimpleapiGoctl() simpleapi_goctl.SimpleapiGoctlInterface {
	return x.simpleapiGoctl
}

func NewClientWithOptions(ops ...restc.Opt) (simpleapi_goctl.SimpleapiGoctlInterface, error) {
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
	cs.simpleapiGoctl, err = simpleapi_goctl.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	
	return cs.SimpleapiGoctl(), nil
}