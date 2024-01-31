package routes

import (
	"meetspace_backend/user/handlers"

	"github.com/gin-gonic/gin"
)


func UserRouter(e *gin.Engine, handler *handlers.UserHandler){
	userRouter := e.Group("/v1/user")
	usersRouter := e.Group("/v1/users")

	usersRouter.POST("", handler.CreateUserHandler)
	usersRouter.GET("", handlers.GetAllUsers)
	usersRouter.GET("/:userId", handlers.GetUserByID)
	usersRouter.PATCH("", handlers.UpdateUser)

	userRouter.GET("/check-email", handlers.CheckUserEmail)
}