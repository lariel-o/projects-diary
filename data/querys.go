package data

import (
	"encoding/json"
	"os"
	// "fmt"
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
	project.ID = DB.LastProjectID

	// Sum 1 in the ProjectsCount indicating a new project is being added
	DB.ProjectsCount++ 
	DB.LastProjectID++

	// Save the new DB at the volatile memory
	DB.World = append(DB.World, project)

	// Write the new DB at the non-volatile memory
	if err := writeAtDatabase(); err != nil { return err }

	return nil
}

func RemoveProject(src uint16) error {
	newWorld := make([]ProjectStructModel, DB.ProjectsCount - 1)

	count := 0
	for i := range DB.ProjectsCount {
		if i == src {
			continue
		}

		newWorld[count] = DB.World[i]
		count++
	}

	DB.World = newWorld
	DB.ProjectsCount--
	if err := writeAtDatabase(); err != nil { return err }
	return nil
}

// change projects at the position src and dst
func SwapProjects(src, dst uint16, permission bool) { 
	if !permission {
		return
	}
	save := DB.World[src]
	DB.World[src] = DB.World[dst]
	DB.World[dst] = save
	writeAtDatabase()
}

func EditProject(n, d string, tracer uint16) error {
	DB.World[tracer].ProjectName = n
	DB.World[tracer].Description = d

	if err := writeAtDatabase(); err != nil { return err }
	return nil
}

func AddNewTask(tracer uint16, task TaskStructModel) error {
	// set task ID
	task.ID = DB.World[tracer].TasksCount

	// sum 1 at TasksCount indicating that a new task is being added
	DB.World[tracer].TasksCount++
	DB.World[tracer].LastTaskID++

	// add the new task to the db
	DB.World[tracer].Tasks = append(DB.World[tracer].Tasks, task)

	// write the new DB
	if err := writeAtDatabase(); err != nil { return err }

	return nil
}

func RemoveTask(src1, src2 uint16) error {
	newTasks := make([]TaskStructModel, DB.World[src1].TasksCount - 1)

	count := 0
	for i := range DB.World[src1].TasksCount {
		if i == src2 {
			continue
		}

		newTasks[count] = DB.World[src1].Tasks[i]
		count++
	}

	DB.World[src1].Tasks = newTasks
	DB.World[src1].TasksCount--
	if err := writeAtDatabase(); err != nil { return err }
	return nil
}

// change tasks at the position src and dst where the project have id tracer
func SwapTasks(src, dst, tracer uint16, permission bool) { 
	if !permission {
		return
	}
	save := DB.World[tracer].Tasks[src]
	DB.World[tracer].Tasks[src] = DB.World[tracer].Tasks[dst]
	DB.World[tracer].Tasks[dst] = save
	writeAtDatabase()
}

func EditTask(c string, pTracer, tTracer uint16) error {
	DB.World[pTracer].Tasks[tTracer].Content = c
	if err := writeAtDatabase(); err != nil { return err }
	return nil
}

