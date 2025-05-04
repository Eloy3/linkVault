package main

import (
	"os/exec"

	"github.com/rivo/tview"
	"github.com/Eloy3/LinkVault/internal/db"
)

func main() {
	
	go err := server.StartServer(":8080")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

	app := tview.NewApplication()

	list := tview.NewList().
		AddItem("LinkVault Home", "https://linkvault.io", '1', nil).
		AddItem("Example Bookmark", "https://example.com", '2', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})

	list.SetSelectedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		
		if shortcut != 'q' {
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
