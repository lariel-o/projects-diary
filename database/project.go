package database

type project struct {
	name string
	description string

	// Ongoing and Finished tasks
	oTasks []task
	fTasks []task
}


func CreateNewProject(name string, desc string) {
	db.oProjects = append(db.oProjects, project{ 
		name: name,
		description: desc,
	})
}

