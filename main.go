package main

import (
	"fmt"
	
	tea "charm.land/bubbletea/v2"
	"github.com/lariel-o/projects-diary/data"
	"github.com/lariel-o/projects-diary/display"
)

func main() {
	// create the initial folder/file's project structure
	if err := data.InitDatas(); err != nil {
		fmt.Println("Could not create the project folder/file\n	~", err)
		return 
	}

	data.AddNewProject(data.ProjectStructModel{
		ProjectName: "Lucas", 
		Description: "Ariel", 
		Time: "93", 
		Failed: false, 
		ID: 0,
	})

	p := tea.NewProgram(display.Daishi{0})
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		return 
	}
}

