package main

import (
	"github.com/lariel-o/projects-diary/database"
	"github.com/lariel-o/projects-diary/database/querys/project"
)

func main() {
	database.Init()
	// project.New("Meu segundo projeto!!")
	// project.New("Meu terceiro projeto!!")
	// project.New("Meu quarto projeto!!")
	// project.New("Meu quinto projeto!!")
	project.New("Meu quinto projeto!!")
	// project.Del(0, true)
}

