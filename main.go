package main

import (
	"github.com/lariel-o/projects-diary/database"
	"github.com/lariel-o/projects-diary/database/querys/project"
)

func main() {
	database.Init()
	project.New()
}

