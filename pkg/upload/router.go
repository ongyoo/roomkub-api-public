package upload

import (
	"github.com/gin-gonic/gin"
)

func SetRoutes(group *gin.RouterGroup, handler Handler, mdlws ...gin.HandlerFunc) {
	//group.POST("/:key", handler.Upload)
	group.GET("private/:key/:id", handler.GetPrivate)
}

func SetPublicRoutes(group *gin.RouterGroup, handler Handler, mdlws ...gin.HandlerFunc) {
	group.GET("/:id", handler.GetPubilc)
	//group.GET("/:key/:id", handler.Get)
}
