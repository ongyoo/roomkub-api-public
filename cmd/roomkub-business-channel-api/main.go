package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/ongyoo/roomkub-api/config"

	//awss3 "github.com/ongyoo/roomkub-api/pkg/awsS3"
	"github.com/ongyoo/roomkub-api/pkg/crypto"
	mongo "github.com/ongyoo/roomkub-api/pkg/database"
	"github.com/ongyoo/roomkub-api/pkg/httpserver"
	"github.com/ongyoo/roomkub-api/pkg/middleware"

	// service
	"github.com/ongyoo/roomkub-api/cmd/roomkub-business-channel-api/businessChannel"
	"github.com/ongyoo/roomkub-api/cmd/roomkub-user-and-customer-api/user/userRolePermission"
)

func getHelp(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "OK")
}

func main() {
	// Config
	conf := config.ReadApi()
	crypto.SetUp(conf.Secrets.WDEKS)

	// Data
	mongoDB := mongo.NewDB()

	// Storage
	//storage := awss3.NewStorage()

	// Repository
	userRolePermissionRepository := userRolePermission.NewRepository(mongoDB)
	businessChannelRepository := businessChannel.NewRepository(mongoDB)

	// Service
	businessChannelService := businessChannel.NewService(businessChannelRepository)
	userRolePermissionService := userRolePermission.NewService(userRolePermissionRepository)

	// Handle
	businessChannelHandler := businessChannel.NewHandler(businessChannelService)

	// Server
	// Server
	server := httpserver.NewServer()
	server.Router.GET("/", getHelp)

	// api v1
	corsMiddleware := middleware.CORS()
	server.Router.Use(corsMiddleware)

	// API Group
	v1 := server.Router.Group("business/api/v1")

	// Middleware
	middlewareService := middleware.NewPermissionService(userRolePermissionService, conf.Root.RootRoleSlug)
	userPermission := middlewareService.ValidatePermission()
	jwtMiddleware := middleware.UserJWT(conf.Secrets.JWTSecret)
	userSecretKey := middleware.ValidateSecret()
	//v1.Use(middleware.HandleError, middleware.HandleEmptyBody, jwtMiddleware, userSecretKey, corsMiddleware, userPermission)
	v1.Use(middleware.HandleError, middleware.HandleEmptyBody, jwtMiddleware, userSecretKey, corsMiddleware, userPermission)

	// Business Channel
	businessChannelGroup := v1.Group("/")
	businessChannel.SetRoutes(businessChannelGroup, businessChannelHandler)

	// Root Business Channel
	rootBusinessChannelGroup := v1.Group("/root")
	businessChannel.SetRootRoutes(rootBusinessChannelGroup, businessChannelHandler)
	// run server
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		server.Router.Run(fmt.Sprintf(":%s", conf.Port))
		wg.Done()
	}()
	wg.Wait()
}
