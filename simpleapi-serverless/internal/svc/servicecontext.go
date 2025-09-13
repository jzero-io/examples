package svc

import (
	configurator "github.com/zeromicro/go-zero/core/configcenter"

	"simpleapi-serverless/internal/config"
)

type ServiceContext struct {
	Config configurator.Configurator[config.Config]
	Middleware
}

func NewServiceContext(cc configurator.Configurator[config.Config]) *ServiceContext {
	sc := &ServiceContext{
		Config: cc,
	}

	sc.SetConfigListener()
	return sc
}
