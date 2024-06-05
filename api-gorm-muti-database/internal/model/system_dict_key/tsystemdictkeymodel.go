package system_dict_key

import (
	"gorm.io/gorm"
)

var _ TSystemDictKeyModel = (*customTSystemDictKeyModel)(nil)

type (
	// TSystemDictKeyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTSystemDictKeyModel.
	TSystemDictKeyModel interface {
		tSystemDictKeyModel
		customTSystemDictKeyLogicModel
	}

	customTSystemDictKeyModel struct {
		*defaultTSystemDictKeyModel
	}

	customTSystemDictKeyLogicModel interface {
	}
)

// NewTSystemDictKeyModel returns a model for the database table.
func NewTSystemDictKeyModel(conn *gorm.DB) TSystemDictKeyModel {
	return &customTSystemDictKeyModel{
		defaultTSystemDictKeyModel: newTSystemDictKeyModel(conn),
	}
}
