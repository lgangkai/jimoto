package model

const TAB_NAME_ORDER = "order_tab"

type Order struct {
	Id uint64 `gorm:"column:id"`
}

func (o Order) TableName() string {
	return TAB_NAME_ORDER
}
