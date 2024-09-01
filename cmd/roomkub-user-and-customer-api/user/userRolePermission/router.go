package userRolePermission

import (
	"github.com/gin-gonic/gin"
)

func SetRoutes(group *gin.RouterGroup, handler Handler, mdlws ...gin.HandlerFunc) {
	// Get
	group.GET("/list/:business_id/:page/:limit", handler.GetUserRoleAllList)
	group.GET("/:business_id/:id", handler.GetUserRoleByID)
	// group.GET("/:id/:business_id/:page/:limit", handler.GetUserRoleByIDs)
	group.GET("/permission/list/:business_id/:page/:limit", handler.GetUserPermissionAllList)
	//group.GET("/permission/:business_id/:id", handler.GetUserPermissionByID)
	group.GET("/permission/:business_id/:id/:page/:limit", handler.GetUserPermissionByIDs)
	// Create
	group.POST("/create", handler.CreateUserRole)
	group.POST("/permission/create", handler.CreateUserPermission)
	// PUT
	group.PUT("/update/info", handler.UpdateUserRoleInfo)
	group.PUT("/update/permission", handler.UpdateUserRolePermission)
	group.PUT("/permission/update/info", handler.UpdateUserPermissionInfo)
	// DELETE
	group.DELETE("/delete", handler.DeleteUserRolePermission)
	group.DELETE("/permission/delete", handler.DeleteUserPermission)
}

func SetRootRoutes(group *gin.RouterGroup, handler Handler, mdlws ...gin.HandlerFunc) {
	// Get
	group.GET("/list/:business_id/:page/:limit/:active/:type", handler.GetRootUserRoleAllList)
	group.GET("/permission/list/:business_id/:page/:limit/:active/:type", handler.GetRootUserPermissionAllList)
}
