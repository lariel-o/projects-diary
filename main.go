package main

import (
	"github.com/lariel-o/projects-diary/internal/pages"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	g := pages.ProjectPage(app)

	if err := app.SetRoot(g, true).SetFocus(g).Run(); err != nil {
		panic(err)
	}
}

