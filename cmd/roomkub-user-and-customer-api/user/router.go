package user

import (
	"github.com/gin-gonic/gin"
)

func SetRoutes(group *gin.RouterGroup, handler Handler, mdlws ...gin.HandlerFunc) {
	// Get
	group.GET("/users/:page/:limit/:channel_id", handler.GetUserAllListByBusinessChannel)
	// PUT
	// update profile
	group.PUT("/update/profile", handler.UpdateProfileInfo)
	group.PUT("/update/my-profile", handler.UpdateMyProfileInfo)
	// update profile image
	group.PUT("/update/profile/image", handler.UpdateProfileImage)
	group.PUT("/update/my-profile/image", handler.UpdateMyProfileImage)
	// update password
	group.PUT("/update/password", handler.UpdatePassword)
	group.PUT("/update/my-password", handler.UpdateMyPassword)
	// update ban user
	group.PUT("/update/ban", handler.BanUser)
	group.PUT("/update/unban", handler.UnBanUser)
	// update verify user
	group.PUT("/update/verify", handler.VerifyUser)
}

func SetGuestRoutes(group *gin.RouterGroup, handler Handler, mdlws ...gin.HandlerFunc) {
	group.POST("/register", handler.CreateUser)
	group.POST("/login", handler.LoginUser)
}

func SetRootRoutes(group *gin.RouterGroup, handler Handler, mdlws ...gin.HandlerFunc) {
	// Get
	group.GET("/users/:page/:limit", handler.GetUserAllList)
}
