package svc

import (
	"quickstart/internal/config"
	"quickstart/internal/custom"
)

type ServiceContext struct {
	Config config.Config

	Custom *custom.Custom
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Custom: custom.New(),
	}
}
