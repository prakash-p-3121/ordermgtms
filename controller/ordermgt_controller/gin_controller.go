package ordermgt_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prakash-p-3121/restlib"
)

func CreateOrder(ctx *gin.Context) {
	ginRestCtx := restlib.NewGinRestContext(ctx)
	controller := NewOrderController()
	controller.CreateOrder(ginRestCtx)
}

func FindOrderByID(ctx *gin.Context) {
	ginRestCtx := restlib.NewGinRestContext(ctx)
	controller := NewOrderController()
	controller.FindOrderByID(ginRestCtx)
}
