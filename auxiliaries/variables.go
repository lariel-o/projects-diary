package auxiliaries

import "os"

type CSVFileInfo struct {
	Path string
	Header string
}

var userHome, _ = os.UserHomeDir()
var ProjectPath = userHome + "/.config/projects-diary"

var ProjectsFile = CSVFileInfo {
	Path: ProjectPath + "/projects.csv",
	Header: "title,finished,description",
}

var DefaultProjectCSV = [2]string {ProjectPath + "/projects.csv", 
	"title,finished,description"}

