package svc

import (
	configurator "github.com/zeromicro/go-zero/core/configcenter"

	"simplerpc/internal/config"
	"simplerpc/internal/custom"
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
