package ordermgt_service

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/ordermgtmodel"
)

type OrderService interface {
	CreateOrder(req *ordermgtmodel.OrderCreateReq) errorlib.AppError
	FindOrderByID(orderID string) (*ordermgtmodel.Order, errorlib.AppError)
}
