package data

import (
	"encoding/json"
	"slices"
	"os"
)

func AddNewProject(path string, project ProjectStructModel) error {
	// checck if the name already exist at the project
	data, err := os.ReadFile(DatabaseInfo.FilesPath["optmized"])
	if err != nil { return err }

	var optmizing OptmizedStructModel
	if err = json.Unmarshal(data, &optmizing); err != nil { return err }
	

	// return an err if the name already exist
	if slices.Index(optmizing.ProjectsNames, project.ProjectName) != -1 { return err }


	// read the file at path
	data, err = os.ReadFile(path)
	if err != nil { return err }

	// Read from db and transform it in a  struct
	var db WorldStructModel
	err = json.Unmarshal(data, &db)
	if err != nil { return err }

	// append the value to the db json-like
	project.ID = optmizing.CurrentProjectID
	db.World = append(db.World, project)

	// transform it in json
	b, err := json.Marshal(db)
	if err != nil { return err }

	// write at db
	if err := os.WriteFile(path, b, 0666); err != nil { return err }



	optmizing.ProjectsNames = append(optmizing.ProjectsNames, project.ProjectName)
	optmizing.CurrentProjectID += 1

	// Make a json of the optmized file content
	b, err = json.Marshal(optmizing)
	if err != nil { return err }

	// add the new project title at ProjectsNames array field
	if err := os.WriteFile(DatabaseInfo.FilesPath["optmized"], b, 0666); err != nil { return err }

	return nil
}

