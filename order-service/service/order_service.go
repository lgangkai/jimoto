package service

import (
	"github.com/lgangkai/golog"
	"jimoto/order-service/dao"
)

type OrderService struct {
	orderDao *dao.OrderDao
	logger   *golog.Logger
}

func NewOrderService(orderDao *dao.OrderDao, logger *golog.Logger) *OrderService {
	return &OrderService{
		orderDao: orderDao,
		logger:   logger,
	}
}

func (s *OrderService) CreateOrder() error {
	return nil
}
