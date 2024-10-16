package handler

import (
	"context"
	"jimoto/order-service/biz"
	"protos/order"
)

type OrderHandlerImpl struct {
	orderBiz *biz.OrderBiz
}

func NewOrderHandlerImpl(orderBiz *biz.OrderBiz) *OrderHandlerImpl {
	return &OrderHandlerImpl{orderBiz: orderBiz}
}

func (h *OrderHandlerImpl) CreateOrder(ctx context.Context, in *order.CreateOrderRequest, out *order.CreateOrderResponse) error {
	return h.CreateOrder(getTraceContext(ctx, in.GetRequestId(), 0), in, out)
}

func getTraceContext(ctx context.Context, requestId string, orderId uint64) context.Context {
	return context.WithValue(ctx, "traceKey", map[string]any{
		"request_id": requestId,
		"order_id":   orderId,
	})
}
