package model

const TAB_NAME_USER = "user_tab"

type User struct {
	Id       uint64 `gorm:"column:id"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	Status   uint8  `gorm:"column:status"`
}

func (c User) TableName() string {
	return TAB_NAME_USER
}
