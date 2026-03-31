package data

import (
	"os"
	auxs "github.com/lariel-o/projects-diary/auxiliaries"
	"fmt"
)

func CreateProjectDir() error {
	// try to create the project folder
	err := os.Mkdir(DatabasePath, 0750)	
	if err != nil && !os.IsExist(err) {
		return err
	}

	// create all files asked to
	for fileType, filePath := range DatabaseInfo.FilesPath {
		switch fileType {
		case "main":
			if err := auxs.WriteIfNotExist(filePath, "{\"World\": []}"); err != nil {
				return err
			}
		}
	}

	return nil
}

// Initialize database in memory
var DB = WorldStructModel{}

func InitDatas() error {
	// create the project dir to save the database
	if err := CreateProjectDir(); err != nil { return err }

	// fill the DB with datas
	err := readDB(&DB)
	if err != nil { return err }
	fmt.Println(DB)

	return nil
}


