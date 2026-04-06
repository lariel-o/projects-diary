package display

import( 
	"fmt"

	"github.com/lariel-o/projects-diary/data"

	tea "charm.land/bubbletea/v2"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/table"
)

type world struct {
	cursor uint16 

	showingDescription uint16
	isShowingDescription bool

	isSwapingProject bool
}

var headers = []string{"ID", "Project Name", "Description", "Created at"}

var worldDisplay = world{0, 0, false, false}

func (m *world) update(msg string, main *Daishi) tea.Cmd {
	switch msg {
	case "q", "esc":
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

		if data.DB.ProjectsCount != 0 {
			deleteDisplay = deleteIt {
				what: 0,
				projectTracer: m.cursor,
				taskTracer: 0,
				confirm: false,
			}

			main.lastOne = main.who
			main.who = 2
			m.isSwapingProject = false
		}

	case "a":
		m.isSwapingProject = false
		main.lastOne = main.who
		main.who = 3

	case "e":
		if data.DB.ProjectsCount != 0 {
			editProjectDisplay.tracer = m.cursor
			editProjectDisplay.setDefaultValues()

			main.lastOne = main.who
			main.who = 5
			m.isSwapingProject = false
		}
	}

	return nil
}

func (m world) view() string {
	// check if exist any project
	if data.DB.ProjectsCount == 0 {
		return "Nothing here"
	}

	if m.isShowingDescription {
		return fmt.Sprintf("%s\n⤷ %s", 
			data.DB.World[m.cursor].ProjectName,
			data.DB.World[m.cursor].Description)
	}

	rows := make([][]string, data.DB.ProjectsCount)
	for i := range data.DB.ProjectsCount {
		rows[i] = []string{
			fmt.Sprint(data.DB.World[i].ID),
			data.DB.World[i].ProjectName + "\nhoho",
			data.DB.World[i].Description,
			"0000",
		}
	}

	t := table.New().Headers(headers...).Rows(rows...).
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground( lipgloss.Color("99") )).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return lipgloss.NewStyle().Align(lipgloss.Center).Padding(0, 2, 0, 2)

			case uint16(row) == m.cursor && !m.isSwapingProject:
				return lipgloss.NewStyle().Background( lipgloss.Color("203") ).Align(lipgloss.Center)

			case uint16(row) == m.cursor && m.isSwapingProject:
				return lipgloss.NewStyle().Background( lipgloss.Color("203") ).Align(lipgloss.Right)


			default:
				return lipgloss.NewStyle().Align(lipgloss.Center)
			}
		})

	return lipgloss.Sprint(t)
}

