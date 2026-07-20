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
	// The logic that trace the pages change and what it need to do when it happens
	// It is used just when the dinamic idea is too complex to be done
	trace := func(pageName string) {
		switch pageName {
		// case Project Page
		case fmt.Sprintf("%d", eProjectPage):
			database.GetProjectsInfo(true)	
		}
	}

	projectPLayout, _, _ := projectPageDinamic(database.GetProjectsInfo(true), trace)
	createProjectPLayout := createProjectPage(database.CreateNewProject)

	pages.AddPage(fmt.Sprintf("%d", eProjectPage),
		projectPLayout,
		true,
		true)

	pages.AddPage(fmt.Sprintf("%d", eCreateProjectPage),
		createProjectPLayout,
		true,
		false)



	return pages
}

