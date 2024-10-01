package model

const TAB_NAME_LIKE = "like_tab"

type Like struct {
	Id          uint64 `gorm:"column:id"`
	CommodityId uint64 `gorm:"column:commodity_id"`
	UserId      uint64 `gorm:"column:user_id"`
}

func (c Like) TableName() string {
	return TAB_NAME_LIKE
}
