package svc

import (
	"github.com/jzero-io/jzero/core/configcenter"

	"simpleapi/internal/config"
)

type ServiceContext struct {
	ConfigCenter configcenter.ConfigCenter[config.Config]
	Middleware
}

func NewServiceContext(cc configcenter.ConfigCenter[config.Config]) *ServiceContext {
	svcCtx := &ServiceContext{
		ConfigCenter: cc,
	}

	svcCtx.SetConfigListener()
	svcCtx.Middleware = svcCtx.NewMiddleware()
	return svcCtx
}
