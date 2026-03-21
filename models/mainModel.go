package models

import (
	"fmt"
	// "encoding/csv"
	
	// auxs "github.com/lariel-o/projects-diary/auxiliaries"

	tea "charm.land/bubbletea/v2"
)

type WorldModel struct {
	ProjectsTitle []string
	ProjectsDescription []string
	Cursor uint

	// Save a pointer to the project who should be showing the description
	ShowingDescription uint
	IsShowingDescription bool

	ViewControl uint8 
	/* 	1 ~ Wold View (show the projects and is the main View)
		2 ~ Project (List the uncompleted task) 
		3 ~ Task    (Show the subtask selected)
		4 ~ Change info
	*/
}

func (m WorldModel) Init() tea.Cmd {
	return nil
}

func (m WorldModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.Cursor == 0 {
				m.Cursor = uint(len(m.ProjectsTitle) - 1)
			} else {
				m.Cursor -= 1
			}

		case "down", "j":
			if m.Cursor == uint(len(m.ProjectsTitle) - 1) {
				m.Cursor = 0
			} else {
				m.Cursor += 1
			}

		case "enter":
			// If there isn't anyone already showing the description
			if m.ShowingDescription == m.Cursor && m.IsShowingDescription {
				m.IsShowingDescription = false
			} else {
				m.ShowingDescription = m.Cursor
				m.IsShowingDescription = true
			}
		}
	}

	return m, nil
}

func (m WorldModel) View() tea.View {
	s := "=== === === === World === === === ===\n\n"
	for interator, projectTitle := range m.ProjectsTitle {
		cursor := " "
		description := ""

		if m.IsShowingDescription && m.ShowingDescription == uint(interator){
			description = "\n      ~ " + m.ProjectsDescription[m.ShowingDescription]
		}

		if uint(interator) == m.Cursor {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s %s\n", cursor, projectTitle, description)
	}

	s += "\nPress q to quit"
	return tea.NewView(s)
}

