package main

import (
	"fmt"
	"os"
	
	tea "charm.land/bubbletea/v2"

	"github.com/lariel-o/projects-diary/models"
)

var userHome, _ = os.UserHomeDir()
var projectPath = userHome + "/.config/projects-diary"

func main() {
	if err := createProjectDir(projectPath); err != nil {
		fmt.Println("Could not create the project folder\n", err)
		return 
	}

	if err := writeIfNotExist(projectPath + "/projects.csv", "title,finished,description"); err != nil {
		fmt.Println("Unexpected err, not able to write file\n", err)
		return
	}

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

func createProjectDir(projectPath string) (error) {
	err := os.Mkdir(projectPath, 0750)	
	if err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}

func writeIfNotExist(filePath string, toWrite string) (error) {
	_, openFileErr := os.OpenFile(filePath, 0, 0644)

	if openFileErr != nil && !os.IsExist(openFileErr) {
		writeErr := os.WriteFile(filePath, []byte(toWrite), 0666)
		if writeErr != nil {
			return writeErr
		}
	} else if os.IsExist(openFileErr) {
		return openFileErr
	}

	return nil
}

