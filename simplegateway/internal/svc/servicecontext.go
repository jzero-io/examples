package svc

import (
	"simplegateway/internal/config"
	"simplegateway/internal/custom"

	configurator "github.com/zeromicro/go-zero/core/configcenter"
)

type ServiceContext struct {
	Config configurator.Configurator[config.Config]

	Custom *custom.Custom
}

func NewServiceContext(c config.Config, cc configurator.Configurator[config.Config]) *ServiceContext {
	sc := &ServiceContext{
		Config: cc,
		Custom: custom.New(),
	}
	sc.DynamicConfListener(c, cc)
	return sc
}
