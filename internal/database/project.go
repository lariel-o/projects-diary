package database

type project struct {
	Name string
	Description string

	// Ongoing and Finished tasks
	OTasks []task
	FTasks []task
}


func CreateNewProject(name string, desc string) error {
	db.OProjects = append(db.OProjects, project{ 
		Name: name,
		Description: desc,
	})

	if err := saveToNVMemory(); err != nil {
		return err
	}

	return nil
}

