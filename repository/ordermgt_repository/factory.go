package ordermgt_repository

import (
	"github.com/prakash-p-3121/ordermgtms/database"
	"github.com/prakash-p-3121/ordermgtms/repository/ordermgt_repository/impl"
)

func NewOrderRepository() OrderRepository {
	return &impl.OrderRepositoryImpl{
		SingleStoreConnection: database.GetSingleStoreConnection(),
		ShardConnectionsMap:   database.GetShardConnectionsMap(),
	}
}
