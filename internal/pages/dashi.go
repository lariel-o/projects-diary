package pages

import (
	"fmt"

	"github.com/lariel-o/projects-diary/internal/database"

	"github.com/rivo/tview"
	// "github.com/gdamore/tcell/v2"
)

const (
	eProjectPage int = iota
	eCreateProjectPage
)

var	pages = tview.NewPages()	

func switchToPage(p string) {
	pages.SwitchToPage(p)
}

func Dashi(app *tview.Application) *tview.Pages {
	layout, _, _ := projectPageDinamic( database.GetProjectsInfo(true) )


	pages.AddPage(fmt.Sprintf("%d", eProjectPage),
		layout,
		true,
		true)

	return pages
}

