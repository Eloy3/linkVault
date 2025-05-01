package bookmarks

import (
	"net/http"
)

func Router(mux *http.ServeMux) {
	mux.HandleFunc("/ping", Pong)
	mux.HandleFunc("/bookmarks", Bookmarks)
}
