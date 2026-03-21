package models

import (
	"fmt"
	
	tea "charm.land/bubbletea/v2"
)

type WorldModel struct {
	Tasks []string
	Cursor int
	CurrentView uint8 
	/*
		1 ~ World View (show the projects and is the main View)
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
		case "ctrl-c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.Cursor == 0 {
				m.Cursor = len(m.Tasks) - 1
			} else {
				m.Cursor -= 1
			}

		case "down", "j":
			if m.Cursor == len(m.Tasks) - 1 {
				m.Cursor = 0
			} else {
				m.Cursor += 1
			}

		case "enter":
			fmt.Println("enter")

		}
	}

	return m, nil
}

func (m WorldModel) View() tea.View {
	s := "=== === === === World === === === ===\n\n"
	for whereIsCursor, task := range m.Tasks {
		cursor := " "
		if whereIsCursor == m.Cursor {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, task)
	}

	s += "\nPress q to quit"
	return tea.NewView(s)
}

