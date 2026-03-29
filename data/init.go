package data

import (
	"os"
	auxs "github.com/lariel-o/projects-diary/auxiliaries"
)

func CreateProjectDir() (error) {
	// try to create the project folder
	err := os.Mkdir(DatabasePath, 0750)	
	if err != nil && !os.IsExist(err) {
		return err
	}

	for _, filePath := range DatabaseInfo.allFilesPath {
		if err := auxs.WriteIfNotExist(filePath, "{\"World\": []}"); err != nil {
			return err
		}
	}

	return nil
}

