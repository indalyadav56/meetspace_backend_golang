package main

import (
	"fmt"
	authHandlers "meetspace_backend/auth/handlers"
	authRepo "meetspace_backend/auth/repositories"
	authRoutes "meetspace_backend/auth/routes"
	authServices "meetspace_backend/auth/services"
	chatHandlers "meetspace_backend/chat/handlers"
	chatRepo "meetspace_backend/chat/repositories"
	chatRoutes "meetspace_backend/chat/routes"
	chatServices "meetspace_backend/chat/services"
	websocketRoute "meetspace_backend/chat/websocket"
	clientRepo "meetspace_backend/client/repositories"
	clientRoutes "meetspace_backend/client/routes"
	clientServices "meetspace_backend/client/services"
	"meetspace_backend/config"
	"meetspace_backend/middlewares"
	userHandlers "meetspace_backend/user/handlers"
	userRepo "meetspace_backend/user/repositories"
	userRoutes "meetspace_backend/user/routes"
	userServices "meetspace_backend/user/services"
	"net/http"

	docs "meetspace_backend/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.basic BasicAuth

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
func main() {
	// load environment
	config.LoadEnv()
	
	// initialize database connection
	db := config.InitDB()

	// repositories
	userRepo := userRepo.NewUserRepository(db)
	verificationRepo := authRepo.NewVerificationRepository(db)
	clientRepo := clientRepo.NewClientRepository(db)
	chatMessageRepo := chatRepo.NewChatMessageRepository(db)
	chatRoomRepo := chatRepo.NewChatRoomRepository(db)

	// services
	userService := userServices.NewUserService(userRepo)
	authService := authServices.NewAuthService(userService)
	verificationService := authServices.NewVerificationService(verificationRepo)
	clientServices.NewClientService(clientRepo, userService)
	clientServices.NewClientUserService(clientRepo, userService)
	chatRoomService := chatServices.NewChatRoomService(chatRoomRepo, userService)
	chatServices.NewChatGroupService(chatRoomRepo, userService)
	chatServices.NewChatMessageService(chatMessageRepo, userService, chatRoomService)

	// handlers
	authHandler := authHandlers.NewAuthHandler(authService, verificationService)
	userHandler := userHandlers.NewUserHandler(userService)
	chatRoomHandler := chatHandlers.NewChatRoomHandler()
	chatGroupHandler := chatHandlers.NewChatGroupHandler()
	chatMessageHandler := chatHandlers.NewChatMessageHandler()
	

	r := gin.Default()

	// static
	r.StaticFS("/uploads", http.Dir("./uploads"))
	
	// middlewares
	r.Use(middlewares.LoggerMiddleware())
	r.Use(middlewares.CorsMiddleware())
	r.Use(middlewares.AuthMiddleware())
	
	// routes
	authRoutes.AuthRouter(r, authHandler)
	userRoutes.UserRouter(r, userHandler)
	chatRoutes.ChatRouter(r, chatRoutes.ChatHandlers{
		ChatRoomHandler: chatRoomHandler,
		ChatGroupHandler: chatGroupHandler, 
		ChatMessageHandler: chatMessageHandler,
	})
	websocketRoute.WebSocketRouter(r)
	clientRoutes.ClientRouter(r)

	// swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1)),
	)
	
	docs.SwaggerInfo.Title = "MeetSpace API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	fmt.Println("server:->", "http://localhost:8080")
	r.Run()
}
