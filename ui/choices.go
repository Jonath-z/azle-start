package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	green = lipgloss.Color("#04B575")
	gray  = lipgloss.Color("#808080")
	white = lipgloss.Color("#fff")
)

type model struct {
	cursor   int
	choices  []string
	selected int
}

func InitialModel() model {
	m := model{
		choices:  []string{"default", "assistant-bot", "chat-completion-bot"},
		cursor:   0,
		selected: 0,
	}

	return m
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Azle starter kits")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "Q", "q":
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
			m.selected = m.cursor
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "What are you starting with?\n\n"

	for i, choice := range m.choices {
		cursor := "   "
		if m.cursor == i {
			activeCursor := lipgloss.NewStyle().Foreground(green).Render(">>>")
			activeChoice := lipgloss.NewStyle().Foreground(green).Underline(true).Render(choice)
			s += fmt.Sprintf("%s %s\n", activeCursor, activeChoice)
		} else {
			inactiveChoice := lipgloss.NewStyle().Foreground(white).Underline(false).Render(choice)
			s += fmt.Sprintf("%s %s\n", cursor, inactiveChoice)
		}
	}
	text := lipgloss.NewStyle().Foreground(gray).Render("Press q, Q, ctrl+c to quit.")
	s += "\n" + text + "\n"

	return s
}
