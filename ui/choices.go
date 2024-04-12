package ui

import (
	"fmt"
	"log"
	"os"

	"github.com/Jonath-z/azle-start/helpers"
	colors "github.com/Jonath-z/azle-start/ui/Colors"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

type model struct {
	cursor         int
	choices        []string
	selected       int
	scrollPosition int
}

func InitialModel() model {
	return model{
		choices:        helpers.GetExamplesList(),
		cursor:         0,
		selected:       0,
		scrollPosition: 0,
	}
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
				if m.cursor < m.scrollPosition {
					m.scrollPosition = m.cursor
				}
			}
		case "down", "j":
			_, terminalHeight, err := term.GetSize(int(os.Stdout.Fd()))
			if err != nil {
				log.Println("Error getting terminal height", err)
			} else {
				if m.cursor < len(m.choices)-1 {
					m.cursor++
					if m.cursor >= m.scrollPosition+terminalHeight-3 {
						m.scrollPosition++
					}
				}
			}
		case "enter", " ":
			m.selected = m.cursor
			nameState := tea.NewProgram(InitializeNameState(m.choices[m.selected]))
			_, err := nameState.Run()
			if err != nil {
				log.Fatal("Error while initializing the name layout")
				os.Exit(1)
			}
			return m, nil
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "What are you starting with?\n\n"

	visibleChoices := make([]string, 0)
	_, terminalHeight, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Println("Error getting terminal size:", err)
	}
	startIndex := m.scrollPosition
	endIndex := startIndex + terminalHeight - 4

	for i := startIndex; i < endIndex && i < len(m.choices); i++ {
		visibleChoices = append(visibleChoices, m.choices[i])
	}

	for i, choice := range visibleChoices {
		cursor := "   "
		if m.cursor == i {
			activeCursor := lipgloss.NewStyle().Foreground(colors.Green).Render(">>>")
			activeChoice := lipgloss.NewStyle().Foreground(colors.Green).Underline(true).Render(choice)
			s += fmt.Sprintf("%s %s\n", activeCursor, activeChoice)
		} else {
			inactiveChoice := lipgloss.NewStyle().Foreground(colors.White).Underline(false).Render(choice)
			s += fmt.Sprintf("%s %s\n", cursor, inactiveChoice)
		}
	}
	text := lipgloss.NewStyle().Foreground(colors.Gray).Render("Press q, Q, ctrl+c to quit.")
	s += "\n" + text + "\n"

	return s
}
