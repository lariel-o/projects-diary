package database

import (
	"os"
	"github.com/lariel-o/projects-diary/utilities/err"
)

const PROJECT_DIR = "/projects-diary"

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
	os.WriteFile(absPath + "/db.json", []byte("{\"projects\": []}"), 0666)
}

