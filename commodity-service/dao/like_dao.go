package dao

import (
	"commodity-service/dao/model"
	"context"
	"github.com/lgangkai/golog"
)

type LikeDao struct {
	dbMaster *DBMaster
	dbSlave  *DBSlave
	logger   *golog.Logger
}

func NewLikeDao(dbMaster *DBMaster, dbSlave *DBSlave, logger *golog.Logger) *LikeDao {
	return &LikeDao{
		dbMaster: dbMaster,
		dbSlave:  dbSlave,
		logger:   logger,
	}
}

func (d *LikeDao) Query(ctx context.Context, like *model.Like) ([]*model.Like, error) {
	d.logger.Info(ctx, "Call LikeDao.Query.")
	var likes []*model.Like
	if err := d.dbSlave.Table(model.TAB_NAME_LIKE).Where(like).Order("update_time desc").Find(&likes).Error; err != nil {
		d.logger.Error(ctx, "Fail to query likes, err: ", err.Error())
		return nil, err
	}
	return likes, nil
}

func (d *LikeDao) Count(ctx context.Context, like *model.Like) (uint64, error) {
	d.logger.Info(ctx, "Call LikeDao.Count.")
	var count int64
	if err := d.dbSlave.Table(model.TAB_NAME_LIKE).Where(like).Count(&count).Error; err != nil {
		d.logger.Error(ctx, "Fail to count likes, err: ", err.Error())
		return 0, err
	}
	return uint64(count), nil
}

func (d *LikeDao) Insert(ctx context.Context, like *model.Like) error {
	d.logger.Info(ctx, "Call LikeDao.Insert.")
	if err := d.dbMaster.Table(model.TAB_NAME_LIKE).Create(like).Error; err != nil {
		d.logger.Error(ctx, "Fail to insert like, err: ", err.Error())
		return err
	}
	return nil
}

func (d *LikeDao) Delete(ctx context.Context, like *model.Like) error {
	d.logger.Info(ctx, "Call LikeDao.Delete.")
	if err := d.dbMaster.Table(model.TAB_NAME_LIKE).Where("user_id = ? AND commodity_id = ?", like.UserId, like.CommodityId).Delete(like).Error; err != nil {
		d.logger.Error(ctx, "Fail to delete like, err: ", err.Error())
		return err
	}
	return nil
}
