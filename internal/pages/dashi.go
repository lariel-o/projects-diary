package pages

import (
	"fmt"

	"github.com/lariel-o/projects-diary/internal/database"

	"github.com/rivo/tview"
)

const (
	eProjectPage int = iota
	eCreateProjectPage
	eWarningPage 
)

var (
	pages = tview.NewPages()
)

func switchToPage(p string) {
	pages.SwitchToPage(p)
}

func callWarningPage(message string, app *tview.Application) {
	onClose := func() {
		pages.RemovePage(fmt.Sprintf("%d", eWarningPage))
	}

	warning := warningPage(message, app, onClose)

	pages.AddPage(fmt.Sprintf("%d", eWarningPage), warning, true, false)

	switchToPage(fmt.Sprintf("%d", eWarningPage))
}


func Dashi(app *tview.Application) *tview.Pages {
	projectPLayout, _, _, projectPChangeHandler := projectPageDinamic(database.GetProjectsInfo(true))
	createProjectPLayout := createProjectPage(app, database.CreateNewProject)

	pages.AddPage(fmt.Sprintf("%d", eProjectPage),
		projectPLayout,
		true,
		true)

	pages.AddPage(fmt.Sprintf("%d", eCreateProjectPage),
		createProjectPLayout,
		true,
		false)

	pages.SetChangedFunc(func() {
		currentPage, _ := pages.GetFrontPage()
		switch currentPage {
		case fmt.Sprintf("%d", eProjectPage):
			database.GetProjectsInfo(true)
			projectPChangeHandler()
		}
	})

	return pages
}
