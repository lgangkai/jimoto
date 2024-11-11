package biz

import (
	"commodity-service/dao/model"
	"commodity-service/service/commodity"
	"commodity-service/service/like"
	"context"
	"github.com/lgangkai/golog"
	cmdt "protos/commodity"
)

type CommodityBiz struct {
	commodityService *commodity.CommodityService
	likeService      *like.LikeService
	logger           *golog.Logger
}

func NewCommodityBiz(commodityService *commodity.CommodityService, likeService *like.LikeService, logger *golog.Logger) *CommodityBiz {
	return &CommodityBiz{
		commodityService: commodityService,
		likeService:      likeService,
		logger:           logger,
	}
}

func (b *CommodityBiz) GetCommodity(ctx context.Context, in *cmdt.GetCommodityRequest, out *cmdt.GetCommodityResponse) error {
	b.logger.Info(ctx, "Call CommodityBiz.GetCommodity, request: ", in)
	c, err := b.commodityService.GetCommodity(ctx, in.GetId())
	if err != nil {
		b.logger.Error(ctx, "Get commodity failed, err: ", err.Error())
		return err
	}
	imgs, err := b.commodityService.GetCommodityImages(ctx, in.GetId())
	if err != nil {
		b.logger.Error(ctx, "Get images failed, err: ", err.Error())
		return err
	}
	out.Commodity = &cmdt.CommodityItem{
		Id:        c.Id,
		CreatorId: c.CreatorId,
		Title:     c.Title,
		Detail:    c.Detail,
		Price:     c.Price,
		Cover:     c.Cover,
		Images:    imgs,
		Type:      c.Type,
		Status:    c.Status,
		Latitude:  c.Latitude,
		Longitude: c.Longitude,
	}
	b.logger.Info(ctx, "Call CommodityBiz.GetCommodity successfully.")
	return nil
}

func (b *CommodityBiz) GetCommodities(ctx context.Context, in *cmdt.GetCommoditiesRequest, out *cmdt.GetCommoditiesResponse) error {
	b.logger.Info(ctx, "Call CommodityBiz.GetCommodities, request: ", in)
	cList, count, err := b.commodityService.GetCommodities(ctx, in.GetFilterType(), in.GetOrderType(), in.GetLimit(), in.GetOffset())
	if err != nil {
		b.logger.Error(ctx, "Get commodity list failed, err: ", err.Error())
		return err
	}
	out.CommodityList = make([]*cmdt.CommodityItem, 0, in.GetLimit())
	out.Count = uint64(count)
	for _, c := range cList {
		out.CommodityList = append(out.CommodityList, &cmdt.CommodityItem{
			Id:        c.Id,
			CreatorId: c.CreatorId,
			Title:     c.Title,
			Detail:    c.Detail,
			Price:     c.Price,
			Cover:     c.Cover,
			Type:      c.Type,
			Status:    c.Status,
			Latitude:  c.Latitude,
			Longitude: c.Longitude,
		})
	}
	b.logger.Info(ctx, "Call CommodityBiz.GetCommodities successfully.")
	return nil
}

func (b *CommodityBiz) GetUserSoldCommodities(ctx context.Context, in *cmdt.GetUserSoldCommoditiesRequest, out *cmdt.GetUserSoldCommoditiesResponse) error {
	b.logger.Info(ctx, "Call CommodityBiz.GetUserSoldCommodities, request: ", in)
	cList, err := b.commodityService.GetUserSoldCommodities(ctx, in.GetUserId())
	if err != nil {
		b.logger.Error(ctx, "Get commodity list failed, err: ", err.Error())
		return err
	}
	out.CommodityList = make([]*cmdt.CommodityItem, 0, len(cList))
	for _, c := range cList {
		out.CommodityList = append(out.CommodityList, &cmdt.CommodityItem{
			Id:        c.Id,
			CreatorId: c.CreatorId,
			Title:     c.Title,
			Detail:    c.Detail,
			Price:     c.Price,
			Cover:     c.Cover,
			Type:      c.Type,
			Status:    c.Status,
			Latitude:  c.Latitude,
			Longitude: c.Longitude,
		})
	}
	b.logger.Info(ctx, "Call CommodityBiz.GetUserSoldCommodities successfully.")
	return nil
}

func (b *CommodityBiz) GetLatestCommodityList(ctx context.Context, in *cmdt.GetLatestCommodityListRequest, out *cmdt.GetLatestCommodityListResponse) error {
	b.logger.Info(ctx, "Call CommodityBiz.GetLatestCommodityList, request: ", in)
	cList, err := b.commodityService.GetLatestCommodityList(ctx, in.GetLimit(), in.GetOffset())
	if err != nil {
		b.logger.Error(ctx, "Get commodity list failed, err: ", err.Error())
		return err
	}
	out.CommodityList = make([]*cmdt.CommodityItem, 0, in.GetLimit())
	for _, c := range cList {
		out.CommodityList = append(out.CommodityList, &cmdt.CommodityItem{
			Id:        c.Id,
			CreatorId: c.CreatorId,
			Title:     c.Title,
			Detail:    c.Detail,
			Price:     c.Price,
			Cover:     c.Cover,
			Type:      c.Type,
			Status:    c.Status,
			Latitude:  c.Latitude,
			Longitude: c.Longitude,
		})
	}
	b.logger.Info(ctx, "Call CommodityBiz.GetLatestCommodityList successfully.")
	return nil
}

func (b *CommodityBiz) UpdateCommodity(ctx context.Context, in *cmdt.UpdateCommodityRequest, out *cmdt.UpdateCommodityResponse) error {
	b.logger.Info(ctx, "Call CommodityBiz.UpdateCommodity, request: ", in)
	c := in.GetCommodity()
	mc := &model.Commodity{
		Id:        c.Id,
		CreatorId: c.CreatorId,
		Title:     c.Title,
		Detail:    c.Detail,
		Price:     c.Price,
		Cover:     c.Cover,
		Type:      c.Type,
		Status:    c.Status,
		Latitude:  c.Latitude,
		Longitude: c.Longitude,
	}
	err := b.commodityService.UpdateCommodity(ctx, mc)
	if err != nil {
		b.logger.Error(ctx, "Update failed, err: ", err.Error())
		return err
	}
	b.logger.Info(ctx, "Call CommodityBiz.UpdateCommodity successfully.")
	return nil
}

func (b *CommodityBiz) DeleteCommodity(ctx context.Context, in *cmdt.DeleteCommodityRequest, out *cmdt.DeleteCommodityResponse) error {
	b.logger.Info(ctx, "Call CommodityBiz.DeleteCommodity, request: ", in)
	id := in.GetId()
	err := b.commodityService.DeleteCommodity(ctx, id)
	if err != nil {
		b.logger.Error(ctx, "Delete failed, err: ", err.Error())
		return err
	}
	b.logger.Info(ctx, "Call CommodityBiz.DeleteCommodity successfully.")
	return nil
}

func (b *CommodityBiz) PublishCommodity(ctx context.Context, in *cmdt.PublishCommodityRequest, out *cmdt.PublishCommodityResponse) error {
	b.logger.Info(ctx, "Call CommodityBiz.PublishCommodity, request: ", in)
	c := in.GetCommodity()
	mc := &model.Commodity{
		Id:        c.Id,
		CreatorId: c.CreatorId,
		Title:     c.Title,
		Detail:    c.Detail,
		Price:     c.Price,
		Cover:     c.Cover,
		Type:      c.Type,
		Status:    c.Status,
		Latitude:  c.Latitude,
		Longitude: c.Longitude,
		IsDeleted: false,
	}
	err := b.commodityService.PublishCommodity(ctx, mc, in.GetCommodity().Images)
	if err != nil {
		b.logger.Error(ctx, "PublishCommodity failed, err: ", err.Error())
		return err
	}
	b.logger.Info(ctx, "Call CommodityBiz.PublishCommodity successfully.")
	return nil
}

func (b *CommodityBiz) GetCommodityImages(ctx context.Context, in *cmdt.GetCommodityImagesRequest, out *cmdt.GetCommodityImagesResponse) error {
	b.logger.Info(ctx, "Call CommodityBiz.GetCommodityImages, request: ", in)
	images, err := b.commodityService.GetCommodityImages(ctx, in.GetId())
	if err != nil {
		b.logger.Error(ctx, "GetCommodityImages failed, err: ", err.Error())
		return err
	}
	b.logger.Info(ctx, "Call CommodityBiz.GetCommodityImages successfully.")
	out.Images = images
	return nil
}

func (b *CommodityBiz) LikeCommodity(ctx context.Context, in *cmdt.LikeCommodityRequest, out *cmdt.LikeCommodityResponse) error {
	b.logger.Info(ctx, "Call CommodityBiz.LikeCommodity, request: ", in)
	err := b.likeService.Like(ctx, in.GetUserId(), in.GetId())
	if err != nil {
		b.logger.Error(ctx, "LikeCommodity failed, err: ", err.Error())
		return err
	}
	return nil
}

func (b *CommodityBiz) UnlikeCommodity(ctx context.Context, in *cmdt.UnlikeCommodityRequest, out *cmdt.UnlikeCommodityResponse) error {
	b.logger.Info(ctx, "Call CommodityBiz.UnlikeCommodity, request: ", in)
	err := b.likeService.Unlike(ctx, in.GetUserId(), in.GetId())
	if err != nil {
		b.logger.Error(ctx, "UnlikeCommodity failed, err: ", err.Error())
		return err
	}
	return nil
}

func (b *CommodityBiz) GetCommodityLikedUsers(ctx context.Context, in *cmdt.GetCommodityLikedUsersRequest, out *cmdt.GetCommodityLikedUsersResponse) error {
	b.logger.Info(ctx, "Call CommodityBiz.GetCommodityLikedUsers, request: ", in)
	userIds, err := b.likeService.GetCommodityLikedUsers(ctx, in.GetId())
	if err != nil {
		b.logger.Error(ctx, "GetCommodityLikedUsers failed, err: ", err.Error())
		return err
	}
	b.logger.Info(ctx, "Call CommodityBiz.GetCommodityLikedUsers successfully.")
	out.UserIds = userIds
	return nil
}

func (b *CommodityBiz) GetUserLikeCommodities(ctx context.Context, in *cmdt.GetUserLikeCommoditiesRequest, out *cmdt.GetUserLikeCommoditiesResponse) error {
	b.logger.Info(ctx, "Call CommodityBiz.GetUserLikeCommodities, request: ", in)
	commodityList, err := b.likeService.GetUserLikeCommodities(ctx, in.GetId())
	if err != nil {
		b.logger.Error(ctx, "GetUserLikeCommodities failed, err: ", err.Error())
		return err
	}
	b.logger.Info(ctx, "Call CommodityBiz.GetUserLikeCommodities successfully.")
	out.CommodityList = make([]*cmdt.CommodityItem, 0)
	for _, c := range commodityList {
		out.CommodityList = append(out.CommodityList, &cmdt.CommodityItem{
			Id:        c.Id,
			CreatorId: c.CreatorId,
			Title:     c.Title,
			Detail:    c.Detail,
			Price:     c.Price,
			Cover:     c.Cover,
			Type:      c.Type,
			Status:    c.Status,
			Latitude:  c.Latitude,
			Longitude: c.Longitude,
		})
	}
	return nil
}
