package database

type project struct {
	Name string
	Description string

	// Ongoing and Finished tasks
	OTasks []task
	FTasks []task
}

type ReturnableProjectInfos struct {
	Name string
	Description string
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

func GetProjectNames(isOngoing bool) []ReturnableProjectInfos {
	readFromNVMemory()

	var interactor *([]project)
	if isOngoing {
		interactor = &(db.OProjects)
	} else {
		interactor = &(db.FProjects)
	}

	arr := make([]ReturnableProjectInfos, len(*interactor))

	for i := range *interactor {
		arr[i].Name = (*interactor)[i].Name
		arr[i].Description = (*interactor)[i].Description

	}

	return arr
}

