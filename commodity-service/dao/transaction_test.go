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

func TestTransaction(t *testing.T) {
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
	m := &DBMaster{dbm}
	ctx := context.Background()
	commodityDao := NewCommodityDao(m, &DBSlave{dbs}, rdb, golog.Default())
	commodityImageDao := NewCommodityImageDao(m, &DBSlave{dbs}, golog.Default())
	m.ExecTx(ctx, func(ctx context.Context) error {
		err = commodityDao.Insert(ctx, &model.Commodity{Detail: "test1"})
		err = commodityImageDao.Insert(ctx, []*model.CommodityImage{{
			CommodityId: 111,
			Image:       "111",
		}})
		return err
	})
}
