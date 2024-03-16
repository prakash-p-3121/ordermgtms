package impl

import (
	database_clustermgt_client "github.com/prakash-p-3121/database-clustermgt-client"
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenclient"
	"github.com/prakash-p-3121/ordermgtmodel"
	"github.com/prakash-p-3121/ordermgtms/cfg"
	"github.com/prakash-p-3121/ordermgtms/database"
	"github.com/prakash-p-3121/ordermgtms/repository/ordermgt_repository"
)

type OrderServiceImpl struct {
	OrderRepository ordermgt_repository.OrderRepository
}

func (service *OrderServiceImpl) CreateOrder(req *ordermgtmodel.OrderCreateReq) errorlib.AppError {
	appErr := req.Validate()
	if appErr != nil {
		return appErr
	}
	idGenMSCfg, err := cfg.GetMsConnectionCfg("idgenms")
	if err != nil {
		return errorlib.NewInternalServerError(err.Error())
	}
	idGenClient := idgenclient.NewIDGenClient(idGenMSCfg.Host, uint(idGenMSCfg.Port))
	resp, appErr := idGenClient.NextID(database.OrdersTable)
	if appErr != nil {
		return appErr
	}
	orderID := resp.ID

	databaseClstrMgtMsCfg, err := cfg.GetMsConnectionCfg("database-clustermgt-ms")
	if err != nil {
		return errorlib.NewInternalServerError(err.Error())
	}
	client := database_clustermgt_client.NewDatabaseClusterMgtClient(databaseClstrMgtMsCfg.Host, uint(databaseClstrMgtMsCfg.Port))
	shardPtr, appErr := client.FindShard(database.OrdersTable, orderID)
	if appErr != nil {
		return appErr
	}

	appErr = service.OrderRepository.CreateOrder(*shardPtr.ID, resp, req)
	if appErr != nil {
		return appErr
	}
	return nil
}

func (service *OrderServiceImpl) FindOrderByID(orderID string) (*ordermgtmodel.Order, errorlib.AppError) {

	databaseClstrMgtMsCfg, err := cfg.GetMsConnectionCfg("database-clustermgt-ms")
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	client := database_clustermgt_client.NewDatabaseClusterMgtClient(databaseClstrMgtMsCfg.Host, uint(databaseClstrMgtMsCfg.Port))
	shardPtr, appErr := client.FindShard(database.OrdersTable, orderID)
	if appErr != nil {
		return nil, appErr
	}

	orderPtr, appErr := service.OrderRepository.FindOrderByID(*shardPtr.ID, orderID)
	if appErr != nil {
		return orderPtr, appErr
	}
	return orderPtr, nil
}
