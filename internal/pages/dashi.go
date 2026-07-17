package pages

import (
	"fmt"

	"github.com/rivo/tview"
	// "github.com/gdamore/tcell/v2"
)

const (
	cPROJECT int = iota
	cTASK
	cPROJECT_CREATE
	cPROJECT_TAKS
)

func Dashi(app *tview.Application) *tview.Pages {
	pagesCollection := tview.NewPages()	

	pagesCollection.AddPage(fmt.Sprintf("%d", cPROJECT),
		project(),
		true,
		true)

	return pagesCollection
}

