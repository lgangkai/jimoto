package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"protos/commodity"
	"strconv"
	"strings"
)

func (c *Client) PublishCommodity(context *gin.Context) {
	uId := c.getAuthedData(context, KEY_USER_ID)
	p := &PublishRequest{}
	if err := context.ShouldBind(p); err != nil {
		c.HandleRequestError(context, err)
		return
	}
	imgs := strings.Split(p.Images, ",")
	cm := &commodity.CommodityItem{
		CreatorId: uId.(uint64),
		Title:     p.Title,
		Detail:    p.Detail,
		Price:     p.Price,
		Images:    imgs,
		Cover:     imgs[0],
	}
	pr := &commodity.PublishCommodityRequest{
		Commodity: cm,
		RequestId: GetRequestId(context),
	}
	_, err := c.commodityClient.PublishCommodity(context, pr)
	if err != nil {
		c.HandleRpcError(context, err)
		return
	}
	c.HandleSuccess(context, nil)
}

func (c *Client) GetCommodity(context *gin.Context) {
	cId, _ := strconv.Atoi(context.Param("commodity_id"))
	r := &commodity.GetCommodityRequest{
		Id:        uint64(cId),
		RequestId: GetRequestId(context),
	}
	resp, err := c.commodityClient.GetCommodity(context, r)
	if err != nil {
		c.HandleRpcError(context, err)
		return
	}

	cr := resp.GetCommodity()
	cr.Images = c.CompleteImageUrls(cr.Images)
	cm, err := json.Marshal(cr)
	if err != nil {
		c.HandleJsonError(context, err)
		return
	}
	c.logger.Info(c.context, "Handle get commodity success, commodity: ", string(cm))
	c.HandleSuccess(context, cm)
}

func (c *Client) GetLatestCommodityList(context *gin.Context) {
	r := &commodity.GetLatestCommodityListRequest{
		Limit:     10,
		Offset:    0,
		RequestId: GetRequestId(context),
	}
	resp, err := c.commodityClient.GetLatestCommodityList(context, r)
	if err != nil {
		c.HandleRpcError(context, err)
		return
	}

	cl := FromList(resp.GetCommodityList())
	cl.CompleteImageUrlForList(c)
	cms, err := json.Marshal(cl)
	if err != nil {
		c.HandleJsonError(context, err)
		return
	}

	c.logger.Info(c.context, "Handle get commodity list success, commodities: ", string(cms))
	c.HandleSuccess(context, cms)
}

func (c *Client) GetCommodityImages(context *gin.Context) {
	cId, _ := strconv.Atoi(context.Param("commodity_id"))
	r := &commodity.GetCommodityImagesRequest{
		Id:        uint64(cId),
		RequestId: GetRequestId(context),
	}
	resp, err := c.commodityClient.GetCommodityImages(context, r)
	if err != nil {
		c.HandleRpcError(context, err)
		return
	}

	cis := make([]*Image, len(resp.Images))
	for i, image := range resp.Images {
		cis[i] = &Image{Image: c.CompleteImageUrl(image)}
	}
	cms, err := json.Marshal(cis)
	if err != nil {
		c.HandleJsonError(context, err)
		return
	}

	c.logger.Info(c.context, "Handle get commodity images success, images: ", string(cms))
	c.HandleSuccess(context, cms)
}

func (c *Client) LikeCommodity(context *gin.Context) {
	uId := c.getAuthedData(context, KEY_USER_ID)
	cId, _ := strconv.Atoi(context.Param("commodity_id"))
	if uId == nil {
		return
	}
	r := &commodity.LikeCommodityRequest{
		Id:        uint64(cId),
		UserId:    uId.(uint64),
		RequestId: GetRequestId(context),
	}
	_, err := c.commodityClient.LikeCommodity(context, r)
	if err != nil {
		c.HandleRpcError(context, err)
		return
	}
	c.HandleSuccess(context, nil)
}

func (c *Client) UnlikeCommodity(context *gin.Context) {
	uId := c.getAuthedData(context, KEY_USER_ID)
	cId, _ := strconv.Atoi(context.Param("commodity_id"))
	if uId == nil {
		return
	}
	r := &commodity.UnlikeCommodityRequest{
		Id:        uint64(cId),
		UserId:    uId.(uint64),
		RequestId: GetRequestId(context),
	}
	_, err := c.commodityClient.UnlikeCommodity(context, r)
	if err != nil {
		c.HandleRpcError(context, err)
		return
	}
	c.HandleSuccess(context, nil)
}

func (c *Client) GetLikedUsers(context *gin.Context) {
	cId, _ := strconv.Atoi(context.Param("commodity_id"))
	r := &commodity.GetCommodityLikedUsersRequest{
		Id:        uint64(cId),
		RequestId: GetRequestId(context),
	}
	resp, err := c.commodityClient.GetCommodityLikedUsers(context, r)
	if err != nil {
		c.HandleRpcError(context, err)
		return
	}
	users := make([]*User, len(resp.UserIds))
	for i, id := range resp.UserIds {
		users[i] = &User{UserId: id}
	}
	urs, err := json.Marshal(users)
	if err != nil {
		c.HandleJsonError(context, err)
		return
	}
	c.logger.Info(c.context, "Handle GetLikedUsers success, users: ", string(urs))
	c.HandleSuccess(context, urs)
}

func (c *Client) GetUserLikeCommodities(context *gin.Context) {
	uId := c.getAuthedData(context, KEY_USER_ID)
	if uId == nil {
		return
	}
	r := &commodity.GetUserLikeCommoditiesRequest{
		Id:        uId.(uint64),
		RequestId: GetRequestId(context),
	}
	resp, err := c.commodityClient.GetUserLikeCommodities(context, r)
	if err != nil {
		c.HandleRpcError(context, err)
		return
	}
	cl := FromList(resp.GetCommodityList())
	cl.CompleteImageUrlForList(c)
	cms, err := json.Marshal(cl)
	if err != nil {
		c.HandleJsonError(context, err)
		return
	}
	c.logger.Info(c.context, "Handle GetUserLikeCommodities success, commodities: ", string(cms))
	c.HandleSuccess(context, cms)
}
