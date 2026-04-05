package display

import( 
	"fmt"

	"github.com/lariel-o/projects-diary/data"

	tea "charm.land/bubbletea/v2"
)

type project struct {
	cursor uint16

	showingDescription uint16
	isShowingDescription bool

	isSwapingTask bool

	projectTracer uint16
}

var projectDisplay = project{0, 0, false, false, 0}

func (m *project) update(msg string, main *Daishi) tea.Cmd {
	switch msg {
	case "q", "esc":
		main.lastOne = main.who
		main.who = 0

	// move cursor up
	case "k", "up":
		if m.cursor == 0 { 
			m.cursor = uint16(data.DB.World[m.projectTracer].TasksCount - 1)

			// try to swap if isSwaping is true
			data.SwapTasks(0, m.cursor, m.projectTracer, m.isSwapingTask)
		} else {
			m.cursor -= 1

			// try to swap if isSwaping is true
			data.SwapTasks(m.cursor + 1, m.cursor, m.projectTracer, m.isSwapingTask)
		}
		return nil
	
	// move cursor down
	case "j", "down":
		if m.cursor == uint16(data.DB.World[m.projectTracer].TasksCount - 1) {
			m.cursor = 0

			// try to swap if isSwaping is true
			data.SwapTasks(uint16(data.DB.World[m.projectTracer].TasksCount - 1), 0, m.projectTracer, m.isSwapingTask)
		} else {
			m.cursor += 1

			// try to swap if isSwaping is true
			data.SwapTasks(m.cursor - 1, m.cursor, m.projectTracer, m.isSwapingTask)
		}

		return nil

	// show description
	case "l", "right":
		m.isShowingDescription = true
		m.showingDescription = m.cursor
		return nil
	
	// unshow description
	case "h", "left":
		m.isShowingDescription = false

	// active and un active the swaping mode
	case "s":
		m.isSwapingTask = !m.isSwapingTask

	case "d":
		m.isSwapingTask = false

		if data.DB.World[m.projectTracer].TasksCount != 0 {
			deleteDisplay = deleteIt {
				what: 1,
				projectTracer: m.projectTracer,
				taskTracer: m.cursor,
				confirm: false,
			}

			main.lastOne = main.who
			main.who = 2
		}

	case "a":
		m.isSwapingTask = false
		main.lastOne = main.who
		main.who = 4

		addTaskDisplay.tracer = m.projectTracer

	case "e":
		m.isSwapingTask = false
		main.lastOne = main.who
		main.who = 6

		editTaskDisplay.projectTracer 	= m.projectTracer
		editTaskDisplay.taskTracer 		= m.cursor
		editTaskDisplay.setDefaultValues()
	}

	return nil
}

func (m project) view() string {
	// check if exist any task
	if data.DB.World[m.projectTracer].TasksCount == 0 {
		return "Nothing here"
	}

	toReturn := ""

	// cursor* title 
	for i := range data.DB.World[m.projectTracer].TasksCount {
		swapingPadding := ""
		cursor := " "
		content := ""

		if m.isSwapingTask && m.cursor == i {
			swapingPadding = "    "
		}

		// set cursor
		if m.cursor == i {
			cursor = ">"
		}

		// set task content
		content = data.DB.World[m.projectTracer].Tasks[i].Content

		toReturn += fmt.Sprintf("%s%s %s\n", swapingPadding, cursor, content)
	}

	return toReturn
}

