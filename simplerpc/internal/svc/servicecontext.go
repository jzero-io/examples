package svc

import (
	"simplerpc/internal/config"
	"simplerpc/internal/custom"

	configurator "github.com/zeromicro/go-zero/core/configcenter"
)

type ServiceContext struct {
	Config configurator.Configurator[config.Config]

	Custom *custom.Custom
}

func NewServiceContext(cc configurator.Configurator[config.Config]) *ServiceContext {
	sc := &ServiceContext{
		Config: cc,
		Custom: custom.New(),
	}
	sc.SetConfigListener()
	return sc
}
