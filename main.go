package main

import (
	"fmt"
	"os"
	tea "charm.land/bubbletea/v2"
)

type worldModel struct {
	tasks []string
	cursor int
}

func (m worldModel) Init() tea.Cmd {
	return nil
}

func (m worldModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:

		switch msg.String() {
		case "ctrl-c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor == 0 {
				m.cursor = len(m.tasks) - 1
			} else {
				m.cursor -= 1
			}

		case "down", "j":
			if m.cursor == len(m.tasks) - 1 {
				m.cursor = 0
			} else {
				m.cursor += 1
			}
		}
	}

	return m, nil
}

func (m worldModel) View() tea.View {
	s := "=== === === === World === === === ===\n\n"
	for whereIsCursor, task := range m.tasks {
		cursor := " "
		if whereIsCursor == m.cursor {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, task)
	}

	s += "\nPress q to quit"
	return tea.NewView(s)
}

func main() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		fmt.Println("ERRRRR")
		os.Exit(1)
	}
}

func initialModel() worldModel {
	return worldModel {
		tasks: []string{"Task 1", "Task 2", "Task 3", "Task 4", },
		cursor: 0,
	}
}

