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


	// ########################## Adding the pages
	// Add the "Project" page
	projectGrid, projectList, projectDescription := projectPage()


	// Redraw always when the pages are changed
	pagesCollection.SetChangedFunc(func() {

		currentPage, _ := pagesCollection.GetFrontPage() 
		switch currentPage {
		case fmt.Sprintf("%d", cPROJECT):
			// get the infos for the Project Page
			projectInfos := database.GetProjectNames(true)
			if len(*projectInfos) == 0 {
				for _, j := range *projectInfos {
					projectList.AddItem(j.Name, "", '*', nil)
				}

				// add the matched description
				projectList.SetChangedFunc(func(i int, _, _ string, _ rune) {
					projectDescription.SetText((*projectInfos)[i].Description)
				})
			}
		}

		// redraw
		go app.Draw()
	})







	pagesCollection.AddPage(fmt.Sprintf("%d", cPROJECT),
		projectGrid,
		true,
		true)

	// Add the "Create Project" Page
	pagesCollection.AddPage(fmt.Sprintf("%d", cPROJECT_CREATE),
		createProjectPage(pagesCollection),
		true,
		false)

	
	// Keys logical
	lastEvent := projectGrid.GetInputCapture()



	// ---- Project page logic
	projectGrid.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'c':
			pagesCollection.SwitchToPage(fmt.Sprintf("%d", cPROJECT_CREATE))
		}

		if lastEvent != nil {
			return lastEvent(event)
		}


		return nil
	})

	return pagesCollection
}

