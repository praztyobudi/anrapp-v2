package route

import (
	"backend/internal/handler"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine, authHandler *handler.AuthHandler, userHandler *handler.UserHandler) {
	router.POST("/login", authHandler.Login)
	router.POST("/register", authHandler.Register)
	// Protected routes (require token)
	auth := router.Group("/", middleware.JWTMiddleware())
	{
		auth.POST("/refresh", authHandler.Refresh)
		auth.GET("/users", userHandler.GetAll)
		auth.PUT("/users", userHandler.Update)
		auth.DELETE("/users/delete", userHandler.Delete)
	}
}
