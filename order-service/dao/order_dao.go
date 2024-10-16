package dao

import (
	"context"
	"github.com/lgangkai/golog"
	"jimoto/order-service/dao/model"
)

type OrderDao struct {
	dbMaster *DBMaster
	dbSlave  *DBSlave
	logger   *golog.Logger
}

func NewOrderDao(dbMaster *DBMaster, dbSlave *DBSlave, logger *golog.Logger) *OrderDao {
	return &OrderDao{
		dbMaster: dbMaster,
		dbSlave:  dbSlave,
		logger:   logger,
	}
}

func (d *OrderDao) Insert(ctx context.Context, order *model.Order) error {
	d.logger.Info(ctx, "Call OrderDao.Insert.")
	if err := d.dbMaster.Table(model.TAB_NAME_ORDER).Create(order).Error; err != nil {
		d.logger.Error(ctx, "Fail to insert order, err: ", err.Error())
		return err
	}
	return nil
}
