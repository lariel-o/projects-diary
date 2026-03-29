package data

import (
	"encoding/json"
	"fmt"
	"os"
)

func AddNewProject(path string, projectModel ProjectStructModel) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var db WorldStructModel

	// Read from db and transform it in a  struct
	err = json.Unmarshal(data, &db)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// append the value at db
	db.World = append(db.World, projectModel)

	// transform it in json
	b, err := json.Marshal(db)
	if err != nil {
		return err
	}

	if err := os.WriteFile(path, b, 0666); err != nil {
		return err
	}


	return nil
}

