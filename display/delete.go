package display

import (
	"github.com/lariel-o/projects-diary/data"

	tea "charm.land/bubbletea/v2"
)

type deleteIt struct  {
	what uint8 // what do you want to delete? a task? a project?
	projectTracer uint16
	taskTracer uint16

	confirm bool
}

var deleteDisplay = deleteIt{0, 0, 0, false}

func (m deleteIt) update(msg string, main *Daishi) tea.Cmd {
	switch msg {
	case "enter":
	 	if m.confirm {
			if m.what == 0 {
				data.RemoveProject(m.projectTracer)
				worldDisplay.cursor = 0
			} else if m.what == 1 {
				data.RemoveTask(m.projectTracer, m.taskTracer)
				projectDisplay.cursor = 0
			}
		}

		main.who = main.lastOne
		main.lastOne = 2

	case "q":
		main.who = main.lastOne
		main.lastOne = 2

	case "l", "left", "h", "right":
		deleteDisplay.confirm = !deleteDisplay.confirm	

	}


	return nil
}

func (m deleteIt) view() string{
	s := "Are you sure you want to delete?\n"
	if !m.confirm {
		s += "*NO        YES"
	} else {
		s += " NO       *YES"
	}

	return s
}

