// This file have all structs needed to work with 
// json

// Note that the limit of projects are 2^16 and 
// tasks are 2^16 per project
package database

type Task struct {
	Title string
}

type Project struct {
	// Project title
	Title string

	// Have all ongoing projects
	OTasks []Task

	// Have all finished projects
	FTasks []Task

	// Count how many tasks have the status ongoing/finished
	OTasksCount uint16
	FTasksCount uint16

	NextTaskID uint16

	ID uint16
}

type database struct {
	// Have all ongoing projects
	OProjects []Project 

	// Have all finished projects
	FProjects []Project

	// Count how many projects have the status ongoing/finished
	OProjectsCount uint16
	FProjectsCount uint16

	NextProjectID uint16
}

