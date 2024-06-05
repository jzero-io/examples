package engine

import (
	"api-gorm-muti-database/internal/model/system_dict_key"
	"api-gorm-muti-database/internal/model/system_dict_value"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&system_dict_key.TSystemDictKey{},
		&system_dict_value.TSystemDictValue{})
}
