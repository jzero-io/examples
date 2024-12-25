package svc

import (
	"simpleapi-serverless/server/config"
	"simpleapi-serverless/server/middleware"
	"simpleapi-serverless/server/custom"
)

type ServiceContext struct {
	Config	config.Config
	middleware.Middleware
	Custom	*custom.Custom
}

func NewServiceContext(c config.Config) *ServiceContext {
	sc := &ServiceContext{
		Config:		c,
		Custom:		custom.New(),
		Middleware:	middleware.New(),
	}
	return sc
}

func (sc *ServiceContext) MustGetConfig() config.Config {
	return sc.Config
}
