package data

import "os"

var userHome, _ = os.UserHomeDir()
var DatabasePath = userHome + "/.config/projects-diary"

type TaskStructModel struct {
	Status string	 // Tell about how the task is going
	Content string   // Tell about what is the tesk objective
	Time string      // Tell about the task's time limit (optional)
	Failed bool      // Tell about if the task has run out of time
	ID uint16
}

type ProjectStructModel struct {
	ProjectName string
	Description string
	Time string 
	Failed bool 
	Tasks [] TaskStructModel // List all the tasks
	ID uint16
}

type WorldStructModel struct {
	World []ProjectStructModel
}

type databaseInfo struct  {
	mainDBPath string
	allFilesPath []string
}

var DatabaseInfo = databaseInfo {
	mainDBPath: DatabasePath + "/main.json",
	allFilesPath: []string {
		DatabasePath + "/main.json",
	},
}

