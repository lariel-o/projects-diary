package pages

import (
	"fmt"

	"github.com/lariel-o/projects-diary/internal/database"

	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
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
	projectPage := projectPage()
	pagesCollection.AddPage(fmt.Sprintf("%d", cPROJECT),
		projectPage,
		true,
		true)


	// Add the "Create Project Page"
	pagesCollection.AddPage(fmt.Sprintf("%d", cPROJECT_CREATE),
		createProject(),
		true,
		false)




	// Keys logical
	lastEvent := projectPage.GetInputCapture()



	// ---- Project page logic
	projectPage.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'c':
			pagesCollection.SwitchToPage(fmt.Sprintf("%d", cPROJECT_CREATE))
		}

		if lastEvent != nil {
			return lastEvent(event)
		}


		return nil
	})

	database.CreateNewProject("Test", "Some random description")

	return pagesCollection
}

