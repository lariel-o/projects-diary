package display

import (
	"fmt"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type create struct {
	inputs []textinput.Model
	texts []string
	cursor uint8
	inputsCount uint8
}

var createDisplay = create{[]textinput.Model{}, []string{}, 0, 0}

func (m *create) init() {
	m.inputsCount = 1

	m.inputs = make([]textinput.Model, m.inputsCount)
	m.texts = make([]string, m.inputsCount)

	t := textinput.New()

	for i := range m.inputsCount {
		switch i {
		case 0:
			t.Prompt = "⤷ "
			t.Placeholder = "Task name" 
			t.Focus() 
			t.CharLimit = 120
			t.SetWidth(90)
			m.inputs[i] = t

			m.texts[i] = "Type the task name here"
		}
	}
}

func (m *create) update(msg string, realMsg tea.Msg, main *Daishi) tea.Cmd {
	switch msg {
	case "ctrl+c", "esc":
		return tea.Quit
	
	case "down":
		if m.cursor == m.inputsCount - 1 {
			m.cursor = 0
			m.inputs[m.inputsCount - 1].Blur()
		} else {
			m.cursor++
			m.inputs[m.cursor - 1].Blur()
		}

		m.inputs[m.cursor].Focus()

	case "up":
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

func (m create) view() (string, *tea.Cursor) {
	var c *tea.Cursor

	toReturn := ""

	for i := range m.inputsCount {
		// decide where the cursor is supposed to be
		switch i {
		case m.cursor:
			c = m.inputs[i].Cursor()
		}

		// create the str to be returned
		toReturn += fmt.Sprintf("%s\n%s\n", m.texts[i], m.inputs[i].View())
	}

	return toReturn, c
}

