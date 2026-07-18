package pages

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/lariel-o/projects-diary/database"
)

const (
	cPROJECT int = iota
	cTASK
	cPROJECT_CREATE
	cPROJECT_TAKS
)

func Dashi(app *tview.Application) *tview.Pages {
	// create the pages logical
	pagesCollection := tview.NewPages()	

	// Add the project page
	pagesCollection.AddPage(fmt.Sprintf("%d", cPROJECT),
		project(),
		true,
		true)

	database.CreateNewProject("Lucas", "Essa é uma descrição muito interessante")

	return pagesCollection
}

