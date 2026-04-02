package data

import (
	"encoding/json"
	"os"
)

func loadDatabase() error {
	data, err := os.ReadFile(DatabaseInfo.FilesPath["main"])
	if err != nil { return err }

	// Read from db and transform it in a  struct
	var db WorldStructModel
	err = json.Unmarshal(data, &db)
	if err != nil { return err }

	DB = db
	return nil
}

func writeAtDatabase() error {
	b, err := json.Marshal(DB)
	if err != nil { return err }

	if err := os.WriteFile(DatabaseInfo.FilesPath["main"], b, 0666); err != nil { 
		return err
	}

	return nil
}

func AddNewProject(project ProjectStructModel) error {
	// Put the proper ID
	project.ID = DB.ProjectsCount

	// Sum 1 in the ProjectsCount indicating a new project is being added
	DB.ProjectsCount++ 

	// Save the new DB at the volatile memory
	DB.World = append(DB.World, project)

	// Write the new DB at the non-volatile memory
	if err := writeAtDatabase(); err != nil { return err }

	return nil
}

// change the projects at the position src and dst
func SwapProjects(src, dst uint16, permission bool) { 
	if !permission {
		return
	}
	save := DB.World[src]
	DB.World[src] = DB.World[dst]
	DB.World[dst] = save
	writeAtDatabase()
}

func AddNewTask(tracer uint16, task TaskStructModel) error {
	// set task ID
	task.ID = DB.World[tracer].TasksCount	

	// sum 1 at TasksCount indicating that a new task is being added
	DB.World[tracer].TasksCount++

	// add the new task to the db
	DB.World[tracer].Tasks = append(DB.World[tracer].Tasks, task)

	// write the new DB
	if err := writeAtDatabase(); err != nil { return err }

	return nil
}

func SwapTasks(src, dst, tracer uint16, permission bool) { 
	if !permission {
		return
	}
	save := DB.World[tracer].Tasks[src]
	DB.World[tracer].Tasks[src] = DB.World[tracer].Tasks[dst]
	DB.World[tracer].Tasks[dst] = save
	writeAtDatabase()
}


