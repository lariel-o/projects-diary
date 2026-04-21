package main

import (
	"github.com/lariel-o/projects-diary/database"
	"github.com/lariel-o/projects-diary/database/querys/task"
)

func main() {
	database.Init()
	task.New(0, database.Task{
		Title: "Ola minha nova task",
	})
}

