package main

import (
	"log"
	// "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	payload string
}

func newModel() model {
	return model{
		payload: "DEFAULT VALUE",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if val, ok := msg.(tea.KeyMsg); ok {
		key := val.String()
		if key == "ctrl+c" {
			return m, tea.Quit
		}
		m.payload = key
	}
	return m, nil
}

func (m model) View() string {
	return m.payload
}

func main() {
	p := tea.NewProgram(newModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
