package engine

import (
	"api-gorm-muti-database/internal/config"
	"fmt"
	"github.com/SpectatorNan/gorm-zero/gormc/config/mysql"
	dameng "github.com/godoes/gorm-dameng"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

func NewGormConn(c config.Config) (*gorm.DB, error) {
	switch c.DataBaseType {
	case "mysql":
		return mysql.Connect(c.Mysql)
	case "dm":
		options := map[string]string{
			"schema":         "SYSDBA",
			"compatibleMode": "Mysql",
		}
		dsn := dameng.BuildUrl(c.Dm.Username, c.Dm.Password, c.Dm.Path, c.Dm.Port, options)
		return gorm.Open(dameng.Open(dsn), &gorm.Config{})
	}
	return nil, fmt.Errorf("unsupported database type: %s", c.DataBaseType)
}

func BuildDataSource(c config.Config) string {
	switch c.DataBaseType {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			c.Mysql.Username,
			c.Mysql.Password,
			c.Mysql.Path+":"+cast.ToString(c.Mysql.Port),
			c.Mysql.Dbname)
	case "dm":
		return fmt.Sprintf("dm://%s:%s@%s?compatibleMode=mysql&columnNameCase=lower",
			c.Dm.Username,
			c.Dm.Password,
			c.Dm.Path+":"+cast.ToString(c.Dm.Port),
		)
	}
	return ""
}
