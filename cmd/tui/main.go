package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/Eloy3/LinkVault/internal/models"
	"github.com/rivo/tview"
)

func main() {

	requestURL := "http://localhost:8080/bookmarks"
	res, err := http.Get(requestURL)
	if err != nil {
		log.Fatalf("failed to request %v. \n %v", requestURL, err)
		os.Exit(1)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
	}

	app := tview.NewApplication()

	var bookmarks []models.Bookmark
	err = json.Unmarshal(body, &bookmarks)
	if err != nil {
		log.Fatalf("failed to decode json response: %v", err)
	}
	var list = tview.NewList()

	for i, bookmark := range bookmarks {
		list.AddItem(bookmark.Title, bookmark.URL, rune('a'+i), nil)
		i++
	}

	list.AddItem("Quit", "", '0', func() {
		app.Stop()
	})

	list.SetSelectedFunc(func(index int, mainText, secondaryText string, shortcut rune) {

		if shortcut != '0' {
			openInBrowser(secondaryText)
		}
	})

	if err := app.SetRoot(list, true).Run(); err != nil {
		panic(err)
	}

}

func openInBrowser(url string) {
	cmd := exec.Command("cmd.exe", "/C", "start", url)
	_ = cmd.Start()
}
