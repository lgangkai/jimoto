package dao

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lgangkai/golog"
	"github.com/redis/go-redis/v9"
	"jimoto/account-service/model"
	"math/rand"
	"time"
)

const (
	REDIS_KEY_GET_PROFILE_PREFIX           = "userinfo:get_profile:"
	REDIS_KEY_GET_PROFILE_EXPIRE_BASE      = time.Second * 60
	REDIS_KEY_GET_PROFILE_EXPIRE_MAX_SHIFT = 30
)

type ProfileDao struct {
	dbMaster *DBMaster
	dbSlave  *DBSlave
	dbRedis  *redis.Client
	logger   *golog.Logger
}

func NewProfileDao(dbMaster *DBMaster, dbSlave *DBSlave, dbRedis *redis.Client, logger *golog.Logger) *ProfileDao {
	return &ProfileDao{
		dbMaster: dbMaster,
		dbSlave:  dbSlave,
		dbRedis:  dbRedis,
		logger:   logger,
	}
}

func (d *ProfileDao) GetProfileByUserId(ctx context.Context, userId uint64) (*model.Profile, error) {
	d.logger.Info(ctx, "Call ProfileDao.GetProfile.")
	profile := &model.Profile{}

	// 1. try to get value from redis first.
	rKey := fmt.Sprintf("%v%d", REDIS_KEY_GET_PROFILE_PREFIX, userId)
	profileStr, err := d.dbRedis.Get(ctx, rKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			d.logger.Info(ctx, "Can not find in cache, go to sql DB.")
		} else {
			d.logger.Error(ctx, "Can not get from cache, err: ", err.Error(), ". Go to sql DB")
		}
	} else {
		d.logger.Info(ctx, "Get profile json from cache, profile: ", profileStr)
		err = json.Unmarshal([]byte(profileStr), profile)
		if err != nil {
			d.logger.Error(ctx, "json.Unmarshal failed, err: ", err.Error(), ". Go to sql DB")
		} else {
			d.logger.Info(ctx, "Get profile from cache succeeded.")
			return profile, nil
		}
	}

	// 2. get value from mysql-slave if not found in redis.
	if err = d.dbSlave.Where("user_id = ? AND is_deleted = ?", userId, false).Take(profile).Error; err != nil {
		d.logger.Error(ctx, "Fail to get data, err: ", err.Error())
		return nil, err
	}
	d.logger.Info(ctx, "Get profile done, profile: ", profile)

	// 3. write profile as json string back to cache.
	pBytes, err := json.Marshal(profile)
	if err != nil {
		d.logger.Error(ctx, "json.Marshal failed, err: ", err.Error(), ". It will not be saved to cache.")
		return profile, nil
	}
	// set key expiration time as base time plus random time to avoid cache avalanche.
	randExp := time.Duration(rand.Intn(REDIS_KEY_GET_PROFILE_EXPIRE_MAX_SHIFT)) * time.Second
	err = d.dbRedis.Set(ctx, rKey, string(pBytes), REDIS_KEY_GET_PROFILE_EXPIRE_BASE+randExp).Err()
	if err != nil {
		d.logger.Error(ctx, "redis set failed, err: ", err.Error(), ". It will not be saved to cache.")
		return profile, nil
	}
	d.logger.Info(ctx, "Write back to redis done. key: ", rKey, ", value: ", string(pBytes),
		", expiration: ", REDIS_KEY_GET_PROFILE_EXPIRE_BASE+randExp)

	return profile, nil
}

func (d *ProfileDao) Update(ctx context.Context, profile *model.Profile) error {
	// use cache aside pattern to update DB and then delete from cache.
	// 1. update data to mysql-master.
	d.logger.Info(ctx, "Call ProfileDao.Update.")
	if err := d.dbMaster.db(ctx).Where("user_id = ? ", profile.UserId).UpdateColumns(profile).Error; err != nil {
		d.logger.Error(ctx, "Fail to update to sql DB, err: ", err.Error())
		return err
	}
	d.logger.Info(ctx, "Update data to sql DB succeed.")

	// 2. delete data from redis.
	d.deleteFromCache(ctx, profile.UserId)

	return nil
}

func (d *ProfileDao) Delete(ctx context.Context, userId uint64) error {
	// use cache aside pattern to delete from DB and then delete from cache.
	// 1. delete data from mysql-master.
	// not really deleted, set is_deleted field to 1.
	d.logger.Info(ctx, "Call ProfileDao.Delete.")
	//if err := d.dbMaster.Where("userId = ?", userId).Delete(&model.Commodity{}).Error; err != nil {
	//	d.logger.Error(ctx, "Fail to delete from sql DB, err: ", err.Error())
	//	return err
	//}
	if err := d.dbMaster.db(ctx).Save(&model.Profile{UserId: userId, IsDeleted: true}).Error; err != nil {
		d.logger.Error(ctx, "Fail to delete from sql DB, err: ", err.Error())
		return err
	}
	d.logger.Info(ctx, "Delete data from sql DB succeed.")

	// 2. delete data from redis.
	d.deleteFromCache(ctx, userId)

	return nil
}

func (d *ProfileDao) Insert(ctx context.Context, profile *model.Profile) error {
	// we don't operate cache in insert. Cache data will be load when read.
	d.logger.Info(ctx, "Call ProfileDao.Insert, profile: ", profile)

	if err := d.dbMaster.db(ctx).Create(profile).Error; err != nil {
		d.logger.Error(ctx, "Fail to insert into sql DB, err: ", err.Error())
		return err
	}
	d.logger.Info(ctx, "Insert data into sql DB succeed.")

	return nil
}

func (d *ProfileDao) deleteFromCache(ctx context.Context, id uint64) {
	rKey := fmt.Sprintf("%v%d", REDIS_KEY_GET_PROFILE_PREFIX, id)
	err := d.dbRedis.Del(ctx, rKey).Err()
	// if delete failed, other process might read the dirty data.
	// since we have set an expiration time on it, it'll be eventually consist.
	if err != nil {
		d.logger.Error(ctx, "Fail to delete from cache, err: ", err.Error())
	}
	d.logger.Info(ctx, "Delete data from cache succeed.")
}
