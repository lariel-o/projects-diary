package pages

// import (
// 	"fmt"
//
// 	"github.com/lariel-o/projects-diary/internal/database"
//
// 	"github.com/rivo/tview"
// )


// The create project page
// func createProjectPage(p *tview.Pages) *tview.Grid {
// 	form := tview.NewForm()
//
// 	form.SetBorder(true).SetTitle(" Create a new project ")
// 	form.SetBorderPadding(2, 2, 2, 2) 
//
//
// 	form.AddInputField("Project name", "", 90, nil, nil)
// 	form.AddTextArea("Description", "", 90, 15, 0, nil)
// 	form.AddButton("Submit", func() {
// 		nameField := form.GetFormItem(0).(*tview.InputField)
// 		descField := form.GetFormItem(1).(*tview.TextArea)
//
// 		name := nameField.GetText()
// 		desc := descField.GetText()
//
// 		if name == "" || desc == "" {
// 			// janela de erro popup (adicionar futuramente)
// 		} else {
// 			database.CreateNewProject(name, desc)
// 			p.SwitchToPage(fmt.Sprintf("%d", cPROJECT))
// 		}
// 	})
//
// 	form.AddButton("Quit", func() {
// 		p.SwitchToPage(fmt.Sprintf("%d", cPROJECT))
// 	})
//
//
// 	grid := tview.NewGrid().
// 		SetRows(0, 30, 0).      
// 		SetColumns(0, 150, 0).  
// 		AddItem(form, 1, 1, 1, 1, 0, 0, true)
//
// 	return grid
// }
//
