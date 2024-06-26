// Code generated by jzero. DO NOT EDIT.
// type: fake_clientset

package fake

import (
	"simplegateway-go/typed"
	"simplegateway-go/typed/fake"

    simplegateway "simplegateway-go/typed/simplegateway"
    fakesimplegateway "simplegateway-go/typed/simplegateway/fake"

)

type Clientset struct {}

func (f *Clientset) Direct() typed.DirectInterface {
	return &fake.FakeDirect{}
}

func (f *Clientset) Simplegateway() simplegateway.SimplegatewayInterface {
    return &fakesimplegateway.FakeSimplegateway{}
}

