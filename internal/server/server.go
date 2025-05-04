package server

import (
	"log"
	"net/http"

	"github.com/Eloy3/LinkVault/internal/bookmarks"
	"github.com/Eloy3/LinkVault/internal/db"
)

func StartServer(addr string) error {
	// Initialize the database
	err := db.Init()
	if err != nil {
		log.Fatalf("failed to initialize DB: %v", err)
	}

	// Initialize the server
	mux := http.NewServeMux()
	bookmarks.Router(mux)
	return http.ListenAndServe(addr, mux)
}
