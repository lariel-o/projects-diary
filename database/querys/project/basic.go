package project

import (
	"github.com/lariel-o/projects-diary/database"
	"github.com/lariel-o/projects-diary/database/querys"
)

var DB = &database.DB

// Create a new project and add it to the non-volatile
func New(title string) {
	// save the new project who will be added
	nP := database.Project {
		Title: title,
		ID: DB.NextProjectID,
	}

	// add the new project to the non-volatile memory
	DB.OProjects = append(DB.OProjects, nP)

	// ## Manipulate the geral database section (the database struct)
	DB.OProjectsCount++
	DB.NextProjectID++
	
	// write the new project at the non-volatile memory
	querys.WriteAtDatabase()
}


// Execute an algorithm to remove an element from a slice.
// In this context, it remove a project from the 
// non-volatiole memory using a indice!
func Del(IND uint16, isO bool) {
	// IND == INDICE
	// isF == Is Ongoing?
	

	if isO {
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
	}

	querys.WriteAtDatabase()
}

