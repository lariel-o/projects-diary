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
	if err := auxs.WriteIfNotExist(auxs.ProjectsFile.Path, auxs.ProjectsFile.Header); err != nil {
		fmt.Println("Unexpected err, not able to write file\n", err)
		return
	}

	// get the projects infos
	projectData, err := auxs.GetCSV(auxs.ProjectsFile.Path)
	if err != nil {
		fmt.Println("Unexpected err\n", projectData)
	}

	// initialize the program
	p := tea.NewProgram(models.WorldModel {
		ProjectsTitle: projectData["title"],
		ProjectsDescription: projectData["description"],

		ShowingDescription: 0,
		IsShowingDescription: false,

		Cursor: 0,
		ViewControl: 0,
	})

	if _, err := p.Run(); err != nil {
		fmt.Println("ERRRRR")
		os.Exit(1)
	}
}

