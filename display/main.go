package display

import tea "charm.land/bubbletea/v2"

type Daishi struct {
	who uint8
	lastOne uint8

	// This variable save the current view display who is supposed to be renderezided
	// 0 - World
	// 1 - Project
	// 2 - Delete
	// 3 - Create projects
	// 4 - Create tasks
	// 5 - Edit projects
	// 6 - Edit tasks
	// 7 - mark task or project as finished
} 

func (m Daishi) Init() tea.Cmd {
	m.who = 0
	m.lastOne = 0

	addProjectDisplay.init()
	addTaskDisplay.init()
	editProjectDisplay.init()
	editTaskDisplay.init()

	return nil
}

func (m Daishi) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch m.who {
		case 0:
			return m, worldDisplay.update(msg.String(), &m)	

		case 1:
			return m, projectDisplay.update(msg.String(), &m)

		case 2:
			return m, deleteDisplay.update(msg.String(), &m)
		
		case 3:
			return m, addProjectDisplay.update(msg.String(), msg, &m)
		
		case 4:
			return m, addTaskDisplay.update(msg.String(), msg, &m)

		case 5:
			return m, editProjectDisplay.update(msg.String(), msg, &m)

		case 6:
			return m, editTaskDisplay.update(msg.String(), msg, &m)

		case 7:
			return m, markAsFinishedDisplay.update(msg.String(), &m)
		}
	}

	return m, nil
}

func (m Daishi) View() tea.View {
	s := ""
	var c *tea.Cursor

	switch m.who {
	case 0:
		s = worldDisplay.view()

	case 1:
		s = projectDisplay.view()

	case 2:
		s = deleteDisplay.view()

	case 3:
		s, c = addProjectDisplay.view()

	case 4:
		s, c = addTaskDisplay.view()

	case 5:
		s, c = editProjectDisplay.view()

	case 6:
		s, c = editTaskDisplay.view()

	case 7:
		s = markAsFinishedDisplay.view()
	}

	v := tea.NewView(s)
	v.Cursor = c
	return v
}

