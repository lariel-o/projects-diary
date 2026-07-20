package pages

import (
	"fmt"

	"github.com/rivo/tview"
)

func createProjectPage(createNewProject func(string, string) error ) *tview.Grid {
	form := tview.NewForm()

	form.SetBorder(true).SetTitle(" Create a new project ")
	form.SetBorderPadding(2, 2, 2, 2) 


	form.AddInputField("Project name", "", 90, nil, nil)
	form.AddTextArea("Description", "", 90, 15, 0, nil)
	form.AddButton("Submit", func() {
		nameField := form.GetFormItem(0).(*tview.InputField)
		descField := form.GetFormItem(1).(*tview.TextArea)

		name := nameField.GetText()
		desc := descField.GetText()

		if name == "" || desc == "" {
			// janela de erro popup (adicionar futuramente)
		} else {
			createNewProject(name, desc)
			switchToPage(fmt.Sprintf("%d", eProjectPage))
		}
	})

	form.AddButton("Quit", func() {
		switchToPage(fmt.Sprintf("%d", eProjectPage))
	})


	grid := tview.NewGrid().
		SetRows(0, 30, 0).      
		SetColumns(0, 150, 0).  
		AddItem(form, 1, 1, 1, 1, 0, 0, true)

	return grid
}
