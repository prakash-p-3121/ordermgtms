package ordermgt_controller

import (
	"github.com/prakash-p-3121/ordermgtms/controller/ordermgt_controller/impl"
	"github.com/prakash-p-3121/ordermgtms/service/ordermgt_service"
)

func NewOrderController() OrderController {
	service := ordermgt_service.NewOrderService()
	return &impl.OrderControllerImpl{OrderService: service}
}
