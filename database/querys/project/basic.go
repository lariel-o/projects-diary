package project

import (
	"fmt"
	"github.com/lariel-o/projects-diary/database"
)

var DB = &database.DB

// Create a new project and add it to database 
func New() {
	fmt.Println(DB)
}

