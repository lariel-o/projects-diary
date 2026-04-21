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

