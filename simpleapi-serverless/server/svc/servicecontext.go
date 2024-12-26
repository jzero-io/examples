package svc

import (
	"simpleapi-serverless/server/config"
	"simpleapi-serverless/server/middleware"
	"simpleapi-serverless/server/custom"

	configurator "github.com/zeromicro/go-zero/core/configcenter"
)

type ServiceContext struct {
	Config	configurator.Configurator[config.Config]
	middleware.Middleware
	Custom	*custom.Custom
}

func NewServiceContext(cc configurator.Configurator[config.Config]) *ServiceContext {
	sc := &ServiceContext{
		Config:		cc,
		Custom:		custom.New(),
		Middleware:	middleware.New(),
	}
	sc.SetConfigListener()
	return sc
}
