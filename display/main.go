package display

import tea "charm.land/bubbletea/v2"

type Daishi struct {
	CurrentView uint8
} 

func (m Daishi) Init() tea.Cmd {
	return nil
}

func (m Daishi) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch m.CurrentView {
		case 0:
			return m, worldDisplay.update(msg.String())	
		}
	}

	return m, nil
}

func (m Daishi) View() tea.View {
	s := ""

	return tea.NewView(s)
}

