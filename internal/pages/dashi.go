package pages

import (
	"fmt"

	// "github.com/lariel-o/projects-diary/internal/database"

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
	layout, list, textView := projectPageDinamic()

	// NAO ESQUECA DE APAGAR ESSA LINHA
	fmt.Sprintf("%d %d %d", layout, list, textView)


	pages.AddPage(fmt.Sprintf("%d", eProjectPage),
		layout,
		true,
		true)

	return pages
}

