package pages

import (
	"github.com/rivo/tview"
)



func createProject() *tview.Form {
	// Create the form
	form := tview.NewForm()


	// Create the field named Project name
	projectNameInput := tview.NewInputField().
		SetLabel("Project name").
		SetFieldWidth(30)
	form.AddFormItem(projectNameInput)


	// Create the field named Description
	descriptionArea := tview.NewTextArea().
		SetLabel("Description")
	form.AddFormItem(descriptionArea)


	// Create the submit button
	form.AddButton("Submit", func() {
	})


	// Create the Quit button
	form.AddButton("Quit", func() {
	})

	return form
}

