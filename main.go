package main

import (
	"fmt"
	"os"
	
	tea "charm.land/bubbletea/v2"

	"github.com/lariel-o/projects-diary/models"
	auxs "github.com/lariel-o/projects-diary/auxiliaries"
)

func main() {
	// create the project dir
	if err := auxs.CreateProjectDir(auxs.ProjectPath); err != nil {
		fmt.Println("Could not create the project folder\n", err)
		return 
	}

	// write the projects.csv file
	if err := auxs.WriteIfNotExist(auxs.DefaultProjectCSV[0], auxs.DefaultProjectCSV[1]); err != nil {
		fmt.Println("Unexpected err, not able to write file\n", err)
		return
	}

	auxs.GetCSV(auxs.DefaultProjectCSV[0])

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

