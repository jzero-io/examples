package svc

import (
	"simplerpc-goctl/internal/config"
	"simplerpc-goctl/internal/custom"
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
