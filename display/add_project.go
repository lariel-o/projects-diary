package display

import (
	"fmt"
	"time"
	"strconv"

	"github.com/lariel-o/projects-diary/data"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type addProject struct {
	inputs []textinput.Model
	texts []string

	cursor uint8
	inputsCount uint8

	dateMode bool // if true add if false pass as literaly
}

var addProjectDisplay = addProject{}

func eraseProjectsInput() {
	addProjectDisplay.init()
}

func (m *addProject) init() {
	// format
	addProjectDisplay = addProject{[]textinput.Model{}, []string{}, 0, 0, false}

	m.inputsCount = 7

	m.inputs = make([]textinput.Model, m.inputsCount)
	m.texts = make([]string, m.inputsCount)

	t := textinput.New()
	t.Prompt = "⤷ "
	t.CharLimit = 120
	t.SetWidth(90)

	for i := range m.inputsCount {
		if i > 0 { t.Blur() }

		if i == 0{
			t.Focus() 

			m.texts[i] = "Project name"
		}

		if i == 1 { m.texts[i] = "Project description" }
		if i >= 2 { t.Prompt = "" }

		if i == 2 {
			t.Placeholder = "YYYY"
			t.CharLimit = 4
			t.SetWidth(4)
		}

		if i >= 3 {
			t.CharLimit = 2
			t.SetWidth(2)
		}

		if i == 3 { t.Placeholder = "MM" }
		if i == 4 { t.Placeholder = "DD" }
		if i == 5 { t.Placeholder = "hh" }
		if i == 6 { t.Placeholder = "mm" }

		m.inputs[i] = t
	}
}

func (m *addProject) update(msg string, realMsg tea.Msg, main *Daishi) tea.Cmd {
	switch msg {
	case "enter":
		// don't allow to create projects with empty title
		if m.inputs[0].Value() != "" {
			// check if the user have typed any date
			willAddExpire := false
			haveAnyInputNull := false

			for i := 2; i < 7; i++ {
				if m.inputs[i].Value() != "" { 
					willAddExpire = true 
				} else { haveAnyInputNull = true }
			}

			currentTime := time.Now()
			var expireAt time.Time

			// put the dates typed into variables
			tYear, _		:= strconv.Atoi(m.inputs[2].Value())
			tMonth, _ 		:= strconv.Atoi(m.inputs[3].Value())
			tDay, _ 		:= strconv.Atoi(m.inputs[4].Value())
			tHour, _ 		:= strconv.Atoi(m.inputs[5].Value())
			tMinute, _ 		:= strconv.Atoi(m.inputs[6].Value())

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


			data.AddNewProject(data.ProjectStructModel{
				ProjectName: m.inputs[0].Value(),
				Description: m.inputs[1].Value(),
				CreatedAt: currentTime,
				ExpireAt: expireAt,
				HaveExpireTime: willAddExpire,
			})
		
			main.who = main.lastOne
			main.lastOne = 3
			worldDisplay.cursor = worldDisplay.cursor + 1

			eraseProjectsInput()
		} 

	case "ctrl+c", "esc":
		main.who = main.lastOne
		main.lastOne = 3

		eraseProjectsInput()
	
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

func (m addProject) view() (string, *tea.Cursor) {
	var c *tea.Cursor

	toReturn := ""

	for i := range m.inputsCount {
		// decide where the cursor is supposed to be
		switch i {
		case m.cursor:
			c = m.inputs[i].Cursor()
		}

		// create the str to be returned
		if i < 2 {
			toReturn += fmt.Sprintf("%s\n%s\n\n", m.texts[i], m.inputs[i].View())
		} 
		if i == 2 { toReturn += fmt.Sprintf("Date Limit\n⤷ ") } 
		if i >= 2 { toReturn += m.inputs[i].View() }
	}

	if !m.dateMode {
		toReturn += "\n\n===Date mode: ADD==="
	} else {
		toReturn += "\n\n===Date mode: DATE==="
	}

	return toReturn, c
}

