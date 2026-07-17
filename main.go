package main

import (
	"github.com/lariel-o/projects-diary/internal/pages"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	p := pages.Dashi(app)

	if err := app.SetRoot(p, true).SetFocus(p).Run(); err != nil {
		panic(err)
	}
}

