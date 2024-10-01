package dao

import (
	model2 "commodity-service/dao/model"
	"context"
	"github.com/lgangkai/golog"
)

type CommodityImageDao struct {
	dbMaster *DBMaster
	dbSlave  *DBSlave
	logger   *golog.Logger
}

func NewCommodityImageDao(dbMaster *DBMaster, dbSlave *DBSlave, logger *golog.Logger) *CommodityImageDao {
	return &CommodityImageDao{
		dbMaster: dbMaster,
		dbSlave:  dbSlave,
		logger:   logger,
	}
}

func (d *CommodityImageDao) GetByCommodityId(ctx context.Context, id uint64) ([]*model2.CommodityImage, error) {
	d.logger.Info(ctx, "Call CommodityImageDao.GetByCommodityId.")
	var cis []*model2.CommodityImage

	if err := d.dbSlave.Table(model2.TAB_NAME_COMMODITY_IMAGE).Where("commodity_id = ?", id).Find(&cis).Error; err != nil {
		d.logger.Error(ctx, "Fail to get images, err: ", err.Error())
		return nil, err
	}
	return cis, nil
}

func (d *CommodityImageDao) Insert(ctx context.Context, commodityImages []*model2.CommodityImage) error {
	d.logger.Info(ctx, "Call CommodityImageDao.Insert.")
	if err := d.dbMaster.db(ctx).Create(commodityImages).Error; err != nil {
		d.logger.Error(ctx, "Fail to Insert images, err: ", err.Error())
		return err
	}
	return nil
}
