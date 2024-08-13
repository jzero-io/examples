package tables

type User struct {
	Id         uint   `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Username   string `gorm:"column:username;unique;not null"`
	Password   string `gorm:"column:password;not null"`
	CreateTime string `gorm:"column:create_time;not null"`
}

func (User) TableName() string {
	return "users"
}
