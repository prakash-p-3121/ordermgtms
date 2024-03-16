package impl

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/ordermgtmodel"
	"github.com/prakash-p-3121/ordermgtms/service/ordermgt_service"
	"github.com/prakash-p-3121/restlib"
)

type OrderControllerImpl struct {
	OrderService ordermgt_service.OrderService
}

func (controller *OrderControllerImpl) CreateOrder(restCtx restlib.RestContext) {
	ginRestCtx, ok := restCtx.(*restlib.GinRestContext)
	if !ok {
		internalServerErr := errorlib.NewInternalServerError("Expected GinRestContext")
		internalServerErr.SendRestResponse(ginRestCtx.CtxGet())
		return
	}

	ctx := ginRestCtx.CtxGet()
	var req ordermgtmodel.OrderCreateReq
	err := ctx.BindJSON(&req)
	if err != nil {
		badReqErr := errorlib.NewBadReqError("payload-serialization" + err.Error())
		badReqErr.SendRestResponse(ctx)
		return
	}

	idResp, appErr := controller.OrderService.CreateOrder(&req)
	if appErr != nil {
		appErr.SendRestResponse(ctx)
		return
	}
	restlib.OkResponse(ctx, idResp)
}

func (controller *OrderControllerImpl) FindOrderByID(restCtx restlib.RestContext) {
	ginRestCtx, ok := restCtx.(*restlib.GinRestContext)
	if !ok {
		internalServerErr := errorlib.NewInternalServerError("Expected GinRestContext")
		internalServerErr.SendRestResponse(ginRestCtx.CtxGet())
		return
	}
	ctx := ginRestCtx.CtxGet()
	orderID := ctx.Query("order-id")
	if restlib.TrimAndCheckForEmptyString(&orderID) {
		badReqErr := errorlib.NewBadReqError("order-id-empty")
		badReqErr.SendRestResponse(ctx)
		return
	}

	orderPtr, appErr := controller.OrderService.FindOrderByID(orderID)
	if appErr != nil {
		appErr.SendRestResponse(ctx)
		return
	}
	restlib.OkResponse(ctx, orderPtr)
}
