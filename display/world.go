package display

import( 
	"fmt"

	"github.com/lariel-o/projects-diary/data"

	tea "charm.land/bubbletea/v2"
)

type world struct {
	cursor uint16 

	showingDescription uint16
	isShowingDescription bool

	isSwapingProject bool
}

var worldDisplay = world{0, 0, false, false}

func (m *world) update(msg string, main *Daishi) tea.Cmd {
	switch msg {
	case "q":
		return tea.Quit

	case "enter":
		// set swap to false to avoid bugs
		m.isSwapingProject = false

		// set the default struct of the project display (daishi will manage it well)
		projectDisplay = project {
			cursor: 0,
			showingDescription: 0,
			isShowingDescription: false,
			isSwapingTask: false,
			projectTracer: m.cursor,
		}
		
		main.lastOne = main.who
		main.who = 1


	// move cursor up
	case "k", "up":
		if m.cursor == 0 { 
			m.cursor = uint16(data.DB.ProjectsCount - 1)

			// try to swap if isSwaping is true
			data.SwapProjects(0, m.cursor, m.isSwapingProject)
		} else {
			m.cursor -= 1

			// try to swap if isSwaping is true
			data.SwapProjects(m.cursor + 1, m.cursor, m.isSwapingProject)
		}
		return nil
	
	// move cursor down
	case "j", "down":
		if m.cursor == uint16(data.DB.ProjectsCount - 1) {
			m.cursor = 0

			// try to swap if isSwaping is true
			data.SwapProjects(uint16(data.DB.ProjectsCount - 1), 0, m.isSwapingProject)
		} else {
			m.cursor++

			// try to swap if isSwaping is true
			data.SwapProjects(m.cursor - 1, m.cursor, m.isSwapingProject)
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
		m.isSwapingProject = !m.isSwapingProject

	case "d":
		m.isSwapingProject = false

		deleteDisplay = deleteIt {
			what: 0,
			projectTracer: m.cursor,
			taskTracer: 0,
			confirm: false,
		}
		
		main.lastOne = main.who
		main.who = 2

	case "c":
		m.isSwapingProject = false
		main.lastOne = main.who
		main.who = 3
	}

	return nil
}

func (m world) view() string {
	// check if exist any project
	if data.DB.ProjectsCount == 0 {
		return "Nothing here"
	}

	toReturn := ""

	// cursor* title 
	for i := range data.DB.ProjectsCount {
		swapingPadding := ""
		cursor := " "
		title := ""
		description := ""

		if m.isSwapingProject && m.cursor == i {
			swapingPadding = "    "
		}

		// set cursor
		if m.cursor == i {
			cursor = ">"
		}

		// set description
		// if is trying to swap don't allow to show description
		if m.isShowingDescription && m.showingDescription == i && !m.isSwapingProject {
			description = "\n    ~ " + data.DB.World[i].Description
		}

		// set title
		title = data.DB.World[i].ProjectName

		toReturn += fmt.Sprintf("%s%s %s%s\n", swapingPadding, cursor, title, description)
	}

	return toReturn
}

