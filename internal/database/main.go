package database

import (
	"fmt"
	"path/filepath"
	"os"
	"encoding/json"
)

type database struct {
	OProjects []project
	FProjects []project
}


var home string
var folderPath string
var filePath string


var db database



// Read from the database in the non-volatile memory
func readFromNVMemory() error {
	// Try to read the file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open database file: %w", err)
	}
	defer file.Close()


	// Decode the information from the .json to the db variable in the volatile memory
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&db); err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}

	return nil
}



// Save at the database in the non-volatile memory
func saveToNVMemory() error {
	data, err := json.MarshalIndent(db, "", "	")

	if err != nil {
		return fmt.Errorf("failed to marshal database: %w", err)
	}

	// Try to write to the database in the NV memory
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write database file: %w", err)
	}
	return nil
}



// InitDatabase make and save the database in the volatile memory
func InitDatabase() error {
	// Verify if the folder exist
	userHome, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Error while trying to get the user home", err)
	}


	folderPath = filepath.Join(userHome, ".config", "projects-diary")
	filePath = filepath.Join(folderPath, "database.json")


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

	// Load the db from the non-volatile memory
	readFromNVMemory()

	defer file.Close()
	return nil
}

