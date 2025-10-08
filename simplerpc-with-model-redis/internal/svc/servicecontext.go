package svc

import (
	"time"

	"github.com/jzero-io/jzero/core/configcenter"
	"github.com/jzero-io/jzero/core/stores/cache"
	"github.com/jzero-io/jzero/core/stores/modelx"
	"github.com/jzero-io/jzero/core/stores/redis"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"simplerpc-with-model-redis/internal/config"
	"simplerpc-with-model-redis/internal/model"
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
	svcCtx.Cache = cache.NewRedisNode(svcCtx.RedisConn, errors.New("cache not found"), cache.WithExpiry(time.Duration(5)*time.Second))
	svcCtx.Model = model.NewModel(svcCtx.SqlxConn, modelx.WithCachedConn(modelx.NewConnWithCache(svcCtx.SqlxConn, svcCtx.Cache)))
	svcCtx.SetConfigListener()
	return svcCtx
}
