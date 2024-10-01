package dao

import (
	"context"
	"github.com/lgangkai/golog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jimoto/account-service/model"
	"testing"
)

func TestAccountDao_GetUserByEmail(t *testing.T) {
	dbm, err := gorm.Open(mysql.Open("root:1234qwer@tcp(127.0.0.1:3306)/account"), &gorm.Config{})
	if err != nil {
		t.Errorf(err.Error())
		//return
	}
	dao := NewAccountDao(&DBMaster{dbm}, golog.Default())
	user, err := dao.GetUserByEmail(context.Background(), "123@qq.com")
	if err != nil {
		return
	}
	t.Log(user)
}

func TestAccountDao_Insert(t *testing.T) {
	dbm, err := gorm.Open(mysql.Open("root:1234qwer@tcp(127.0.0.1:3306)/account"), &gorm.Config{})
	if err != nil {
		t.Errorf(err.Error())
		//return
	}
	dao := NewAccountDao(&DBMaster{dbm}, golog.Default())
	err = dao.Insert(context.Background(), &model.User{
		Email:    "123@qq.com",
		Password: "123456",
		Status:   0,
	})
	if err != nil {
		t.Error(err)
	}
}
