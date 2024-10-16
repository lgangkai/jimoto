// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: order/order.proto

package order

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Order service

func NewOrderEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Order service

type OrderService interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...client.CallOption) (*CreateOrderResponse, error)
}

type orderService struct {
	c    client.Client
	name string
}

func NewOrderService(name string, c client.Client) OrderService {
	return &orderService{
		c:    c,
		name: name,
	}
}

func (c *orderService) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...client.CallOption) (*CreateOrderResponse, error) {
	req := c.c.NewRequest(c.name, "Order.CreateOrder", in)
	out := new(CreateOrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Order service

type OrderHandler interface {
	CreateOrder(context.Context, *CreateOrderRequest, *CreateOrderResponse) error
}

func RegisterOrderHandler(s server.Server, hdlr OrderHandler, opts ...server.HandlerOption) error {
	type order interface {
		CreateOrder(ctx context.Context, in *CreateOrderRequest, out *CreateOrderResponse) error
	}
	type Order struct {
		order
	}
	h := &orderHandler{hdlr}
	return s.Handle(s.NewHandler(&Order{h}, opts...))
}

type orderHandler struct {
	OrderHandler
}

func (h *orderHandler) CreateOrder(ctx context.Context, in *CreateOrderRequest, out *CreateOrderResponse) error {
	return h.OrderHandler.CreateOrder(ctx, in, out)
}
