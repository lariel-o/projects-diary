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
	task.ID = DB.World[tracer].LastTaskID

	// sum 1 at TasksCount indicating that a new task is being added
	DB.World[tracer].GTasksCount++
	DB.World[tracer].LastTaskID++

	// add the new task to the db
	DB.World[tracer].Tasks = append(DB.World[tracer].Tasks, task)

	// write the new DB
	if err := writeAtDatabase(); err != nil { return err }

	return nil
}

func RemoveTask(src1, src2 uint16) error {
	newTasks := make([]TaskStructModel, DB.World[src1].GTasksCount - 1)

	count := 0
	for i := range DB.World[src1].GTasksCount {
		if i == src2 {
			continue
		}

		newTasks[count] = DB.World[src1].Tasks[i]
		count++
	}

	DB.World[src1].Tasks = newTasks
	DB.World[src1].GTasksCount--
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

func MarkTaskAsFinished(src1, src2 uint16) error {
	// this function make something like a shift left in the tasks slice (i don't know the correct algorithm name), ex:
	// [GTsk1, GTsk2, GTsk3, GTsk4, FTsk1, FTsk2]
	// [GTsk1, NTsk0, GTsk3, GTsk4, FTsk1, FTsk2]

	// [GTsk1, GTsk3, GTsk3, GTsk4, FTsk1, FTsk2]
	// [GTsk1, GTsk3, GTsk4, GTsk4, FTsk1, FTsk2]
	// [GTsk1, GTsk3, GTsk3, NTsk0, FTsk1, FTsk2]

	currentProject := &DB.World[src1]
	currentFinishedTask := currentProject.Tasks[src2]

	// Look that LCTI (last completed task indice) =
	// (Total tasks) - FTasksCount - 1 =
	// GTasksCount + FTasksCount - FTasksCount - 1 =
	// GTasksCount - 1
	LCTI := currentProject.GTasksCount - 1 
	for i := src2; i < LCTI; i++ {
		DB.World[src1].Tasks[i] = DB.World[src1].Tasks[i+1]
	}

	currentFinishedTask.Finished = true
	DB.World[src1].Tasks[LCTI] = currentFinishedTask

	currentProject.GTasksCount--
	currentProject.FTasksCount++

	if e := writeAtDatabase(); e != nil { return e }

	return nil
}

func MarkTaskAsOngoing(src1, src2 uint16) error {
	currentProject := &DB.World[src1]
	currentOngoingTask := currentProject.Tasks[src2]

	// Look that LFTI (last finished task indice) =
	// (Total tasks) - FTasksCount =
	// GTasksCount + FTasksCount - FTasksCount =
	// GTasksCount
	LFTI := currentProject.GTasksCount
	for i := src2; i > LFTI; i-- {
		currentProject.Tasks[i] = currentProject.Tasks[i-1]
	}

	currentOngoingTask.Finished = false
	DB.World[src1].Tasks[LFTI] = currentOngoingTask

	currentProject.GTasksCount++
	currentProject.FTasksCount--

	if e := writeAtDatabase(); e != nil { return e }

	return nil
}
