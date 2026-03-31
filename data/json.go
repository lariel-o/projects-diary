package data

import (
	"encoding/json"
	"os"
)

func AddNewProject(project ProjectStructModel) error {
	// read the file at path
	data, err := os.ReadFile(DatabaseInfo.FilesPath["main"])
	if err != nil { return err }

	// Read from db and transform it in a  struct
	var db WorldStructModel
	err = json.Unmarshal(data, &db)
	if err != nil { return err }

	// [START] append project to db
	project.ID = uint16(len(db.World))
	db.World = append(db.World, project)
	// [END]

	// transform it in json
	b, err := json.Marshal(db)
	if err != nil { return err }


	if err := os.WriteFile(DatabaseInfo.FilesPath["main"], b, 0666); err != nil { 
		return err 
	}

	return nil
}



