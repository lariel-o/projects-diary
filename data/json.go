package data

import (
	"encoding/json"
	"os"
)

func readDB(database *WorldStructModel) error {
	data, err := os.ReadFile(DatabaseInfo.FilesPath["main"])
	if err != nil { return err }

	// Read from db and transform it in a  struct
	var db WorldStructModel
	err = json.Unmarshal(data, &db)
	if err != nil { return err }

	*database = db
	return nil
}

func WriteAtDatabase() error {
	b, err := json.Marshal(DB)
	if err != nil { return err }

	if err := os.WriteFile(DatabaseInfo.FilesPath["main"], b, 0666); err != nil { 
		return err
	}

	return nil
}

func AddNewProject(project ProjectStructModel) error {
	// [START] append project to db
	project.ID = uint16(len(DB.World))
	DB.World = append(DB.World, project)
	// [END]

	WriteAtDatabase()

	return nil
}


