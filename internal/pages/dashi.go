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
	projectPLayout, _, _, projectPChangeHandler := projectPageDinamic(database.GetProjectsInfo(true))
	createProjectPLayout := createProjectPage(database.CreateNewProject)

	pages.AddPage(fmt.Sprintf("%d", eProjectPage),
		projectPLayout,
		true,
		true)

	pages.AddPage(fmt.Sprintf("%d", eCreateProjectPage),
		createProjectPLayout,
		true,
		false)

	// The logic that trace the pages change and what it need to do when it happens
	// It is used just when the dinamic idea is too complex to be done
	pages.SetChangedFunc(func() {
		currentPage, _ := pages.GetFrontPage()

		switch currentPage {
		// case Project Page
		case fmt.Sprintf("%d", eProjectPage):
			database.GetProjectsInfo(true)	
			projectPChangeHandler()
		}
	})



	return pages
}

