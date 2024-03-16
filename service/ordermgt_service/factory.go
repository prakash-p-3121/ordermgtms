package ordermgt_service

import (
	"github.com/prakash-p-3121/ordermgtms/repository/ordermgt_repository"
	"github.com/prakash-p-3121/ordermgtms/service/ordermgt_service/impl"
)

func NewOrderService() OrderService {
	repository := ordermgt_repository.NewOrderRepository()
	return &impl.OrderServiceImpl{OrderRepository: repository}
}
