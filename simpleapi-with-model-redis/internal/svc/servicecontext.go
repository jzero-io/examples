package svc

import (
	"github.com/jzero-io/jzero/core/configcenter"
	"github.com/jzero-io/jzero/core/stores/modelx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"simpleapi-with-model-redis/internal/config"
	"simpleapi-with-model-redis/internal/model"
)

type ServiceContext struct {
	ConfigCenter configcenter.ConfigCenter[config.Config]
	Middleware
	SqlxConn  sqlx.SqlConn
	Model     model.Model
	RedisConn *redis.Redis
}

func NewServiceContext(cc configcenter.ConfigCenter[config.Config]) *ServiceContext {
	svcCtx := &ServiceContext{
		ConfigCenter: cc,
	}

	svcCtx.SqlxConn = modelx.MustNewConn(cc.MustGetConfig().Sqlx.SqlConf)
	svcCtx.RedisConn = redis.MustNewRedis(cc.MustGetConfig().Redis.RedisConf)

	svcCtx.Model = model.NewModel(svcCtx.SqlxConn)
	svcCtx.SetConfigListener()
	svcCtx.Middleware = svcCtx.NewMiddleware()
	return svcCtx
}
