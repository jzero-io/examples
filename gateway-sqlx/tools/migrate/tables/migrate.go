package tables

import (
	"github.com/SpectatorNan/gorm-zero/gormc/config/mysql"
	"gorm.io/gorm"
)

func NewGormConn(c mysql.Mysql) (*gorm.DB, error) {
	return mysql.Connect(c)
}

// Migrate https://gorm.io/docs/models.html
func Migrate(gormConn *gorm.DB) error {
	return gormConn.AutoMigrate(
		&User{},
	)
}
