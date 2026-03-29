package main

import (
	"fmt"
	
	// tea "charm.land/bubbletea/v2"

	"github.com/lariel-o/projects-diary/data"
)

func main() {
	// create the initial folder/file's project structure
	if err := data.CreateProjectDir(); err != nil {
		fmt.Println("Could not create the project folder/file\n	~", err)
		return 
	}

	test := data.ProjectStructModel{
		ProjectName: "Lucas Ariel",
	}

	if err := data.AddNewProject(data.DatabaseInfo.FilesPath["main"], test); err != nil {
		fmt.Println(err)
		return
	}
}

