package model

const TAB_NAME_COMMODITY = "commodity_tab"

type Commodity struct {
	Id        uint64 `gorm:"column:id"`
	CreatorId uint64 `gorm:"column:creator_id"`
	Title     string `gorm:"column:title"`
	Detail    string `gorm:"column:detail"`
	Price     uint64 `gorm:"column:price"`
	Cover     string `gorm:"column:cover"`
	Type      uint32 `gorm:"column:type"`
	Status    uint32 `gorm:"column:status"`
	IsDeleted bool   `gorm:"column:is_deleted"`
}

func (c Commodity) TableName() string {
	return TAB_NAME_COMMODITY
}
