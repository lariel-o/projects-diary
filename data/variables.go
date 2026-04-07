package data

import (
	"os"
	"time"
)

var userHome, _ = os.UserHomeDir()
var DatabasePath = userHome + "/.config/projects-diary"

// START ~ Database struct in general
type TaskStructModel struct {
	Content string   // Tell about what is the tesk objective
	Status string	 // Tell about how the task is going
	Time string      // Tell about the task's time limit (optional)
	Failed bool      // Tell about if the task has run out of time
	ID uint16
}

type ProjectStructModel struct {
	ProjectName string
	Description string

	CreatedAt time.Time
	ExpireAt time.Time
	HaveExpireTime bool

	Tasks [] TaskStructModel // List all the tasks

	ID uint16

	TasksCount uint16
	LastTaskID uint16
}

type WorldStructModel struct {
	World []ProjectStructModel
	ProjectsCount uint16
	LastProjectID uint16
}

type databaseInfo struct  {
	FilesPath map[string] string
}
// END

var DatabaseInfo = databaseInfo {
	FilesPath: map[string]string {
		"main": DatabasePath + "/main.json",
	},
}
