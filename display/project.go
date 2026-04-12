package display

import( 
	"fmt"
	"time"
	"slices"

	"github.com/lariel-o/projects-diary/data"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/table"
)

type project struct {
	cursor uint16

	isSwapingTask bool
	isShowingFinishedTasks bool

	projectTracer uint16
}

var projectDisplayHeader = []string{"ID", "Task content", "Expires in"}

var projectDisplay = project{0, false, false, 0}

func (m *project) update(msg string, main *Daishi) tea.Cmd {
	switch msg {
	case "q", "esc":
		main.lastOne = main.who
		main.who = 0

	// move cursor up
	case "k", "up":
		if m.cursor == 0 { 
			m.cursor = uint16(data.DB.World[m.projectTracer].GTasksCount - 1)

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
		currentProject := data.DB.World[m.projectTracer] 
		floor := uint16(0)

		if m.isShowingFinishedTasks {
			floor = currentProject.GTasksCount + currentProject.FTasksCount
		} else {
			floor = currentProject.GTasksCount
		}

		if m.cursor == floor - 1 {
			m.cursor = 0

			// try to swap if isSwaping is true
			data.SwapTasks(floor - 1, 0, m.projectTracer, m.isSwapingTask)
		} else {
			m.cursor += 1

			// try to swap if isSwaping is true
			data.SwapTasks(m.cursor - 1, m.cursor, m.projectTracer, m.isSwapingTask)
		}

		return nil

	// active and un active the swaping mode
	case "s":
		if m.cursor < data.DB.World[m.projectTracer].GTasksCount {
			m.isSwapingTask = !m.isSwapingTask
		}

	case "d":
		m.isSwapingTask = false

		if data.DB.World[m.projectTracer].GTasksCount != 0 {
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

	case "ctrl+f":
		markAsFinishedDisplay.what = 1
		markAsFinishedDisplay.projectTracer = m.projectTracer
		markAsFinishedDisplay.taskTracer = m.cursor

		main.lastOne = main.who
		main.who = 7
		m.isSwapingTask = false

	case "ctrl+h": 
		m.isShowingFinishedTasks = !m.isShowingFinishedTasks
	}

	return nil
}

func (m project) view() string {
	// check if exist any task
	currentProject := &data.DB.World[m.projectTracer] 

	if currentProject.GTasksCount == 0 {
		return "Nothing here"
	}

	ongoing := make([][]string, currentProject.GTasksCount)
	for i := range currentProject.GTasksCount {
		currentTask := currentProject.Tasks[i]
		expiresIn := ""

		// set the expire time if it exist
		if currentTask.HaveExpireTime {
			s := currentTask.ExpireAt.Sub(time.Now())
			expiresIn = fmt.Sprintf("%dh%dm", int(s.Hours()), int(s.Minutes()) - int(s.Hours())*60)
		}
		ongoing[i] = []string{
			fmt.Sprint(currentTask.ID),
			currentTask.Content,
			expiresIn,
		}
	}

	finished := make([][]string, currentProject.FTasksCount)
	fCount := 0
	for i := currentProject.GTasksCount; i < currentProject.FTasksCount + currentProject.GTasksCount; i++{
		currentTask := currentProject.Tasks[i]
		expiresIn := ""

		// set the expire time if it exist
		if currentTask.HaveExpireTime {
			s := currentTask.ExpireAt.Sub(time.Now())
			expiresIn = fmt.Sprintf("%dh%dm", int(s.Hours()), int(s.Minutes()) - int(s.Hours())*60)
		}
		finished[fCount] = []string{
			fmt.Sprint(currentTask.ID),
			currentTask.Content,
			expiresIn,
		}

		fCount++
	}

	rows := [][]string{}
	if m.isShowingFinishedTasks {
		rows = slices.Concat(ongoing, finished)
	} else {
		rows = ongoing
	}
	
	t := table.New().Headers(projectDisplayHeader...).Rows(rows...).
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground( lipgloss.Color("99") )).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return lipgloss.NewStyle().Align(lipgloss.Center).Padding(0, 2, 0, 2)

			case uint16(row) == m.cursor && !m.isSwapingTask:
				return lipgloss.NewStyle().Background(lipgloss.Color("203")).Align(lipgloss.Center)

			case uint16(row) == m.cursor && m.isSwapingTask:
				return lipgloss.NewStyle().Background(lipgloss.Color("203")).Align(lipgloss.Right)


			default:
				return lipgloss.NewStyle().Align(lipgloss.Center)
			}
		})

	return lipgloss.Sprint(t)
}

