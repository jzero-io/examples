package main

import (
	"gateway-sqlx/tools/migrate/tables"
	"github.com/SpectatorNan/gorm-zero/gormc/config/mysql"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("etc/etc.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	conn, err := tables.NewGormConn(mysql.Mysql{
		Path:     viper.GetString("mysql.path"),
		Port:     viper.GetInt("mysql.port"),
		Dbname:   viper.GetString("mysql.dbname"),
		Username: viper.GetString("mysql.username"),
		Password: viper.GetString("mysql.password"),
	})
	if err != nil {
		panic(err)
	}
	err = tables.Migrate(conn)
	if err != nil {
		panic(err)
	}
}
