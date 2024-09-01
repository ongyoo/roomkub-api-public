package roomContract

import (
	"github.com/gin-gonic/gin"
)

func SetRoutes(group *gin.RouterGroup, handler Handler, mdlws ...gin.HandlerFunc) {
	// Create
	group.POST("/create", handler.CreateRoomContract)
	// Get
	group.GET("/list/:room_id/:page/:limit", handler.GetRoomContractList)
	// Put
	group.PUT("/update", handler.UpdateRoomContract)
}

func SetRootRoutes(group *gin.RouterGroup, handler Handler, mdlws ...gin.HandlerFunc) {
	// Get
	group.GET("/list/:room_id/:page/:limit", handler.GetRootRoomContractList)
}
