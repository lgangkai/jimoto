package handler

import (
	commodity2 "commodity-service/biz"
	"context"
	"protos/commodity"
)

type CommodityHandlerImpl struct {
	commodityBiz *commodity2.CommodityBiz
}

func NewCommodityHandlerImpl(commodityBiz *commodity2.CommodityBiz) *CommodityHandlerImpl {
	return &CommodityHandlerImpl{commodityBiz: commodityBiz}
}

func (h *CommodityHandlerImpl) GetCommodity(ctx context.Context, in *commodity.GetCommodityRequest, out *commodity.GetCommodityResponse) error {
	return h.commodityBiz.GetCommodity(getTraceContext(ctx, in.RequestId, in.Id), in, out)
}

func (h *CommodityHandlerImpl) GetLatestCommodityList(ctx context.Context, in *commodity.GetLatestCommodityListRequest, out *commodity.GetLatestCommodityListResponse) error {
	return h.commodityBiz.GetLatestCommodityList(getTraceContext(ctx, in.RequestId, 0), in, out)
}

func (h *CommodityHandlerImpl) GetUserSoldCommodities(ctx context.Context, in *commodity.GetUserSoldCommoditiesRequest, out *commodity.GetUserSoldCommoditiesResponse) error {
	return h.commodityBiz.GetUserSoldCommodities(getTraceContext(ctx, in.RequestId, 0), in, out)
}

func (h *CommodityHandlerImpl) PublishCommodity(ctx context.Context, in *commodity.PublishCommodityRequest, out *commodity.PublishCommodityResponse) error {
	return h.commodityBiz.PublishCommodity(getTraceContext(ctx, in.RequestId, in.Commodity.Id), in, out)
}

func (h *CommodityHandlerImpl) DeleteCommodity(ctx context.Context, in *commodity.DeleteCommodityRequest, out *commodity.DeleteCommodityResponse) error {
	return h.commodityBiz.DeleteCommodity(getTraceContext(ctx, in.RequestId, in.Id), in, out)
}

func (h *CommodityHandlerImpl) UpdateCommodity(ctx context.Context, in *commodity.UpdateCommodityRequest, out *commodity.UpdateCommodityResponse) error {
	return h.commodityBiz.UpdateCommodity(getTraceContext(ctx, in.RequestId, in.Commodity.Id), in, out)
}

func (h *CommodityHandlerImpl) GetCommodityImages(ctx context.Context, in *commodity.GetCommodityImagesRequest, out *commodity.GetCommodityImagesResponse) error {
	return h.commodityBiz.GetCommodityImages(getTraceContext(ctx, in.RequestId, in.Id), in, out)
}

func (h *CommodityHandlerImpl) LikeCommodity(ctx context.Context, in *commodity.LikeCommodityRequest, out *commodity.LikeCommodityResponse) error {
	return h.commodityBiz.LikeCommodity(getTraceContext(ctx, in.RequestId, in.Id), in, out)
}

func (h *CommodityHandlerImpl) UnlikeCommodity(ctx context.Context, in *commodity.UnlikeCommodityRequest, out *commodity.UnlikeCommodityResponse) error {
	return h.commodityBiz.UnlikeCommodity(getTraceContext(ctx, in.RequestId, in.Id), in, out)
}

func (h *CommodityHandlerImpl) GetCommodityLikedUsers(ctx context.Context, in *commodity.GetCommodityLikedUsersRequest, out *commodity.GetCommodityLikedUsersResponse) error {
	return h.commodityBiz.GetCommodityLikedUsers(getTraceContext(ctx, in.RequestId, in.Id), in, out)
}

func (h *CommodityHandlerImpl) GetUserLikeCommodities(ctx context.Context, in *commodity.GetUserLikeCommoditiesRequest, out *commodity.GetUserLikeCommoditiesResponse) error {
	return h.commodityBiz.GetUserLikeCommodities(getTraceContext(ctx, in.RequestId, in.Id), in, out)
}

func (h *CommodityHandlerImpl) GetCommodities(ctx context.Context, in *commodity.GetCommoditiesRequest, out *commodity.GetCommoditiesResponse) error {
	return h.commodityBiz.GetCommodities(getTraceContext(ctx, in.RequestId, 0), in, out)
}

func getTraceContext(ctx context.Context, requestId string, commodityId uint64) context.Context {
	return context.WithValue(ctx, "traceKey", map[string]any{
		"request_id":   requestId,
		"commodity_id": commodityId,
	})
}
