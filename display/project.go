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

func (m project) update(msg string, main *Daishi) tea.Cmd {
	switch msg {
	case "q":
		who = 0

	// move cursor up
	case "k", "up":
		if projectDisplay.cursor == 0 { 
			projectDisplay.cursor = uint16(data.DB.World[m.projectTracer].TasksCount - 1)

			// try to swap if isSwaping is true
			data.SwapTasks(0, projectDisplay.cursor, m.projectTracer, projectDisplay.isSwapingTask)
		} else {
			projectDisplay.cursor -= 1

			// try to swap if isSwaping is true
			data.SwapTasks(projectDisplay.cursor + 1, projectDisplay.cursor, m.projectTracer, projectDisplay.isSwapingTask)
		}
		return nil
	
	// move cursor down
	case "j", "down":
		if projectDisplay.cursor == uint16(data.DB.World[m.projectTracer].TasksCount - 1) {
			projectDisplay.cursor = 0

			// try to swap if isSwaping is true
			data.SwapTasks(uint16(data.DB.World[m.projectTracer].TasksCount - 1), 0, m.projectTracer, projectDisplay.isSwapingTask)
		} else {
			projectDisplay.cursor += 1

			// try to swap if isSwaping is true
			data.SwapTasks(projectDisplay.cursor - 1, projectDisplay.cursor, m.projectTracer, projectDisplay.isSwapingTask)
		}

		return nil

	// show description
	case "l", "right":
		projectDisplay.isShowingDescription = true
		projectDisplay.showingDescription = projectDisplay.cursor
		return nil
	
	// unshow description
	case "h", "left":
		projectDisplay.isShowingDescription = false

	// active and un active the swaping mode
	case "s":
		projectDisplay.isSwapingTask = !projectDisplay.isSwapingTask
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

