//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/lgangkai/golog"
	"github.com/redis/go-redis/v9"
	"jimoto/account-service/biz"
	"jimoto/account-service/dao"
	"jimoto/account-service/handler"
	"jimoto/account-service/service"
)

func InitAccountHandler(*dao.DBMaster, *dao.DBSlave, *redis.Client, *golog.Logger) *handler.AccountHandlerImpl {
	wire.Build(dao.NewAccountDao, biz.NewProfileBiz, service.NewAccountService, biz.NewAccountBiz, handler.NewAccountHandlerImpl)
	return &handler.AccountHandlerImpl{}
}
