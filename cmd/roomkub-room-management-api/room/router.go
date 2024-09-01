package room

import (
	"github.com/gin-gonic/gin"
)

func SetRoutes(group *gin.RouterGroup, handler Handler, mdlws ...gin.HandlerFunc) {
	// POST
	group.POST("/create", handler.CreateRoom)
	// PUT
	group.PUT("/info/update", handler.UpdateInfo)
	group.PUT("/thumbnail/upload", handler.UpdateThumbnail)
	group.PUT("/image/upload", handler.UpdateImage)
	group.PUT("/publisher/update", handler.UpdatePublisher)
	group.PUT("/unpublisher/update", handler.UpdateUnPublisher)
	group.PUT("/active/update", handler.UpdateActive)
	group.PUT("/unactive/update", handler.UpdateUnActive)
	// GET
	group.GET("/:page/:limit", handler.GetRoomAllListByBusinessId)
}

func SetRootRoutes(group *gin.RouterGroup, handler Handler, mdlws ...gin.HandlerFunc) {
	// Get
	group.GET("/:page/:limit", handler.GetRootRoomAllList)
}
