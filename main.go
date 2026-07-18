package main

import (
	"github.com/lariel-o/projects-diary/internal/pages"
	"github.com/lariel-o/projects-diary/database"

	"github.com/rivo/tview"
)

func main() {
	// Init the database
	database.InitDatabase()

	app := tview.NewApplication()

	p := pages.Dashi(app)

	if err := app.SetRoot(p, true).SetFocus(p).Run(); err != nil {
		panic(err)
	}
}

