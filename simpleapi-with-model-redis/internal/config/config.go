package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	Rest  RestConf
	Log   LogConf
	Sqlx  SqlxConf
	Redis RedisConf

	Banner BannerConf
}

type RestConf struct {
	rest.RestConf
}

type LogConf struct {
	logx.LogConf
}

type SqlxConf struct {
	sqlx.SqlConf
}

type RedisConf struct {
	redis.RedisConf
}

type BannerConf struct {
	Text     string `json:",default=JZERO"`
	Color    string `json:",default=green"`
	FontName string `json:",default=starwars,options=big|larry3d|starwars|standard"`
}
