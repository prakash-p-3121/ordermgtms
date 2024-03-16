package ordermgt_controller

import "github.com/prakash-p-3121/restlib"

type OrderController interface {
	CreateOrder(restCtx restlib.RestContext)
	FindOrderByID(restCtx restlib.RestContext)
	//UpdateOrderStateByID(restCtx restlib.RestContext)
}
