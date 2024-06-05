package svc

import (
	"api-gorm-muti-database/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config   config.Config
	SqlxConn sqlx.SqlConn
	GormConn *gorm.DB
}

func NewServiceContext(c config.Config, sqlxConn sqlx.SqlConn, gormConn *gorm.DB) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		SqlxConn: sqlxConn,
		GormConn: gormConn,
	}
}
