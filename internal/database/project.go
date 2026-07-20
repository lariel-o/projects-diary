package database

import (
)

type project struct {
	Name string
	Description string

	// Ongoing and Finished tasks
	OTasks []task
	FTasks []task
}

type ReturnableProjectsInfo struct {
	Name string
	Description string
}

// A pointer to this variable is returned when need to interact with the projects
// info dinamicly
var returnableProjectsInfo []ReturnableProjectsInfo



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

func GetProjectsInfo(isOngoing bool) *[]ReturnableProjectsInfo {
	readFromNVMemory()

	// Free the slice
	returnableProjectsInfo = nil


	var interactor *([]project)
	if isOngoing {
		interactor = &(db.OProjects)
	} else {
		interactor = &(db.FProjects)
	}

	arr := make([]ReturnableProjectsInfo, len(*interactor))

	for i := range *interactor {
		arr[i].Name = (*interactor)[i].Name
		arr[i].Description = (*interactor)[i].Description

	}

	returnableProjectsInfo = arr

	return &returnableProjectsInfo
}

