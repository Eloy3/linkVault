package bookmarks

import (
	"fmt"
	"net/http"
)

func pong(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong \n")
}
