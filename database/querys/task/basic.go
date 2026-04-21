package task

import (
	"github.com/lariel-o/projects-diary/database"
	"github.com/lariel-o/projects-diary/database/querys"
)

var DB = &database.DB

// Create a new task and add it to the non-volatile memory
func New(TRC uint16, TSK database.Task) {	
	// TSK = Task
	// TRC = Tracer, indicate the wanted project's indice

	// save the project that is going to be operated
	PJ := &DB.OProjects[TRC]

	TSK.ID = PJ.NextTaskID

	// add the new task to the volatile memory
	PJ.OTasks = append(PJ.OTasks, TSK)

	// ## Manipulate the geral project section (the project struct)
	PJ.OTasksCount++
	PJ.NextTaskID++

	// write the new project at the non-volatile memory
	querys.WriteAtDatabase()
}

// Execute an algorithm to remove an element from a slice.
// In this context, it remove a task from the 
// non-volatiole memory using the indice you want!
func Del(IND, TRC uint16, ISOP, ISOT bool) {
	// TSK = Task
	// TRC = Tracer, indicate the wanted project's indice
	// ISOP = Is Ongoing project?
	// ISOT = Is Ongoing task?
	// IND = Indice

	// save the project that is going to be operated
	var PJ *database.Project

	if ISOP {
		PJ = &DB.OProjects[TRC]
	} else {
		PJ = &DB.FProjects[TRC]
	}

	if ISOT {
		// save the number of ongoing tasks that have now
		TSKC := PJ.OTasksCount

		// create the new slice that will replace at
		// databe.OProjects
		to_replace := make([]database.Task, TSKC - 1)

		// ### Interate to set the propers values at the arr
		counter := 0
		for i := range TSKC {
			//make sure that the element at IND will not be
			// in the new array
			if IND == i { continue }
			to_replace[counter] = PJ.OTasks[i]
			counter++
		}

		// replace now
		PJ.OTasks = to_replace

		// Decrement TasksCount to make sense
		PJ.OTasksCount--
	} else {
		// save the number of finished tasks that have now
		TSKC := PJ.FTasksCount

		// create the new slice that will replace at
		// databe.OProjects
		to_replace := make([]database.Task, TSKC - 1)

		// ### Interate to set the propers values at the arr
		counter := 0
		for i := range TSKC {
			//make sure that the element at IND will not be
			// in the new array
			if IND == i { continue }
			to_replace[counter] = PJ.FTasks[i]
			counter++
		}

		// replace now
		PJ.FTasks = to_replace

		// Decrement TasksCount to make sense
		PJ.FTasksCount--
	}

	querys.WriteAtDatabase()
}

// Change informations of a task
func Change(IND, TRC uint16, ISOP, ISOT bool, TSK database.Task) {
	// IND = indice
	// TSK = Task
	// TRC = Tracer, indicate the wanted project's indice
	// ISOP = Is Ongoing Project?
	// ISOT = Is Ongoing Task?

	// Save the current PJ indicated by TRC
	var PJ *database.Project
	if ISOP {
		PJ = &DB.OProjects[TRC]
	} else {
		PJ = &DB.FProjects[TRC]
	}

	// change at ongoing tasks slice if is supposed to
	if ISOT {

		TSK.ID = PJ.OTasks[IND].ID
		PJ.OTasks[IND] = TSK


	// save at finished tasks slice if is supposed to
	} else {
		TSK.ID = PJ.FTasks[IND].ID
		PJ.FTasks[IND] = TSK
	}

	querys.WriteAtDatabase()
}

// Move a task Up or Down

