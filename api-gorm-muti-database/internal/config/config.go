package config

import (
	"github.com/SpectatorNan/gorm-zero/gormc/config/mysql"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	Rest RestConf
	Log  LogConf

	DataBaseType string
	Mysql        mysql.Mysql
	Dm           DmConf
}

type RestConf struct {
	rest.RestConf
}

type LogConf struct {
	logx.LogConf
	// only Log.Mode is file or volume take effect
	LogToConsole bool `json:",default=true"`
}

type DmConf struct {
	Username string `json:",optional"`
	Password string `json:",optional"`
	Path     string `json:",optional"`
	Port     int    `json:",optional"`
	Database string `json:",optional"`
}
