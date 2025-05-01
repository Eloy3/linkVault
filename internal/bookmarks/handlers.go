package bookmarks

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Eloy3/LinkVault/internal/db"
)

func Pong(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong \n")
}

func Bookmarks(w http.ResponseWriter, req *http.Request) {

	bookmarks, err := db.GetBookmarks()
	if err != nil {
		http.Error(w, "Failed to get bookmarks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookmarks)
}
