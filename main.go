package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prakash-p-3121/mysqllib"
	"github.com/prakash-p-3121/ordermgtms/cfg"
	"github.com/prakash-p-3121/ordermgtms/controller/ordermgt_controller"
	"github.com/prakash-p-3121/ordermgtms/database"
	"github.com/prakash-p-3121/restlib"
)

func main() {

	msConnectionsMap, err := restlib.CreateMsConnectionCfg("conf/microservice.toml")
	if err != nil {
		panic(err)
	}
	cfg.SetMsConnectionsMap(msConnectionsMap)

	databaseInst, err := mysqllib.CreateDatabaseConnectionWithRetryByCfg("conf/database.toml")
	if err != nil {
		panic(err)
	}
	database.SetSingleStoreConnection(databaseInst)

	hostPortCfg, err := cfg.GetMsConnectionCfg("database-clustermgt-ms")
	if err != nil {
		panic(err)
	}

	connectionsMap, err := mysqllib.CreateShardConnectionsWithRetry(database.GetShardedTableList(), hostPortCfg.Host, hostPortCfg.Port)
	if err != nil {
		panic(err)
	}
	database.SetShardConnectionsMap(connectionsMap)

	router := gin.Default()
	routerGroup := router.Group("/ordermgtms")

	routerGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routerGroup.POST("/v1/order", ordermgt_controller.CreateOrder)
	routerGroup.GET("/v1/order", ordermgt_controller.FindOrderByID)

	err = router.Run("127.0.0.1:3005")
	if err != nil {
		panic("Error Starting ordermgtms")
	}
}
