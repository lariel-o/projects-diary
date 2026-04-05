package display

import (
	"fmt"

	"github.com/lariel-o/projects-diary/data"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type addTask struct {
	inputs []textinput.Model
	texts []string
	tracer uint16
	cursor uint8
	inputsCount uint8
}

var addTaskDisplay = addTask{}

func eraseTasksInput() {
	addTaskDisplay = addTask{[]textinput.Model{}, []string{}, 0, 0, 0}
	addTaskDisplay.init()
}

func (m *addTask) init() {
	m.inputsCount = 1

	m.inputs = make([]textinput.Model, m.inputsCount)
	m.texts = make([]string, m.inputsCount)

	t := textinput.New()
	t.Prompt = "⤷ "
	t.CharLimit = 120
	t.SetWidth(90)

	for i := range m.inputsCount {
		switch i {
		case 0:
			t.Focus() 
			m.inputs[i] = t

			m.texts[i] = "Task content"
		}
	}
}

func (m *addTask) update(msg string, realMsg tea.Msg, main *Daishi) tea.Cmd {
	switch msg {
	case "enter":
		// don't allow to create tasks with empty content
		if m.inputs[0].Value() != "" {
			data.AddNewTask(m.tracer, data.TaskStructModel{
				Content: m.inputs[0].Value(),
			})
		
			main.who = main.lastOne
			main.lastOne = 4

			eraseTasksInput()
		}
	
	case "ctrl+c", "esc":
		main.who = main.lastOne
		main.lastOne = 4

		eraseTasksInput()
	
	case "down", "shift+tab":
		if m.cursor == m.inputsCount - 1 {
			m.cursor = 0
			m.inputs[m.inputsCount - 1].Blur()
		} else {
			m.cursor++
			m.inputs[m.cursor - 1].Blur()
		}

		m.inputs[m.cursor].Focus()

	case "up", "tab":
		if m.cursor == 0 {
			m.cursor = m.inputsCount - 1
			m.inputs[0].Blur()
		} else {
			m.cursor--
			m.inputs[m.cursor + 1].Blur()
		}

		m.inputs[m.cursor].Focus()
	}	


	m.inputs[m.cursor], _ = m.inputs[m.cursor].Update(realMsg)

	return nil
}

func (m addTask) view() (string, *tea.Cursor) {
	var c *tea.Cursor

	toReturn := ""

	for i := range m.inputsCount {
		// decide where the cursor is supposed to be
		switch i {
		case m.cursor:
			c = m.inputs[i].Cursor()
		}

		// create the str to be returned
		toReturn += fmt.Sprintf("%s\n%s\n\n", m.texts[i], m.inputs[i].View())
	}

	return toReturn, c
}

