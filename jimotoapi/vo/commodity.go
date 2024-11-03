package vo

import (
	"jimotoapi/conf"
	"jimotoapi/util"
	"protos/commodity"
)

type Commodity struct {
	Id     uint64 `json:"id"`
	Title  string `json:"title"`
	Price  uint64 `json:"price"`
	Cover  string `json:"cover"`
	Type   uint32 `json:"type"`
	Status uint32 `json:"status"`
}

type CommodityList []*Commodity

func FromCommodity(c *commodity.CommodityItem, config *conf.Config) *Commodity {
	return &Commodity{
		Id:     c.Id,
		Title:  c.Title,
		Price:  c.Price,
		Cover:  util.CompleteImageUrl(c.Cover, config),
		Type:   c.Type,
		Status: c.Status,
	}
}

func FromCommodityList(cl []*commodity.CommodityItem, config *conf.Config) *CommodityList {
	commodityList := make(CommodityList, len(cl))
	for i, c := range cl {
		commodityList[i] = FromCommodity(c, config)
	}
	return &commodityList
}

type GetCommoditiesResp struct {
	Commodities *CommodityList `json:"commodities"`
	Count       uint64         `json:"count"`
}

type PublishReq struct {
	Title     string  `form:"title"`
	Detail    string  `form:"detail"`
	Price     uint64  `form:"price"`
	Images    string  `form:"images"`
	Latitude  float64 `form:"latitude"`
	Longitude float64 `form:"longitude"`
}

type UpdateProfileReq struct {
	Username     string `form:"username"`
	Introduction string `form:"introduction"`
	AvatarUrl    string `form:"avatar_url"`
}
