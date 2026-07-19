package main

import (
	"fmt"
	"os"

	"github.com/lariel-o/projects-diary/internal/pages"
	"github.com/lariel-o/projects-diary/internal/database"

	"github.com/rivo/tview"
)

func main() {
	// Init the database
	if err := database.InitDatabase(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	app := tview.NewApplication()
	p := pages.Dashi(app)

	if err := app.SetRoot(p, true).SetFocus(p).Run(); err != nil {
		panic(err)
	}
}

