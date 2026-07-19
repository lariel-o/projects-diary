package pages

import (
	"github.com/rivo/tview"
)


// The create project page
func createProjectPage() *tview.Grid {
	form := tview.NewForm()

	form.SetBorder(true).SetTitle(" Create a new project ")
	form.SetBorderPadding(2, 2, 2, 2) 

	
	projectNameInput := tview.NewInputField().
		SetLabel("Project name").
		SetFieldWidth(90)



	form.AddFormItem(projectNameInput)
	form.AddTextArea("Description", "", 90, 15, 0, nil)
	form.AddButton("Submit", func() {})
	form.AddButton("Quit", func() {})


	grid := tview.NewGrid().
		SetRows(0, 30, 0).      
		SetColumns(0, 150, 0).  
		AddItem(form, 1, 1, 1, 1, 0, 0, true)

	return grid
}

