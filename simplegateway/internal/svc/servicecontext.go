package svc

import (
	configurator "github.com/zeromicro/go-zero/core/configcenter"

	"simplegateway/internal/config"
	"simplegateway/internal/custom"
)

type ServiceContext struct {
	Config configurator.Configurator[config.Config]
	Custom *custom.Custom
}

func NewServiceContext(cc configurator.Configurator[config.Config]) *ServiceContext {
	sc := &ServiceContext{
		Config: cc,
		Custom: custom.New(cc),
	}
	sc.SetConfigListener()
	return sc
}
