package database

type database struct {
	oProjects []project
	fProjects []project
}
var db database

// InitDatabase make and save the database in the volatile memory
func InitDatabase() {
	// Load the db from non-volatile memory
}

