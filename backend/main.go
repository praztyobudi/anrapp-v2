package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Menampilkan "Hello, World!" di browser
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Menyeting route dan handler
	http.HandleFunc("/", handler)

	// Menjalankan server di port 8080
	fmt.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
