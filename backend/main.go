package main

import (
	"backend/config"
	"log"
)

func main() {
	db := config.ConnectDB()
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("Gagal menutup koneksi DB: %v", err)
		}
	}()

	// Lanjutkan logika aplikasi...
}
