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

func yes(m *deleteIt) {
	if m.what == 0 {
		if worldDisplay.cursor == data.DB.ProjectsCount - 1 {
			worldDisplay.cursor -= 1
		}
		data.RemoveProject(m.projectTracer)
	} else {
		if projectDisplay.cursor == data.DB.World[m.projectTracer].GTasksCount - 1 {
			projectDisplay.cursor -= 1
		}
		data.RemoveTask(m.projectTracer, m.taskTracer)
	}

}

func (m *deleteIt) update(msg string, main *Daishi) tea.Cmd {
	switch msg {
	case "enter":
	 	if m.confirm {
			yes(m)
		}

		main.who = main.lastOne
		main.lastOne = 2

	case "y":
		yes(m)

		main.who = main.lastOne
		main.lastOne = 2

	case "q", "n", "N", "esc":
		main.who = main.lastOne
		main.lastOne = 2

	case "l", "left", "h", "right":
		m.confirm = !m.confirm	

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

