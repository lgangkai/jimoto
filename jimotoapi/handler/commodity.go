package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"jimotoapi/util"
	"jimotoapi/vo"
	"protos/commodity"
	"strconv"
	"strings"
)

const (
	PARAM_KEY_FILTER_TYPE = "filter_type"
	PARAM_KEY_TYPE        = "type"
	PARAM_KEY_ORDER_TYPE  = "order_type"
	PARAM_KEY_PAGE        = "page"
	PARAM_KEY_PAGE_SIZE   = "page_size"
)

var (
	PARAM_FILTER_PUBLISH_MAP = map[string]commodity.FilterPublish{
		"":           commodity.FilterPublish_ALL,
		"all":        commodity.FilterPublish_ALL,
		"publishing": commodity.FilterPublish_PUBLISHING,
	}
	PARAM_FILTER_SELL_MAP = map[string]commodity.FilterSell{
		"":     commodity.FilterSell_SELL,
		"sell": commodity.FilterSell_SELL,
		"buy":  commodity.FilterSell_BUY,
	}
	PARAM_ORDER_MAP = map[string]commodity.OrderType{
		"":           commodity.OrderType_LATEST,
		"latest":     commodity.OrderType_LATEST,
		"nearest":    commodity.OrderType_NEAREST,
		"cheapest":   commodity.OrderType_CHEAPEST,
		"highest":    commodity.OrderType_HIGHEST,
		"most-liked": commodity.OrderType_MOST_LIKED,
	}
)

func (c *Client) PublishCommodity(context *gin.Context) {
	authData, ok := c.ParseAuthData(context)
	if !ok {
		return
	}
	p := &vo.PublishReq{}
	if err := context.ShouldBind(p); err != nil {
		c.HandleRequestError(context, err)
		return
	}
	imgs := strings.Split(p.Images, ",")
	cm := &commodity.CommodityItem{
		CreatorId: authData.UserId,
		Title:     p.Title,
		Detail:    p.Detail,
		Price:     p.Price,
		Images:    imgs,
		Cover:     imgs[0],
		Latitude:  p.Latitude,
		Longitude: p.Longitude,
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
	cr.Images = util.CompleteImageUrls(cr.Images, c.config)
	cm, err := json.Marshal(cr)
	if err != nil {
		c.HandleJsonError(context, err)
		return
	}
	c.logger.Info(c.context, "Handle get commodity success, commodity: ", string(cm))
	c.HandleSuccess(context, cm)
}

func (c *Client) GetCommodities(context *gin.Context) {
	fType := context.Query(PARAM_KEY_FILTER_TYPE)
	tType := context.Query(PARAM_KEY_TYPE)
	oType := context.Query(PARAM_KEY_ORDER_TYPE)
	page := context.Query(PARAM_KEY_PAGE)
	pageSize := context.Query(PARAM_KEY_PAGE_SIZE)
	r := &commodity.GetCommoditiesRequest{
		Limit:  uint64(util.Str2Num(pageSize)),
		Offset: uint64((util.Str2Num(page) - 1) * util.Str2Num(pageSize)),
		Filter: &commodity.Filter{
			FilterPublish: PARAM_FILTER_PUBLISH_MAP[fType],
			FilterSell:    PARAM_FILTER_SELL_MAP[tType],
		},
		OrderType: PARAM_ORDER_MAP[oType],
		RequestId: GetRequestId(context),
	}
	resp, err := c.commodityClient.GetCommodities(context, r)
	if err != nil {
		c.HandleRpcError(context, err)
		return
	}

	rsp := vo.GetCommoditiesResp{
		Commodities: vo.FromCommodityList(resp.GetCommodityList(), c.config),
		Count:       resp.GetCount(),
	}
	cms, err := json.Marshal(rsp)
	if err != nil {
		c.HandleJsonError(context, err)
		return
	}

	c.logger.Info(c.context, "Handle get commodity list success, commodities: ", string(cms))
	c.HandleSuccess(context, cms)
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

	cl := vo.FromCommodityList(resp.GetCommodityList(), c.config)
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

	cis := make([]*vo.Image, len(resp.Images))
	for i, image := range resp.Images {
		cis[i] = &vo.Image{Image: util.CompleteImageUrl(image, c.config)}
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
	authData, ok := c.ParseAuthData(context)
	if !ok {
		return
	}
	cId, _ := strconv.Atoi(context.Param("commodity_id"))
	r := &commodity.LikeCommodityRequest{
		Id:        uint64(cId),
		UserId:    authData.UserId,
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
	authData, ok := c.ParseAuthData(context)
	if !ok {
		return
	}
	cId, _ := strconv.Atoi(context.Param("commodity_id"))
	r := &commodity.UnlikeCommodityRequest{
		Id:        uint64(cId),
		UserId:    authData.UserId,
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
	users := make([]*vo.User, len(resp.UserIds))
	for i, id := range resp.UserIds {
		users[i] = &vo.User{UserId: id}
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
	authData, ok := c.ParseAuthData(context)
	if !ok {
		return
	}
	r := &commodity.GetUserLikeCommoditiesRequest{
		Id:        authData.UserId,
		RequestId: GetRequestId(context),
	}
	resp, err := c.commodityClient.GetUserLikeCommodities(context, r)
	if err != nil {
		c.HandleRpcError(context, err)
		return
	}
	cl := vo.FromCommodityList(resp.GetCommodityList(), c.config)
	cms, err := json.Marshal(cl)
	if err != nil {
		c.HandleJsonError(context, err)
		return
	}
	c.logger.Info(c.context, "Handle GetUserLikeCommodities success, commodities: ", string(cms))
	c.HandleSuccess(context, cms)
}

func (c *Client) GetUserSoldCommodities(context *gin.Context) {
	authData, ok := c.ParseAuthData(context)
	if !ok {
		return
	}
	r := &commodity.GetUserSoldCommoditiesRequest{
		UserId:    authData.UserId,
		RequestId: GetRequestId(context),
	}
	resp, err := c.commodityClient.GetUserSoldCommodities(context, r)
	if err != nil {
		c.HandleJsonError(context, err)
		return
	}
	cl := vo.FromCommodityList(resp.GetCommodityList(), c.config)
	cms, err := json.Marshal(cl)
	if err != nil {
		c.HandleJsonError(context, err)
		return
	}
	c.logger.Info(c.context, "Handle GetUserSoldCommodities success, commodities: ", string(cms))
	c.HandleSuccess(context, cms)
}
