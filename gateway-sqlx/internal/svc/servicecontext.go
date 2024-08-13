package svc

import (
	"gateway-sqlx/internal/config"
	"gateway-sqlx/internal/custom"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	SqlxConn sqlx.SqlConn
	Model    Model

	Custom *custom.Custom
}

func NewServiceContext(c config.Config) *ServiceContext {
	svcCtx := &ServiceContext{
		Config:   c,
		SqlxConn: MustSqlConn(c),
		Custom:   custom.New(),
	}

	svcCtx.Model = NewModel(svcCtx.SqlxConn)
	return svcCtx
}
