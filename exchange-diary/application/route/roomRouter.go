package route

import (
	"github.com/ExchangeDiary_Server/exchange-diary/application/controller"
	"github.com/gin-gonic/gin"
)

func RoomRoutes(incomingRoutes *gin.Engine, controller controller.RoomController) {
	rooms := incomingRoutes.Group("v1/rooms")
	{
		rooms.GET("/", controller.GetAll())
		rooms.GET("/:room_id", controller.Get())
		rooms.POST("/", controller.Post())
	}
}