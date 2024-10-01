package dao

import (
	"commodity-service/dao/model"
	"context"
	"github.com/lgangkai/golog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestLikeDao_Query(t *testing.T) {
	dbm, err := gorm.Open(mysql.Open("root:1234qwer@tcp(127.0.0.1:3306)/commodity"), &gorm.Config{})
	dbs, err := gorm.Open(mysql.Open("root:1234qwer@tcp(127.0.0.1:3306)/commodity"), &gorm.Config{})
	if err != nil {
		t.Errorf(err.Error())
		//return
	}
	likeDao := NewLikeDao(&DBMaster{dbm}, &DBSlave{dbs}, golog.Default())
	rst, err := likeDao.Query(context.Background(), &model.Like{
		CommodityId: 1,
		UserId:      1,
	})
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(rst)
}

func TestLikeDao_Count(t *testing.T) {
	dbm, err := gorm.Open(mysql.Open("root:1234qwer@tcp(127.0.0.1:3306)/commodity"), &gorm.Config{})
	dbs, err := gorm.Open(mysql.Open("root:1234qwer@tcp(127.0.0.1:3306)/commodity"), &gorm.Config{})
	if err != nil {
		t.Errorf(err.Error())
		//return
	}
	likeDao := NewLikeDao(&DBMaster{dbm}, &DBSlave{dbs}, golog.Default())
	cnt, err := likeDao.Count(context.Background(), &model.Like{
		UserId: 1,
	})
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(cnt)
}

func TestLikeDao_Delete(t *testing.T) {
	dbm, err := gorm.Open(mysql.Open("root:1234qwer@tcp(127.0.0.1:3306)/commodity"), &gorm.Config{})
	dbs, err := gorm.Open(mysql.Open("root:1234qwer@tcp(127.0.0.1:3306)/commodity"), &gorm.Config{})
	if err != nil {
		t.Errorf(err.Error())
		//return
	}
	likeDao := NewLikeDao(&DBMaster{dbm}, &DBSlave{dbs}, golog.Default())
	err = likeDao.Delete(context.Background(), &model.Like{
		CommodityId: 1,
		UserId:      1,
	})
	if err != nil {
		t.Log(err.Error())
		return
	}
}
