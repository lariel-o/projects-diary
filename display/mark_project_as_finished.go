package display

import (
	"github.com/lariel-o/projects-diary/data"

	tea "charm.land/bubbletea/v2"
)

type markAsFinished struct  {
	what uint8 // what do you want to mark as finished? a task? a project?
	projectTracer uint16
	taskTracer uint16

	confirm bool
}

var markAsFinishedDisplay = markAsFinished{0, 0, 0, false}

func yesMark(m *markAsFinished) {
	if m.what == 0 {

	} else if m.what == 1 {
		currentTask := data.DB.World[m.projectTracer].Tasks[m.taskTracer]
		if currentTask.Finished {
			data.MarkTaskAsOngoing(m.projectTracer, m.taskTracer)
		} else {
			data.MarkTaskAsFinished(m.projectTracer, m.taskTracer)
		}
	}
}

func (m *markAsFinished) update(msg string, main *Daishi) tea.Cmd {
	switch msg {
	case "enter":
	 	if m.confirm {
			yesMark(m)
		}

		main.who = main.lastOne
		main.lastOne = 7

	case "y":
		yesMark(m)

		main.who = main.lastOne
		main.lastOne = 7

	case "q", "n", "N", "esc":
		main.who = main.lastOne
		main.lastOne = 7

	case "l", "left", "h", "right":
		m.confirm = !m.confirm	

	}


	return nil
}

func (m markAsFinished) view() string{
	s := "Are you sure you want to mark it as "
	if data.DB.World[m.projectTracer].Finished {
		s += "Unfinished\n"
	} else {
		s += "Finished\n"
	}

	if !m.confirm {
		s += "*NO        yes"
	} else {
		s += " NO       *yes"
	}

	return s
}


