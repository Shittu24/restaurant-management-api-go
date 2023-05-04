package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shittu24/rms/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controllers.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", controllers.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:orderItem_id", controllers.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItems/createorderItem", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:orderItem_id", controllers.UpdateOrderItem())

}
