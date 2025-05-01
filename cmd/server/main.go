package main

import (
	"log"
	"net/http"

	"github.com/Eloy3/LinkVault/internal/bookmarks"

	"github.com/Eloy3/LinkVault/internal/db"
)

func main() {
	err := db.Init()
	if err != nil {
		log.Fatalf("failed to initialize DB: %v", err)
	}
	defer db.DB.Close()

	mux := http.NewServeMux()
	bookmarks.Router(mux)

	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
