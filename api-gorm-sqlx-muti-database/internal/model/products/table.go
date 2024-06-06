package products

import "gorm.io/gorm"

type GormProduct struct {
	gorm.Model
	Code  string
	Price uint

	Remark string `gorm:"column:remark;size:0"`
}

func (g GormProduct) TableName() string {
	return "products"
}
