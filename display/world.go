package display

import tea "charm.land/bubbletea/v2"

type world struct {
	projectsTitle []string
}

var worldDisplay = world {[]string{"Lucas", "Ariel", "Oliveira", "Moreira"}}

func (m world) update(msg string) tea.Cmd {
	switch msg {
	case "q":
		return tea.Quit


	}

	return nil
}

