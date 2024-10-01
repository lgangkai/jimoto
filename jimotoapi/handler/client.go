package handler

import (
	"context"
	"github.com/lgangkai/golog"
	"jimotoapi/conf"
	"protos/account"
	"protos/commodity"
)

type Client struct {
	context         context.Context
	commodityClient commodity.CommodityService
	accountClient   account.AccountService
	config          *conf.Config
	logger          *golog.Logger
}

func NewClient(context context.Context, commodityClient commodity.CommodityService, accountClient account.AccountService, config *conf.Config, logger *golog.Logger) *Client {
	return &Client{
		context:         context,
		commodityClient: commodityClient,
		accountClient:   accountClient,
		config:          config,
		logger:          logger,
	}
}
