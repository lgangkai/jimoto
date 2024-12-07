package dao

import (
	model2 "commodity-service/dao/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lgangkai/golog"
	"github.com/redis/go-redis/v9"
	"math/rand"
	cmdt "protos/commodity"
	"time"
)

const (
	REDIS_KEY_GET_COMMODITY_PREFIX           = "commodity:get_commodity:"
	REDIS_KEY_GET_COMMODITY_EXPIRE_BASE      = time.Second * 60
	REDIS_KEY_GET_COMMODITY_EXPIRE_MAX_SHIFT = 30
)

type CommodityDao struct {
	dbMaster *DBMaster
	dbSlave  *DBSlave
	dbRedis  *redis.Client
	logger   *golog.Logger
}

func NewCommodityDao(dbMaster *DBMaster, dbSlave *DBSlave, dbRedis *redis.Client, logger *golog.Logger) *CommodityDao {
	return &CommodityDao{
		dbMaster: dbMaster,
		dbSlave:  dbSlave,
		dbRedis:  dbRedis,
		logger:   logger,
	}
}

func (d *CommodityDao) GetById(ctx context.Context, id uint64) (*model2.Commodity, error) {
	d.logger.Info(ctx, "Call CommodityDao.GetById.")
	commodity := &model2.Commodity{}

	// 1. try to get value from redis first.
	rKey := fmt.Sprintf("%v%d", REDIS_KEY_GET_COMMODITY_PREFIX, id)
	commStr, err := d.dbRedis.Get(ctx, rKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			d.logger.Info(ctx, "Can not find in cache, go to sql DB.")
		} else {
			d.logger.Error(ctx, "Can not get from cache, err: ", err.Error(), ". Go to sql DB")
		}
	} else {
		d.logger.Info(ctx, "Get commodity json from cache, commodity: ", commStr)
		err = json.Unmarshal([]byte(commStr), commodity)
		if err != nil {
			d.logger.Error(ctx, "json.Unmarshal failed, err: ", err.Error(), ". Go to sql DB")
		} else {
			d.logger.Info(ctx, "Get commodity from cache succeeded.")
			return commodity, nil
		}
	}

	// 2. get value from mysql-slave if not found in redis.
	if err = d.dbSlave.Where("id = ? AND is_deleted = ?", id, false).Take(commodity).Error; err != nil {
		d.logger.Error(ctx, "Fail to get data, err: ", err.Error())
		return nil, err
	}
	d.logger.Info(ctx, "Get commodity done, commodity: ", commodity)

	// 3. write profile as json string back to cache.
	cBytes, err := json.Marshal(commodity)
	if err != nil {
		d.logger.Error(ctx, "json.Marshal failed, err: ", err.Error(), ". It will not be saved to cache.")
		return commodity, nil
	}
	// set key expiration time as base time plus random time to avoid cache avalanche.
	randExp := time.Duration(rand.Intn(REDIS_KEY_GET_COMMODITY_EXPIRE_MAX_SHIFT)) * time.Second
	err = d.dbRedis.Set(ctx, rKey, string(cBytes), REDIS_KEY_GET_COMMODITY_EXPIRE_BASE+randExp).Err()
	if err != nil {
		d.logger.Error(ctx, "redis set failed, err: ", err.Error(), ". It will not be saved to cache.")
		return commodity, nil
	}
	d.logger.Info(ctx, "Write back to redis done. key: ", rKey, ", value: ", string(cBytes),
		", expiration: ", REDIS_KEY_GET_COMMODITY_EXPIRE_BASE+randExp)

	return commodity, nil
}

func (d *CommodityDao) GetListByFilter(ctx context.Context, filter *cmdt.Filter, orderType cmdt.OrderType, pageSize uint64, offset uint64) ([]*model2.Commodity, int64, error) {
	d.logger.Info(ctx, "Call CommodityDao.GetListByFilter.")
	var cmList []*model2.Commodity
	var orderPhrase string
	var count int64
	switch orderType {
	case cmdt.OrderType_LATEST:
		orderPhrase = "update_time desc"
	case cmdt.OrderType_CHEAPEST:
		orderPhrase = "price asc"
	case cmdt.OrderType_HIGHEST:
		orderPhrase = "price desc"
	default:
		orderPhrase = "update_time desc"
	}
	whereClause, args := buildWhereClauseByFilter(filter)
	err := d.dbSlave.Table(model2.TAB_NAME_COMMODITY).Where(whereClause, args...).Order(orderPhrase).Limit(int(pageSize)).Offset(int(offset)).Find(&cmList).Offset(-1).Limit(-1).Count(&count).Error
	if err != nil {
		d.logger.Error(ctx, "Fail to get data, err: ", err.Error())
		return nil, 0, err
	}
	d.logger.Info(ctx, "Get list success.")
	return cmList, count, nil
}

func (d *CommodityDao) GetListByCreatorId(ctx context.Context, creatorId uint64) ([]*model2.Commodity, error) {
	d.logger.Info(ctx, "Call CommodityDao.GetListByCreatorId.")
	var cmList []*model2.Commodity
	if err := d.dbSlave.Where("creator_id = ?", creatorId).Find(&cmList).Error; err != nil {
		d.logger.Error(ctx, "Fail to get data, err: ", err.Error())
		return nil, err
	}
	d.logger.Info(ctx, "Get list success.")
	return cmList, nil
}

func (d *CommodityDao) GetListByIds(ctx context.Context, ids []uint64) ([]*model2.Commodity, error) {
	d.logger.Info(ctx, "Call CommodityDao.GetListByIds.")
	var cmList []*model2.Commodity
	if err := d.dbSlave.Where("id IN ?", ids).Find(&cmList).Error; err != nil {
		d.logger.Error(ctx, "Fail to get data, err: ", err.Error())
		return nil, err
	}
	d.logger.Info(ctx, "Get list success.")
	return cmList, nil
}

func (d *CommodityDao) GetListLatest(ctx context.Context, pageSize uint64, offset uint64) ([]*model2.Commodity, error) {
	d.logger.Info(ctx, "Call CommodityDao.GetListLatest.")
	var cList []*model2.Commodity
	var count int64

	if err := d.dbSlave.Table(model2.TAB_NAME_COMMODITY).Where("is_deleted = ?", false).Limit(int(pageSize)).Order("update_time desc").Offset(int(offset)).Count(&count).Find(&cList).Error; err != nil {
		d.logger.Error(ctx, "Fail to get data, err: ", err.Error())
		return nil, err
	}
	d.logger.Info(ctx, "Get list success. Total number: ", count)
	return cList, nil
}

func (d *CommodityDao) Insert(ctx context.Context, commodity *model2.Commodity) error {
	// we don't operate cache in insert. Cache data will be load when read.
	d.logger.Info(ctx, "Call CommodityDao.Insert, commodity: ", commodity)

	if err := d.dbMaster.db(ctx).Create(commodity).Error; err != nil {
		d.logger.Error(ctx, "Fail to insert into sql DB, err: ", err.Error())
		return err
	}
	d.logger.Info(ctx, "Insert data into sql DB succeed.")

	return nil
}

func (d *CommodityDao) Update(ctx context.Context, commodity *model2.Commodity) error {
	// use cache aside pattern to update DB and then delete from cache.
	// 1. update data to mysql-master.
	d.logger.Info(ctx, "Call CommodityDao.Update.")
	if err := d.dbMaster.db(ctx).Where("id = ? ", commodity.Id).UpdateColumns(commodity).Error; err != nil {
		d.logger.Error(ctx, "Fail to update to sql DB, err: ", err.Error())
		return err
	}
	d.logger.Info(ctx, "Update data to sql DB succeed.")

	// 2. delete data from redis.
	d.deleteFromCache(ctx, commodity.Id)

	return nil
}

func (d *CommodityDao) Delete(ctx context.Context, id uint64) error {
	// use cache aside pattern to delete from DB and then delete from cache.
	// 1. delete data from mysql-master.
	// not really deleted, set is_deleted field to 1.
	d.logger.Info(ctx, "Call CommodityDao.Delete.")
	//if err := d.dbMaster.Where("id = ?", id).Delete(&model.Commodity{}).Error; err != nil {
	//	d.logger.Error(ctx, "Fail to delete from sql DB, err: ", err.Error())
	//	return err
	//}
	if err := d.dbMaster.db(ctx).Save(&model2.Commodity{Id: id, IsDeleted: true}).Error; err != nil {
		d.logger.Error(ctx, "Fail to delete from sql DB, err: ", err.Error())
		return err
	}
	d.logger.Info(ctx, "Delete data from sql DB succeed.")

	// 2. delete data from redis.
	d.deleteFromCache(ctx, id)

	return nil
}

func (d *CommodityDao) deleteFromCache(ctx context.Context, id uint64) {
	rKey := fmt.Sprintf("%v%d", REDIS_KEY_GET_COMMODITY_PREFIX, id)
	err := d.dbRedis.Del(ctx, rKey).Err()
	// if delete failed, other process might read the dirty data.
	// since we have set an expiration time on it, it'll be eventually consist.
	if err != nil {
		d.logger.Error(ctx, "Fail to delete from cache, err: ", err.Error())
	}
	d.logger.Info(ctx, "Delete data from cache succeed.")
}

func buildWhereClauseByFilter(filter *cmdt.Filter) (string, []any) {
	whereClause := ""
	args := make([]any, 0)
	if filter.FilterPublish == cmdt.FilterPublish_PUBLISHING {
		whereClause += "status = ? AND "
		args = append(args, 0)
	}
	whereClause += "type = ?"
	if filter.FilterSell == cmdt.FilterSell_SELL {
		args = append(args, 0)
	} else {
		args = append(args, 1)
	}
	return whereClause, args
}
