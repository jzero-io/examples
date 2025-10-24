package svc

import (
	"github.com/jzero-io/jzero/core/configcenter"
	"github.com/jzero-io/jzero/core/stores/cache"
	"github.com/jzero-io/jzero/core/stores/modelx"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"simplegateway-with-model-redis/internal/config"
	"simplegateway-with-model-redis/internal/model"
)

type ServiceContext struct {
	ConfigCenter configcenter.ConfigCenter[config.Config]
	SqlxConn     sqlx.SqlConn
	Model        model.Model
	RedisConn    *redis.Redis
	Cache        cache.Cache
}

func NewServiceContext(cc configcenter.ConfigCenter[config.Config]) *ServiceContext {
	svcCtx := &ServiceContext{
		ConfigCenter: cc,
	}

	svcCtx.SqlxConn = modelx.MustNewConn(cc.MustGetConfig().Sqlx.SqlConf)
	svcCtx.RedisConn = redis.MustNewRedis(cc.MustGetConfig().Redis.RedisConf)
	svcCtx.Cache = cache.NewRedisNode(svcCtx.RedisConn, errors.New("cache not found"))
	svcCtx.Model = model.NewModel(svcCtx.SqlxConn, modelx.WithCachedConn(modelx.NewConnWithCache(svcCtx.SqlxConn, svcCtx.Cache)))
	svcCtx.SetConfigListener()
	return svcCtx
}
