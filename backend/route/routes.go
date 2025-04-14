package route

import (
	"backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine, authHandler *handler.AuthHandler) {
	router.POST("/login", authHandler.Login)
}
