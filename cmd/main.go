package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ongyoo/roomkub-api/config"

	awss3 "github.com/ongyoo/roomkub-api/pkg/awsS3"
	"github.com/ongyoo/roomkub-api/pkg/crypto"
	mongo "github.com/ongyoo/roomkub-api/pkg/database"
	"github.com/ongyoo/roomkub-api/pkg/httpserver"
	"github.com/ongyoo/roomkub-api/pkg/middleware"
	"github.com/ongyoo/roomkub-api/pkg/upload"

	// service
	"github.com/ongyoo/roomkub-api/cmd/roomkub-business-channel-api/businessChannel"
	roomContract "github.com/ongyoo/roomkub-api/cmd/roomkub-room-management-api/contract"
	contractTemplate "github.com/ongyoo/roomkub-api/cmd/roomkub-room-management-api/contract/template"
	"github.com/ongyoo/roomkub-api/cmd/roomkub-room-management-api/room"
	"github.com/ongyoo/roomkub-api/cmd/roomkub-user-and-customer-api/customer"
	"github.com/ongyoo/roomkub-api/cmd/roomkub-user-and-customer-api/user"
	"github.com/ongyoo/roomkub-api/cmd/roomkub-user-and-customer-api/user/userRolePermission"
)

func getHelp(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "OK")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func main() {
	// Config
	conf := config.ReadApi()
	crypto.SetUp(conf.Secrets.WDEKS)

	// Data
	mongoDB := mongo.NewDB()

	// Storage
	storage := awss3.NewStorage()

	// Repository
	userRepository := user.NewRepository(mongoDB)
	userRolePermissionRepository := userRolePermission.NewRepository(mongoDB)
	businessChannelRepository := businessChannel.NewRepository(mongoDB)
	customerRepository := customer.NewRepository(mongoDB)
	roomRepository := room.NewRepository(mongoDB)
	contractTemplateRepository := contractTemplate.NewRepository(mongoDB)
	roomContractRepository := roomContract.NewRepository(mongoDB)

	// Service
	uploadService := upload.NewService(storage)
	businessChannelService := businessChannel.NewService(businessChannelRepository)
	userRolePermissionService := userRolePermission.NewService(userRolePermissionRepository)
	userService := user.NewService(businessChannelService, userRolePermissionService, userRepository)
	customerService := customer.NewService(customerRepository)
	roomService := room.NewService(roomRepository, uploadService)
	contractTemplateService := contractTemplate.NewService(contractTemplateRepository)
	roomContractService := roomContract.NewService(roomContractRepository, customerService, contractTemplateService)

	// Handle
	userHandler := user.NewHandler(userService)
	userRolePermissionHandler := userRolePermission.NewHandler(userRolePermissionService)
	customerHandler := customer.NewHandler(customerService)
	businessChannelHandler := businessChannel.NewHandler(businessChannelService)
	roomHandler := room.NewHandler(roomService)
	uploadHandler := upload.NewHandler(uploadService)
	contractTemplateHandler := contractTemplate.NewHandler(contractTemplateService)
	contractHandler := roomContract.NewHandler(roomContractService)

	// Server
	server := httpserver.NewServer()
	server.Router.GET("/", getHelp)

	// api v1
	corsMiddleware := middleware.CORS()
	server.Router.Use(corsMiddleware)
	server.Router.Use(CORSMiddleware())
	// CORS configuration
	server.Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*", "http://localhost:3000"}, // Replace with your allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Define your routes
	server.Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello Kub",
		})
	})

	// API Group
	v1 := server.Router.Group("/api/v1")
	publicV1 := server.Router.Group("/api/v1")
	guestV1 := server.Router.Group("/api/v1")

	// Middleware
	middlewareService := middleware.NewPermissionService(userRolePermissionService, conf.Root.RootRoleSlug)
	userPermission := middlewareService.ValidatePermission()
	jwtMiddleware := middleware.UserJWT(conf.Secrets.JWTSecret)
	userSecretKey := middleware.ValidateSecret()
	//v1.Use(middleware.HandleError, middleware.HandleEmptyBody, jwtMiddleware, userSecretKey, corsMiddleware, userPermission)
	v1.Use(middleware.HandleError, middleware.HandleEmptyBody, jwtMiddleware, userSecretKey, corsMiddleware, userPermission)
	publicV1.Use(middleware.HandleError, middleware.HandleEmptyBody, corsMiddleware)
	// Guest
	guestV1.Use(middleware.HandleError, middleware.HandleEmptyBody, userSecretKey, corsMiddleware)

	// Route
	// Guest
	guestUserGroup := guestV1.Group("/authen")
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

	// Business Channel
	businessChannelGroup := v1.Group("/business")
	businessChannel.SetRoutes(businessChannelGroup, businessChannelHandler)

	// Root Business Channel
	rootBusinessChannelGroup := v1.Group("/root/business")
	businessChannel.SetRootRoutes(rootBusinessChannelGroup, businessChannelHandler)

	// Room
	roomGroup := v1.Group("/room")
	room.SetRoutes(roomGroup, roomHandler)

	// Root Room
	rootRoomGroup := v1.Group("/root/room")
	room.SetRootRoutes(rootRoomGroup, roomHandler)

	// contract template
	contractTemplateGroup := v1.Group("/contract-template")
	contractTemplate.SetRoutes(contractTemplateGroup, contractTemplateHandler)

	// Root Contract template
	rootContractTemplateGroup := v1.Group("/root/contract-template")
	contractTemplate.SetRootRoutes(rootContractTemplateGroup, contractTemplateHandler)

	// room contract
	roomContractGroup := v1.Group("/contract")
	roomContract.SetRoutes(roomContractGroup, contractHandler)

	// Root Room contract
	rootRoomContractGroup := v1.Group("/root/contract")
	roomContract.SetRootRoutes(rootRoomContractGroup, contractHandler)

	// Upload
	uploadGroup := v1.Group("/upload")
	upload.SetRoutes(uploadGroup, uploadHandler)

	publicUploadGroup := publicV1.Group("/upload/public")
	upload.SetPublicRoutes(publicUploadGroup, uploadHandler)

	// run server
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		server.Router.Run(fmt.Sprintf(":%s", conf.Port))
		wg.Done()
	}()
	wg.Wait()
}
