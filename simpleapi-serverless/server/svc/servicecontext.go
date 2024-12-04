package svc

import (
	"simpleapi-serverless/server/config"
	"simpleapi-serverless/server/custom"
	"simpleapi-serverless/server/middleware"
)

type ServiceContext struct {
	Config config.Config
	middleware.Middleware
	Custom *custom.Custom
}

func NewServiceContext(c config.Config) *ServiceContext {
	sc := &ServiceContext{
		Config:     c,
		Custom:     custom.New(),
		Middleware: middleware.New(),
	}
	return sc
}
