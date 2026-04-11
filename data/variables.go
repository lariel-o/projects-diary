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

	CreatedAt time.Time
	ExpireAt time.Time
	HaveExpireTime bool

	Finished bool

	ID uint16
}

type ProjectStructModel struct {
	ProjectName string
	Description string

	CreatedAt time.Time
	ExpireAt time.Time
	HaveExpireTime bool

	Tasks [] TaskStructModel // List all the tasks

	Finished bool

	ID uint16

	GTasksCount uint16  // ongoing tasks count
	FTasksCount uint16  // finished tasks count
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
