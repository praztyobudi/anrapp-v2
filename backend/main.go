package main

import (
	"backend/config"
	"backend/internal/handler"
	repository "backend/internal/repo"
	"backend/internal/usecase"
	"backend/route"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	r := gin.Default()

	repo := repository.NewUserRepository(db)
	uc := usecase.NewUserUsecase(repo)
	auth := handler.NewAuthHandler(uc)
	user := handler.NewUserHandler(uc)

	route.Setup(r, auth, user)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
