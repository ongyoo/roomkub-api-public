package contractTemplate

import (
	"github.com/gin-gonic/gin"
)

func SetRoutes(group *gin.RouterGroup, handler Handler, mdlws ...gin.HandlerFunc) {
	// Create
	group.POST("/create", handler.CreateContractTemplate)
	// Get
	group.GET("/my-list/:page/:limit", handler.GetMyContractTemplateList)
	group.GET("/publisher-list/:page/:limit", handler.GetAllPublisherContractTemplateList)
	group.GET("/detail/:id", handler.GetContractTemplateDetail)
	// // Put
	group.PUT("/update", handler.UpdateContract)
	group.PUT("/update/active", handler.UpdateActive)
	group.PUT("/update/publisher", handler.UpdatePublisher)
}

func SetRootRoutes(group *gin.RouterGroup, handler Handler, mdlws ...gin.HandlerFunc) {
	group.GET("/list/:page/:limit", handler.GetAllPublisherContractTemplateList)
	group.GET("/detail/:id", handler.GetContractTemplateDetail)
}
