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
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("Gagal menutup koneksi DB: %v", err)
		}
	}()
	defer db.Close()

	r := gin.Default()

	repo := repository.NewUserRepository(db)
	uc := usecase.NewUserUsecase(repo)
	h := handler.NewAuthHandler(uc)

	route.Setup(r, h)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
	// Lanjutkan logika aplikasi...
}
