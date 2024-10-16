//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/lgangkai/golog"
	"github.com/redis/go-redis/v9"
	"jimoto/order-service/biz"
	"jimoto/order-service/dao"
	"jimoto/order-service/handler"
	"jimoto/order-service/service"
)

func InitOrderHandler(*dao.DBMaster, *dao.DBSlave, *redis.Client, *golog.Logger) *handler.OrderHandlerImpl {
	wire.Build(dao.NewOrderDao, service.NewOrderService, biz.NewOrderBiz, handler.NewOrderHandlerImpl)
	return &handler.OrderHandlerImpl{}
}
