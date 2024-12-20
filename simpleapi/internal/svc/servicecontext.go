package svc

import (
	"simpleapi/internal/config"
	"simpleapi/internal/custom"
	"simpleapi/internal/middleware"
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
