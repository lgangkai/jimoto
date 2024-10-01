package dao

import (
	"commodity-service/dao/model"
	"context"
	"github.com/lgangkai/golog"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestCommodityDao_GetById(t *testing.T) {
	dbm, err := gorm.Open(mysql.Open("root:1234qwer@tcp(127.0.0.1:3306)/commodity"), &gorm.Config{})
	dbs, err := gorm.Open(mysql.Open("root:1234qwer@tcp(127.0.0.1:3306)/commodity"), &gorm.Config{})
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	//err = rdb.Set(context.Background(), "q", "2", 0).Err()
	if err != nil {
		t.Errorf(err.Error())
		//return
	}
	commodityDao := NewCommodityDao(&DBMaster{dbm}, &DBSlave{dbs}, rdb, golog.Default())
	commodity, err := commodityDao.GetById(context.Background(), 1)
	if err != nil {
		t.Errorf("call profileDao.GetProfile failed!")
		return
	}
	if commodity.Id != 1 {
		t.Errorf("got unexpected result!")
		return
	}
}

func TestInsert(t *testing.T) {
	dbm, err := gorm.Open(mysql.Open("root:1234qwer@tcp(127.0.0.1:3306)/commodity"), &gorm.Config{})
	dbs, err := gorm.Open(mysql.Open("root:1234qwer@tcp(127.0.0.1:3306)/commodity"), &gorm.Config{})
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	if err != nil {
		t.Errorf(err.Error())
		//return
	}
	commodityDao := NewCommodityDao(&DBMaster{dbm}, &DBSlave{dbs}, rdb, golog.Default())
	_ = commodityDao.Insert(context.Background(), &model.Commodity{Detail: "fixed"})
}

func TestUpdate(t *testing.T) {
	dbm, err := gorm.Open(mysql.Open("root:1234qwer@tcp(127.0.0.1:3306)/commodity"), &gorm.Config{})
	dbs, err := gorm.Open(mysql.Open("root:1234qwer@tcp(127.0.0.1:3306)/commodity"), &gorm.Config{})
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	if err != nil {
		t.Errorf(err.Error())
		//return
	}
	commodityDao := NewCommodityDao(&DBMaster{dbm}, &DBSlave{dbs}, rdb, golog.Default())
	_ = commodityDao.Update(context.Background(), &model.Commodity{Id: 3, Detail: "fixedddd"})
}

func TestDelete(t *testing.T) {
	dbm, err := gorm.Open(mysql.Open("root:1234qwer@tcp(127.0.0.1:3306)/commodity"), &gorm.Config{})
	dbs, err := gorm.Open(mysql.Open("root:1234qwer@tcp(127.0.0.1:3306)/commodity"), &gorm.Config{})
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	if err != nil {
		t.Errorf(err.Error())
		//return
	}
	commodityDao := NewCommodityDao(&DBMaster{dbm}, &DBSlave{dbs}, rdb, golog.Default())
	_ = commodityDao.Delete(context.Background(), 1)
}

func TestCommodityDao_GetListLatest(t *testing.T) {
	dbm, err := gorm.Open(mysql.Open("root:1234qwer@tcp(127.0.0.1:3306)/commodity"), &gorm.Config{})
	dbs, err := gorm.Open(mysql.Open("root:1234qwer@tcp(127.0.0.1:3306)/commodity"), &gorm.Config{})
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	if err != nil {
		t.Errorf(err.Error())
		//return
	}
	commodityDao := NewCommodityDao(&DBMaster{dbm}, &DBSlave{dbs}, rdb, golog.Default())
	latest, err := commodityDao.GetListLatest(context.Background(), 2, 2)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Log(latest)
}
