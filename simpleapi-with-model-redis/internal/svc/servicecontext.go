package svc

import (
	"time"

	"github.com/jzero-io/jzero/core/stores/cache"
	"github.com/jzero-io/jzero/core/stores/modelx"
	"github.com/jzero-io/jzero/core/stores/redis"
	"github.com/pkg/errors"
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"simpleapi-with-model-redis/internal/config"
	"simpleapi-with-model-redis/internal/model"
)

type ServiceContext struct {
	Config configurator.Configurator[config.Config]
	Middleware
	SqlxConn  sqlx.SqlConn
	Model     model.Model
	RedisConn *redis.Redis
	Cache     cache.Cache
}

func NewServiceContext(cc configurator.Configurator[config.Config]) *ServiceContext {
	sc := &ServiceContext{
		Config: cc,
	}

	sc.SqlxConn = modelx.MustNewConn(sc.MustGetConfig().Sqlx.SqlConf)
	sc.RedisConn = redis.MustNewRedis(sc.MustGetConfig().Redis.RedisConf)
	sc.Cache = cache.NewRedisNode(sc.RedisConn, errors.New("cache not found"), cache.WithExpiry(time.Duration(5)*time.Second))
	sc.Model = model.NewModel(sc.SqlxConn, modelx.WithCachedConn(modelx.NewConnWithCache(sc.SqlxConn, sc.Cache)))
	sc.SetConfigListener()
	return sc
}
