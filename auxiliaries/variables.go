package auxiliaries

import "os"

var userHome, _ = os.UserHomeDir()
var ProjectPath = userHome + "/.config/projects-diary"
var DefaultProjectCSV = [2]string {ProjectPath + "/projects.csv", 
	"title,finished,description"}

