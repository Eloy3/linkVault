package server

import (
    "log"
    "net/http"
	"github.com/Eloy3/LinkVault/internal/db"
	"github.com/Eloy3/LinkVault/internal/bookmarks"
)

func StartServer(addr string) error {
	// Initialize the database
	err := db.Init()
	if err != nil {
		log.Fatalf("failed to initialize DB: %v", err)
	}
	defer db.DB.Close()

	// Initialize the server
	mux := http.NewServeMux()
	bookmarks.Router(mux)
	return http.ListenAndServe(addr, mux)
}
