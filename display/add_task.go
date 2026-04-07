package display

import (
	"fmt"
	"time"
	"strconv"

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

	dateMode bool // if true add if false pass a literaly
}

var addTaskDisplay = addTask{}

func eraseTasksInput() {
	addTaskDisplay = addTask{[]textinput.Model{}, []string{}, 0, 0, 0, false}
	addTaskDisplay.init()
}

func (m *addTask) init() {
	m.inputsCount = 6

	m.inputs = make([]textinput.Model, m.inputsCount)
	m.texts = make([]string, m.inputsCount)

	t := textinput.New()
	t.Prompt = "⤷ "
	t.CharLimit = 120
	t.SetWidth(90)

	for i := range m.inputsCount {
		if i > 0 { t.Blur() }
		if i ==  0 {
			t.Focus() 

			m.inputs[i] = t
			m.texts[i] = "Task content"
		}

		if i >= 1 { t.Prompt = "" }
		if i == 1 {
			t.Placeholder = "YYYY"
			t.CharLimit = 4
			t.SetWidth(4)
		}
		if i >= 2 {
			t.CharLimit = 2
			t.SetWidth(2)
		}

		if i == 2 { t.Placeholder = "MM" }
		if i == 3 { t.Placeholder = "DD" }
		if i == 4 { t.Placeholder = "hh" }
		if i == 5 { t.Placeholder = "mm" }

		m.inputs[i] = t
	}
}

func (m *addTask) update(msg string, realMsg tea.Msg, main *Daishi) tea.Cmd {
	switch msg {
	case "enter":
		// don't allow to create tasks with empty content
		if m.inputs[0].Value() != "" {
			// check if the user have typed any date
			willAddExpire := false
			haveAnyInputNull := false

			for i := 1; i < 6; i++  {
				if m.inputs[i].Value() != "" {
					willAddExpire = true
				} else { haveAnyInputNull = true }
			}

			currentTime := time.Now()
			var expireAt time.Time

			// put the dates typed into variables
			tYear, _		:= strconv.Atoi(m.inputs[1].Value())
			tMonth, _ 		:= strconv.Atoi(m.inputs[2].Value())
			tDay, _ 		:= strconv.Atoi(m.inputs[3].Value())
			tHour, _ 		:= strconv.Atoi(m.inputs[4].Value())
			tMinute, _ 		:= strconv.Atoi(m.inputs[5].Value())

			if willAddExpire && !m.dateMode {
				expireAt = time.Date( 
					currentTime.Year() + 	tYear, 
					currentTime.Month() +	time.Month(tMonth), 
					currentTime.Day() + 	tDay,  
					currentTime.Hour() + 	tHour, 
					currentTime.Minute() + 	tMinute,
					currentTime.Second(), currentTime.Nanosecond(),
					time.Local,
				)
			} else if willAddExpire {
				expireAt = time.Date(tYear, time.Month(tMonth), tDay, tHour, tMinute, 0, 0, time.Local)
				if haveAnyInputNull || expireAt.Sub(currentTime) < 0 { return nil }
			}

			data.AddNewTask(m.tracer, data.TaskStructModel{
				Content: m.inputs[0].Value(),
				CreatedAt: currentTime,
				ExpireAt: expireAt,
				HaveExpireTime: willAddExpire,
			})
		
			main.who = main.lastOne
			main.lastOne = 4

			eraseTasksInput()
		}
	
	case "ctrl+c", "esc":
		main.who = main.lastOne
		main.lastOne = 4

		eraseTasksInput()
	
	case "down", "tab":
		if m.cursor == m.inputsCount - 1 {
			m.cursor = 0
			m.inputs[m.inputsCount - 1].Blur()
		} else {
			m.cursor++
			m.inputs[m.cursor - 1].Blur()
		}

		m.inputs[m.cursor].Focus()

	case "up", "shift+tab":
		if m.cursor == 0 {
			m.cursor = m.inputsCount - 1
			m.inputs[0].Blur()
		} else {
			m.cursor--
			m.inputs[m.cursor + 1].Blur()
		}

		m.inputs[m.cursor].Focus()

	case "ctrl+p":
		m.dateMode = !m.dateMode
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
		if i == 0 {
			toReturn += fmt.Sprintf("%s\n%s\n\n", m.texts[0], m.inputs[0].View())
		}
		if i == 1 { toReturn += fmt.Sprintf("Date Limit\n⤷ ") } 
		if i >= 1 { toReturn += m.inputs[i].View() }
	}

	if !m.dateMode {
		toReturn += "\n\n===Date mode: ADD==="
	} else {
		toReturn += "\n\n===Date mode: DATE==="
	}

	return toReturn, c
}

