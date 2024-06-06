package model

import (
	"api-gorm-sqlx-muti-database/internal/model/dict"
	"api-gorm-sqlx-muti-database/internal/model/products"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&products.GormProduct{}, &dict.TSystemDictKey{})
}
