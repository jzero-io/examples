package system_dict_value

import (
	"gorm.io/gorm"
)

var _ TSystemDictValueModel = (*customTSystemDictValueModel)(nil)

type (
	// TSystemDictValueModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTSystemDictValueModel.
	TSystemDictValueModel interface {
		tSystemDictValueModel
		customTSystemDictValueLogicModel
	}

	customTSystemDictValueModel struct {
		*defaultTSystemDictValueModel
	}

	customTSystemDictValueLogicModel interface {
	}
)

// NewTSystemDictValueModel returns a model for the database table.
func NewTSystemDictValueModel(conn *gorm.DB) TSystemDictValueModel {
	return &customTSystemDictValueModel{
		defaultTSystemDictValueModel: newTSystemDictValueModel(conn),
	}
}
