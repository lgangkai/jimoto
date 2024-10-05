package handler

import (
	"protos/commodity"
)

type PublishRequest struct {
	Title  string `form:"title"`
	Detail string `form:"detail"`
	Price  uint64 `form:"price"`
	Images string `form:"images"`
}

type Commodity struct {
	Id     uint64 `json:"id"`
	Title  string `json:"title"`
	Price  uint64 `json:"price"`
	Cover  string `json:"cover"`
	Type   uint32 `json:"type"`
	Status uint32 `json:"status"`
}

type CreateProfileRequest struct {
	Username string `form:"username"`
	Birthday string `form:"birthday"`
	Email    string `form:"email"`
}

type Image struct {
	Image string `json:"image"`
}

type AccountRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type LoginResp struct {
	UserId uint64 `json:"user_id"`
}

type UploadResp struct {
	Filename string `json:"filename"`
}

type User struct {
	UserId uint64 `json:"user_id"`
}

type CommodityList []*Commodity

func From(c *commodity.CommodityItem) *Commodity {
	return &Commodity{
		Id:     c.Id,
		Title:  c.Title,
		Price:  c.Price,
		Cover:  c.Cover,
		Type:   c.Type,
		Status: c.Status,
	}
}

func FromList(cl []*commodity.CommodityItem) *CommodityList {
	commodityList := make(CommodityList, len(cl))
	for i, c := range cl {
		commodityList[i] = From(c)
	}
	return &commodityList
}

func (c *Commodity) CompleteImageUrl(client *Client) {
	c.Cover = client.CompleteImageUrl(c.Cover)
}

func (cl *CommodityList) CompleteImageUrlForList(client *Client) {
	for i := 0; i < len(*cl); i++ {
		(*cl)[i].CompleteImageUrl(client)
	}
}
