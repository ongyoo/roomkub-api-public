package businessChannel

import (
	"github.com/gin-gonic/gin"
)

func SetRoutes(group *gin.RouterGroup, handler Handler, mdlws ...gin.HandlerFunc) {
	// Create
	group.POST("/create", handler.CreateBusinessChannel)
	// Get
	group.GET("/my-channel/list/:page/:limit", handler.GetBusinessChannelListByUserId)
	// Put
	group.PUT("/delete", handler.UpdateDelete)
	group.PUT("/my-channel/update-info", handler.UpdateSetting)
}

func SetRootRoutes(group *gin.RouterGroup, handler Handler, mdlws ...gin.HandlerFunc) {
	// Get
	group.GET("/all/:page/:limit", handler.GetBusinessChannelAllList)
	// Put
	group.PUT("/ban", handler.UpdateBan)
	group.PUT("/unban", handler.UpdateUnBan)
}
