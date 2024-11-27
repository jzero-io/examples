package svc

import (
	"simpleapi-serverless/server/config"
	"simpleapi-serverless/server/custom"
	"simpleapi-serverless/server/middleware"
)

type ServiceContext struct {
	Config config.Config
	Custom *custom.Custom

	middleware.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Custom:     custom.New(),
		Middleware: middleware.NewMiddleware(),
	}
}
