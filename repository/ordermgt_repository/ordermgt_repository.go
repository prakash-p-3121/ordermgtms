package ordermgt_repository

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/ordermgtmodel"
)

type OrderRepository interface {
	CreateOrder(shardID int64, idGenResp *idgenmodel.IDGenResp, req *ordermgtmodel.OrderCreateReq) errorlib.AppError
	FindOrderByID(shardID int64, orderID string) (*ordermgtmodel.Order, errorlib.AppError)
}
