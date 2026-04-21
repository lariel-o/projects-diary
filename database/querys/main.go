package querys

import (
	"os"
	"encoding/json"

	"github.com/lariel-o/projects-diary/database"
	"github.com/lariel-o/projects-diary/utilities/err"
)

var DB = &database.DB

// write at the non-volatile memory
func WriteAtDatabase() {
	// ### Get the project's database path
	// the user home
	userHome, _ := os.UserHomeDir() 

	// the absolute path to the database
	absPath := userHome + "/.config" + database.PROJECT_DIR

	b, e := json.MarshalIndent(*DB, "", "    ")
	if e != nil { err.LogErr(e) }

	e = os.WriteFile(absPath + "/db.json", b, 0666)
	if e != nil { err.LogErr(e) }
}

