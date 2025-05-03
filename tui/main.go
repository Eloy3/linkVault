package main

import (
	"os/exec"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	list := tview.NewList().
		AddItem("LinkVault Home", "https://linkvault.io", '1', nil).
		AddItem("Example Bookmark", "https://example.com", '2', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})

	list.SetSelectedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		// Only open if not the Quit item
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
