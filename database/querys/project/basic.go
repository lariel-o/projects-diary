package project

import (
	"github.com/lariel-o/projects-diary/database"
	"github.com/lariel-o/projects-diary/database/querys"
)

var DB = &database.DB

// Create a new project and add it to the non-volatile
func New(PJ database.Project) {
	// PJ = New Project

	PJ.ID = DB.NextProjectID 

	// add the new project to the volatile memory
	DB.OProjects = append(DB.OProjects, PJ)

	// ## Manipulate the geral database section (the database struct)
	DB.OProjectsCount++
	DB.NextProjectID++
	
	// write the new project at the non-volatile memory
	querys.WriteAtDatabase()
}

// Execute an algorithm to remove an element from a slice.
// In this context, it remove a project from the 
// non-volatiole memory using a indice!
func Del(IND uint16, ISO bool) {
	// IND == INDICE
	// ISO == Is Ongoing?
	

	if ISO {
		// save the number of ongoing projects that have now
		PJC := DB.OProjectsCount

		// create the new slice that will be replaced at 
		// databe.OProjects
		to_replace := make([]database.Project, PJC - 1)

		// ### Interate to set the propers values at the arr
		counter := 0
		for i := range PJC {
			// make sure that the element at IND will not be
			// in the new array
			if IND == i { continue }

			to_replace[counter] = DB.OProjects[i]
			counter++
		}

		// replace now
		DB.OProjects = to_replace

		// decrement ProjectsCount to make sense
		DB.OProjectsCount--
	} else {
		// save the number of finished projects that have now
		PJC := DB.FProjectsCount

		// create the new slice that will be replaced at 
		// databe.FProjects
		to_replace := make([]database.Project, PJC - 1)

		// ### Interate to set the propers values at the arr
		counter := 0
		for i := range PJC {
			// make sure that the element at IND will not be
			// in the new array
			if IND == i { continue }

			to_replace[counter] = DB.FProjects[i]
			counter++
		}

		// replace now
		DB.FProjects = to_replace

		// decrement ProjectsCount to make sense
		DB.FProjectsCount--
	}

	querys.WriteAtDatabase()
}

// Change informations from a project
func Change(IND uint16, PJ database.Project, ISO bool) {
	// IND = INDICE
	// ISO = Is Ongoing?
	// PJ = Project

	// change at Ongoing projects slice if is supposed to
	if ISO {
		PJ.ID = DB.OProjects[IND].ID
		DB.OProjects[IND] = PJ

	// change at Finished projects slice if is supposed to
	} else {
		PJ.ID = DB.FProjects[IND].ID
		DB.FProjects[IND] = PJ
	}

	querys.WriteAtDatabase()
}

// Move a project Up or Down
func Mov(IND uint16, DRC bool, ISO bool) {
	// DRC = Direction, when true move UP when false move DOWN
	// IND = Indice
	// ISO = Is ongoing project

	// contais or OProjects or FProjects, it will be set
	// by the next condition
	var object *[]database.Project

	// contains length of object
	var objectC uint16

	if ISO { 
		object = &(DB.OProjects)
		objectC = DB.OProjectsCount
	} else { 
		object = &(DB.FProjects)
		objectC = DB.FProjectsCount
	}

	// ### Move
	// Save the one who will be moved
	save := (*object)[IND]

	// When it's to move up
	if DRC {
		// Mov to the final of the slice
		if IND == 0 {
			(*object)[0] = (*object)[objectC - 1]
			(*object)[objectC - 1] = save

		// Move as expected
		} else {
			(*object)[IND] = (*object)[IND - 1]
			(*object)[IND - 1] = save
		}

		
	// When it's to move down
	} else {
		// Move the the start of the slice
		if IND == objectC - 1 {
			(*object)[objectC - 1] = (*object)[0] 
			(*object)[0] = save

		// Move as expected
		} else {
			(*object)[IND] = (*object)[IND + 1]
			(*object)[IND + 1] = save
		}
	}

	// PJ 0
	// PJ 1
	// PJ 2
	// PJ 3
	// PJ 4
	// PJ 5

	querys.WriteAtDatabase()
}

