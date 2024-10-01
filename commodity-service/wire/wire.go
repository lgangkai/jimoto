//go:build wireinject
// +build wireinject

package wire

import (
	commodity2 "commodity-service/biz"
	"commodity-service/dao"
	"commodity-service/handler"
	"commodity-service/service/commodity"
	"commodity-service/service/like"
	"github.com/google/wire"
	"github.com/lgangkai/golog"
	"github.com/redis/go-redis/v9"
)

func InitCommodityHandler(*dao.DBMaster, *dao.DBSlave, *redis.Client, *golog.Logger) *handler.CommodityHandlerImpl {
	wire.Build(dao.NewCommodityDao, dao.NewLikeDao, like.NewLikeService, dao.NewCommodityImageDao, commodity.NewCommodityService, commodity2.NewCommodityBiz, handler.NewCommodityHandlerImpl)
	return &handler.CommodityHandlerImpl{}
}
