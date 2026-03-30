package display

import tea "charm.land/bubbletea/v2"
import "fmt"

type world struct {
	projectsTitle []string  
	cursor uint16 
}

var worldDisplay = world {[]string{"Lucas", "Ariel", "Oliveira", "Moreira"}, 0}

func (m world) update(msg string, main *Daishi) tea.Cmd {
	switch msg {
	case "q":
		return tea.Quit

	case "k", "up":
		if worldDisplay.cursor == 0 { 
			worldDisplay.cursor = uint16(len(m.projectsTitle) - 1)
			return nil
		}
		worldDisplay.cursor -= 1
		return nil
	
	case "j", "down":
		if worldDisplay.cursor == uint16(len(m.projectsTitle) - 1) {
			worldDisplay.cursor = 0
			return nil
		}
		worldDisplay.cursor += 1
		return nil
	}

	return nil
}

func (m world) view() string {
	toReturn := ""
	for i, j := range m.projectsTitle {
		cursor := " "
		content := j

		if worldDisplay.cursor == uint16(i) {
			cursor = ">"
		}
		
		toReturn += fmt.Sprintf("%s %s\n", cursor, content)
	}

	return toReturn
}

