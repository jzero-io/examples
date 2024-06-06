package dict

import "gorm.io/gorm"

type TSystemDictKey struct {
	gorm.Model
	label string
	value string
}

func (TSystemDictKey) TableName() string {
	return "T_system_dict_key"
}
