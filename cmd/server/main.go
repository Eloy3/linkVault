package main

import (
	"log"

	"github.com/Eloy3/LinkVault/internal/server"
)

func main() {
	err := server.StartServer(":8080")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	} else {
		log.Println("Server started on :8080")
	}
}
