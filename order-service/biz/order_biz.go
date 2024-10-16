package biz

import (
	"context"
	"github.com/lgangkai/golog"
	"jimoto/order-service/service"
	"protos/order"
)

type OrderBiz struct {
	orderService *service.OrderService
	logger       *golog.Logger
}

func NewOrderBiz(orderService *service.OrderService, logger *golog.Logger) *OrderBiz {
	return &OrderBiz{
		orderService: orderService,
		logger:       logger,
	}
}

func (b *OrderBiz) CreateOrder(ctx context.Context, in *order.CreateOrderRequest, out *order.CreateOrderResponse) error {
	return nil
}
