package main

import (
	"github.com/lariel-o/projects-diary/database"
	"github.com/lariel-o/projects-diary/database/querys/project"
)

func main() {
	database.Init()
	// project.New(database.Project {
	// 	Title: "Something new 4",
	// })
	// project.Del(1, true)
	project.Mov(3, false, true)
}

