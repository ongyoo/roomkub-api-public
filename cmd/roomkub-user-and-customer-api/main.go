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
	"github.com/ongyoo/roomkub-api/cmd/roomkub-user-and-customer-api/customer"
	"github.com/ongyoo/roomkub-api/cmd/roomkub-user-and-customer-api/user"
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
	userRepository := user.NewRepository(mongoDB)
	userRolePermissionRepository := userRolePermission.NewRepository(mongoDB)
	businessChannelRepository := businessChannel.NewRepository(mongoDB)
	customerRepository := customer.NewRepository(mongoDB)

	// Service
	businessChannelService := businessChannel.NewService(businessChannelRepository)
	userRolePermissionService := userRolePermission.NewService(userRolePermissionRepository)
	userService := user.NewService(businessChannelService, userRolePermissionService, userRepository)
	customerService := customer.NewService(customerRepository)

	// Handle
	userHandler := user.NewHandler(userService)
	userRolePermissionHandler := userRolePermission.NewHandler(userRolePermissionService)
	customerHandler := customer.NewHandler(customerService)

	// Server
	server := httpserver.NewServer()
	server.Router.GET("/", getHelp)

	// api v1
	corsMiddleware := middleware.CORS()
	server.Router.Use(corsMiddleware)

	// API Group
	v1 := server.Router.Group("/authen/api/v1")
	guestV1 := server.Router.Group("/authen/api/v1")

	// Middleware
	middlewareService := middleware.NewPermissionService(userRolePermissionService, conf.Root.RootRoleSlug)
	userPermission := middlewareService.ValidatePermission()
	jwtMiddleware := middleware.UserJWT(conf.Secrets.JWTSecret)
	userSecretKey := middleware.ValidateSecret()
	//v1.Use(middleware.HandleError, middleware.HandleEmptyBody, jwtMiddleware, userSecretKey, corsMiddleware, userPermission)
	v1.Use(middleware.HandleError, middleware.HandleEmptyBody, jwtMiddleware, userSecretKey, corsMiddleware, userPermission)
	// Guest
	guestV1.Use(middleware.HandleError, middleware.HandleEmptyBody, userSecretKey, corsMiddleware)

	// Route
	// Guest
	guestUserGroup := guestV1.Group("/")
	user.SetGuestRoutes(guestUserGroup, userHandler)

	// User
	userGroup := v1.Group("/user")
	user.SetRoutes(userGroup, userHandler)

	// User Root
	rootUserGroup := v1.Group("/root/user")
	user.SetRootRoutes(rootUserGroup, userHandler)

	// Role
	userRolePermissionGroup := v1.Group("/role")
	userRolePermission.SetRoutes(userRolePermissionGroup, userRolePermissionHandler)

	// Root Role
	rootUserRolePermissionGroup := v1.Group("/root/role")
	userRolePermission.SetRootRoutes(rootUserRolePermissionGroup, userRolePermissionHandler)

	// Customer
	customerGroup := v1.Group("/customer")
	customer.SetRoutes(customerGroup, customerHandler)

	rootCustomerGroup := v1.Group("/root/customer")
	customer.SetRootRoutes(rootCustomerGroup, customerHandler)
	// run server
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		server.Router.Run(fmt.Sprintf(":%s", conf.Port))
		wg.Done()
	}()
	wg.Wait()
}
