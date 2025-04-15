package route

import (
	"backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine, authHandler *handler.AuthHandler, userHandler *handler.UserHandler) {
	router.POST("/login", authHandler.Login)
	router.POST("/register", authHandler.Register)
	router.GET("/users", userHandler.GetAll)
	router.PUT("/users", userHandler.Update)
	router.DELETE("/users/:id", userHandler.Delete)
}
