package like

import (
	"commodity-service/dao"
	"commodity-service/dao/model"
	"context"
	errs "errs"
	"github.com/lgangkai/golog"
)

type LikeService struct {
	likeDao      *dao.LikeDao
	commodityDao *dao.CommodityDao
	logger       *golog.Logger
}

func NewLikeService(likeDao *dao.LikeDao, commodityDao *dao.CommodityDao, logger *golog.Logger) *LikeService {
	return &LikeService{
		likeDao:      likeDao,
		commodityDao: commodityDao,
		logger:       logger,
	}
}

func (s *LikeService) Like(ctx context.Context, userId uint64, commodityId uint64) error {
	s.logger.Info(ctx, "Call LikeService.Like")
	err := s.likeDao.Insert(ctx, &model.Like{
		CommodityId: commodityId,
		UserId:      userId,
	})
	if err != nil {
		s.logger.Error(ctx, "Fail to like commodity, err:", err.Error())
		return errs.New(errs.ERR_INTERNAL_ERROR)
	}
	return nil
}

func (s *LikeService) Unlike(ctx context.Context, userId uint64, commodityId uint64) error {
	s.logger.Info(ctx, "Call LikeService.Unlike")
	err := s.likeDao.Delete(ctx, &model.Like{
		CommodityId: commodityId,
		UserId:      userId,
	})
	if err != nil {
		s.logger.Error(ctx, "Fail to unlike commodity, err:", err.Error())
		return errs.New(errs.ERR_INTERNAL_ERROR)
	}
	return nil
}

func (s *LikeService) GetCommodityLikedUsers(ctx context.Context, commodityId uint64) ([]uint64, error) {
	s.logger.Info(ctx, "Call LikeService.GetCommodityLikedUsers")
	likes, err := s.likeDao.Query(ctx, &model.Like{
		CommodityId: commodityId,
	})
	if err != nil {
		s.logger.Error(ctx, "Fail to query commodity liked users, err:", err.Error())
		return nil, errs.New(errs.ERR_INTERNAL_ERROR)
	}
	userIds := make([]uint64, len(likes))
	for i, like := range likes {
		userIds[i] = like.UserId
	}
	return userIds, nil
}

func (s *LikeService) GetUserLikeCommodities(ctx context.Context, id uint64) ([]*model.Commodity, error) {
	s.logger.Info(ctx, "Call LikeService.GetUserLikeCommodities")
	// 1. query ids
	likes, err := s.likeDao.Query(ctx, &model.Like{
		UserId: id,
	})
	if err != nil {
		s.logger.Error(ctx, "Fail to query user like commodities, err:", err.Error())
		return nil, errs.New(errs.ERR_INTERNAL_ERROR)
	}
	cmIds := make([]uint64, len(likes))
	for i, like := range likes {
		cmIds[i] = like.CommodityId
	}
	// 2. query commodities from ids
	cms, err := s.commodityDao.GetListByIds(ctx, cmIds)
	if err != nil {
		s.logger.Error(ctx, "Fail to query commodities, err:", err.Error())
		return nil, errs.New(errs.ERR_INTERNAL_ERROR)
	}
	// 3. sort cms as order in cmIds
	// because we want to show liked items order by the liked time
	id2index := make(map[uint64]int)
	for i, cmId := range cmIds {
		id2index[cmId] = i
	}
	cmsOrdered := make([]*model.Commodity, len(cms))
	for _, cm := range cms {
		index, _ := id2index[cm.Id]
		cmsOrdered[index] = cm
	}

	return cmsOrdered, nil
}
