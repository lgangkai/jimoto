// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"commodity-service/biz"
	"commodity-service/dao"
	"commodity-service/handler"
	"commodity-service/service/commodity"
	"commodity-service/service/like"
	"github.com/lgangkai/golog"
	"github.com/redis/go-redis/v9"
)

// Injectors from wire.go:

func InitCommodityHandler(dbMaster *dao.DBMaster, dbSlave *dao.DBSlave, client *redis.Client, logger *golog.Logger) *handler.CommodityHandlerImpl {
	commodityDao := dao.NewCommodityDao(dbMaster, dbSlave, client, logger)
	commodityImageDao := dao.NewCommodityImageDao(dbMaster, dbSlave, logger)
	commodityService := commodity.NewCommodityService(commodityDao, commodityImageDao, dbMaster, logger)
	likeDao := dao.NewLikeDao(dbMaster, dbSlave, logger)
	likeService := like.NewLikeService(likeDao, commodityDao, logger)
	commodityBiz := biz.NewCommodityBiz(commodityService, likeService, logger)
	commodityHandlerImpl := handler.NewCommodityHandlerImpl(commodityBiz)
	return commodityHandlerImpl
}
