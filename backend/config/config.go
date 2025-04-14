package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Gagal memuat file .env: %v", err)
	}

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
		user, password, dbname, host, port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Gagal membuka koneksi: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Gagal koneksi ke database: %v", err)
	}

	fmt.Println("Berhasil terkoneksi dengan PostgreSQL!")
	return db
}
