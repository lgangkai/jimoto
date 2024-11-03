package commodity

import (
	"commodity-service/dao"
	"commodity-service/dao/model"
	"context"
	errs "errs"
	"github.com/lgangkai/golog"
	cmdt "protos/commodity"
)

const (
	COMMODITY_STATUS_PUBLISHING = 0
	COMMODITY_STATUS_DONE       = 1
	COMMODITY_STATUS_REMOVED    = 2
)

type CommodityService struct {
	commodityDao      *dao.CommodityDao
	commodityImageDao *dao.CommodityImageDao
	db                *dao.DBMaster
	logger            *golog.Logger
}

func NewCommodityService(commodityDao *dao.CommodityDao, commodityImageDao *dao.CommodityImageDao, db *dao.DBMaster, logger *golog.Logger) *CommodityService {
	return &CommodityService{
		commodityDao:      commodityDao,
		commodityImageDao: commodityImageDao,
		db:                db,
		logger:            logger,
	}
}

func (s *CommodityService) GetCommodity(ctx context.Context, id uint64) (*model.Commodity, error) {
	s.logger.Info(ctx, "Call CommodityService.GetCommodity")
	commodity, err := s.commodityDao.GetById(ctx, id)
	if err != nil {
		s.logger.Error(ctx, "Fail to get commodity, err:", err.Error())
		return nil, errs.New(errs.ERR_GET_COMMODITY_FAILED)
	}
	return commodity, nil
}

func (s *CommodityService) GetCommodities(ctx context.Context, filterType cmdt.FilterType, orderType cmdt.OrderType, pageSize uint64, offset uint64) ([]*model.Commodity, int64, error) {
	s.logger.Info(ctx, "Call CommodityService.GetCommodities")
	cList, count, err := s.commodityDao.GetListByFilter(ctx, filterType, orderType, pageSize, offset)
	if err != nil {
		s.logger.Error(ctx, "Fail to get commodity list, err:", err.Error())
		return nil, 0, errs.New(errs.ERR_GET_COMMODITY_LIST_FAILED)
	}
	return cList, count, err
}

func (s *CommodityService) GetLatestCommodityList(ctx context.Context, pageSize uint64, offset uint64) ([]*model.Commodity, error) {
	s.logger.Info(ctx, "Call CommodityService.GetLatestCommodityList")
	cList, err := s.commodityDao.GetListLatest(ctx, pageSize, offset)
	if err != nil {
		s.logger.Error(ctx, "Fail to get commodity list, err:", err.Error())
		return nil, errs.New(errs.ERR_GET_COMMODITY_LIST_FAILED)
	}
	return cList, err
}

func (s *CommodityService) UpdateCommodity(ctx context.Context, commodity *model.Commodity) error {
	s.logger.Info(ctx, "Call CommodityService.UpdateCommodity")
	err := s.commodityDao.Update(ctx, commodity)
	if err != nil {
		s.logger.Error(ctx, "Fail to update commodity, err:", err.Error())
		return errs.New(errs.ERR_UPDATE_COMMODITY_FAILED)
	}
	return nil
}

func (s *CommodityService) DeleteCommodity(ctx context.Context, id uint64) error {
	s.logger.Info(ctx, "Call CommodityService.DeleteCommodity")
	err := s.commodityDao.Delete(ctx, id)
	if err != nil {
		s.logger.Error(ctx, "Fail to delete commodity, err:", err.Error())
		return errs.New(errs.ERR_DELETE_COMMODITY_FAILED)
	}
	return nil
}

func (s *CommodityService) PublishCommodity(ctx context.Context, commodity *model.Commodity, images []string) error {
	s.logger.Info(ctx, "Call CommodityService.CreateCommodity, commodity: ", commodity)
	// set status to publishing
	commodity.Status = COMMODITY_STATUS_PUBLISHING
	commodityImages := make([]*model.CommodityImage, len(images))
	// start transaction
	return s.db.ExecTx(ctx, func(ctx context.Context) error {
		err := s.commodityDao.Insert(ctx, commodity)
		// get id after insert
		for i, image := range images {
			commodityImages[i] = &model.CommodityImage{
				CommodityId: commodity.Id,
				Image:       image,
			}
		}
		err = s.commodityImageDao.Insert(ctx, commodityImages)
		if err != nil {
			s.logger.Error(ctx, "Fail to publish commodity, err:", err.Error())
			return errs.New(errs.ERR_CREATE_COMMODITY_FAILED)
		}
		return nil
	})
}

func (s *CommodityService) GetCommodityImages(ctx context.Context, id uint64) ([]string, error) {
	s.logger.Info(ctx, "Call CommodityService.GetByCommodityId.")
	cis, err := s.commodityImageDao.GetByCommodityId(ctx, id)
	if err != nil {
		s.logger.Error(ctx, "Fail to get commodity images, err:", err.Error())
		return nil, errs.New(errs.ERR_GET_COMMODITY_IMAGES_FAILED)
	}
	images := make([]string, len(cis))
	for i, ci := range cis {
		images[i] = ci.Image
	}
	return images, nil
}
