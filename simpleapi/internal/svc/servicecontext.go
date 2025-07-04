package svc

import (
	configurator "github.com/zeromicro/go-zero/core/configcenter"

	"simpleapi/internal/config"
	"simpleapi/internal/custom"
	"simpleapi/internal/middleware"
)

type ServiceContext struct {
	Config configurator.Configurator[config.Config]
	middleware.Middleware
	Custom *custom.Custom
}

func NewServiceContext(cc configurator.Configurator[config.Config]) *ServiceContext {
	sc := &ServiceContext{
		Config:     cc,
		Custom:     custom.New(),
		Middleware: middleware.New(),
	}
	sc.SetConfigListener()
	return sc
}
