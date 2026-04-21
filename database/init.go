package database

import (
	"os"
	"encoding/json"
	"github.com/lariel-o/projects-diary/utilities/err"
)

const PROJECT_DIR = "/projects-diary"

// ### LOAD EMPTY DATABASE MODEL IN MEMORY
var DB = database{}

func Init() {
	// ### Get the project's database path
	// the user home
	userHome, _ := os.UserHomeDir() 

	// the absolute path to the database
	absPath := userHome + "/.config" + PROJECT_DIR

	// ### Create and set things about database
	// create the project's folder if it doesn't exist 
	e := os.Mkdir(absPath, 0750); 
	if e != nil && !os.IsExist(e) { err.LogErr(e) }

	// create the db file (db.json) if not exist and set
	// the initial configuration
	_, e = os.ReadFile(absPath + "/db.json")
	if e != nil && !os.IsExist(e) {
		INIT_DB := []byte(`{"OProjects":{},"FProjects":{},"OProjectsCount":0,"FProjectsCount":0,"ID":0}`)
		os.WriteFile(absPath + "/db.json", INIT_DB, 0666)
	} else if e != nil { err.LogErr(e) }

	// ### Put datas from non-volatile memory to the volatile database
	// read from database
	data, e := os.ReadFile(absPath + "/db.json")
	if e != nil { err.LogErr(e) }

	// Call json pkg to convert it in a proper database-struct-type
	e = json.Unmarshal(data, &DB)
	if e != nil { err.LogErr(e) }
}

