package model

const TAB_NAME_COMMODITY_IMAGE = "commodity_image_tab"

type CommodityImage struct {
	Id          uint64 `gorm:"column:id"`
	CommodityId uint64 `gorm:"column:commodity_id"`
	Image       string `gorm:"column:image"`
}

func (c CommodityImage) TableName() string {
	return TAB_NAME_COMMODITY_IMAGE
}
