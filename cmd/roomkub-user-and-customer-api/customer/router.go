package customer

import (
	"github.com/gin-gonic/gin"
)

func SetRoutes(group *gin.RouterGroup, handler Handler, mdlws ...gin.HandlerFunc) {
	// POST
	group.POST("/create", handler.CreateCustomer)
	// PUT
	group.PUT("/profile/update", handler.UpdateCustomerProfileInfo)
	// GET
	group.GET("/find/:phone", handler.FindCustomerByPhone)
}

func SetRootRoutes(group *gin.RouterGroup, handler Handler, mdlws ...gin.HandlerFunc) {
	// Get
	group.GET("/:page/:limit", handler.GetCustomerAllList)
}
