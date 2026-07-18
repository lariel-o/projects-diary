package database

import (
	"fmt"
	"path/filepath"
	"os"
)

type database struct {
	oProjects []project
	fProjects []project
}

var db database

// InitDatabase make and save the database in the volatile memory
func InitDatabase() error {
	// ########
	// Load the db from non-volatile memory 


	// Verify if the folder exist
	userHome, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Error while trying to get the user home", err)
	}


	folderPath := filepath.Join(userHome, ".config", "projects-diary")
	filePath := filepath.Join(folderPath, "database.json")


	// Verify if the "pojects-diary" folder exists
	if _, err := os.Stat(folderPath); err != nil && os.IsNotExist(err) {
		// Create a new "projects-diary" folder
		if err := os.Mkdir(folderPath, 0755); err != nil {
			return fmt.Errorf("Could not create the project's folder: %w", err)
		}
	} 

	// Verify if the database file exists
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsExist(err) {
			return fmt.Errorf("The file exist but there's something wrong: %w", err)
		}

		// Create the new database.json file
		if err := os.WriteFile(filePath, []byte("{}"), 0644); err != nil {
			return fmt.Errorf("err trying to create database file: %w", err)
		}
	}

	defer file.Close()
	return nil
}

