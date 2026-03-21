package main

import (
	"fmt"
	"os"
	
	tea "charm.land/bubbletea/v2"

	"github.com/lariel-o/projects-diary/models"
	aux "github.com/lariel-o/projects-diary/auxiliaries"
)

var userHome, _ = os.UserHomeDir()
var projectPath = userHome + "/.config/projects-diary"

func main() {
	// create the project dir
	if err := aux.CreateProjectDir(projectPath); err != nil {
		fmt.Println("Could not create the project folder\n", err)
		return 
	}

	// write the projects.csv file
	if err := aux.WriteIfNotExist(projectPath + "/projects.csv", "title,finished,description"); err != nil {
		fmt.Println("Unexpected err, not able to write file\n", err)
		return
	}


	// initialize the program
	p := tea.NewProgram(models.WorldModel {
		Cursor: 0,
		Tasks: []string{"Task1", "Task2", "Task3"},
		CurrentView: 0,
	})

	if _, err := p.Run(); err != nil {
		fmt.Println("ERRRRR")
		os.Exit(1)
	}
}

