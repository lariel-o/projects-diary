package display

import tea "charm.land/bubbletea/v2"

// This variable save the current view display who is supposed to be renderezided
var who uint8

type Daishi struct {
} 

func (m Daishi) Init() tea.Cmd {
	who = 0
	return nil
}

func (m Daishi) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch who {
		case 0:
			// the update function at worlddisplay will deal with everything
			return m, worldDisplay.update(msg.String(), &m)	

		case 1:
			return m, projectDisplay.update(msg.String(), &m)
		}
	}

	return m, nil
}

func (m Daishi) View() tea.View {
	s := ""

	switch who {
	case 0:
		s = worldDisplay.view()

	case 1:
		s = projectDisplay.view()
	}

	return tea.NewView(s)
}

