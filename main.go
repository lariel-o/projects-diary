package main

import (
	"github.com/lariel-o/projects-diary/database"
	"github.com/lariel-o/projects-diary/database/querys/task"
)

func main() {
	database.Init()
	task.Change(0, 1, true, true, database.Task{
		Title: "Meu novo titulo",
	})
}

