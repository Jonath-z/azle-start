package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor   int
	choices  []string
	selected int
}

func InitialModel() model {
	return model{
		choices:  []string{"default", "assistant-bot", "chat-completion-bot"},
		selected: 0,
	}
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Azle starter kits")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctl+c", "Q", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			fmt.Println(m.cursor)
			m.selected = m.cursor
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "What are you starting with?\n\n"

	for i, choice := range m.choices {
		cursor := "   "
		if m.cursor == i {
			cursor = ">>>"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}
