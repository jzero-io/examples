package svc

import (
	"simpleapi-serverless/internal/config"
	"simpleapi-serverless/internal/custom"
	"simpleapi-serverless/internal/middleware"

	configurator "github.com/zeromicro/go-zero/core/configcenter"
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
