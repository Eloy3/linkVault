package main

import (
	"net/http",
	"github.com/eloy3/linkvault/bookmarks/internal/bookmarks"
)

func main() {
	http.HandleFunc("/ping", handlers.pong)
	http.ListenAndServe(":8080", nil)
}
