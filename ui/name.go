package ui

import (
	"fmt"
	"unicode"

	colors "github.com/Jonath-z/azle-start/ui/Colors"
	"github.com/Jonath-z/azle-start/utils"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type nameState struct {
	name        string
	boilerplate string
}

func InitializeNameState(boilerplate string) nameState {
	state := nameState{
		name:        "",
		boilerplate: boilerplate,
	}

	return state
}

func (state nameState) Init() tea.Cmd {
	return tea.SetWindowTitle("Select project name")
}

func (state nameState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if len(msg.Runes) > 0 {
			if !unicode.IsSymbol(msg.Runes[0]) {
				state.name += msg.String()
			}
		}

		switch msg.String() {
		case "ctrl+c":
			return state, tea.Quit
		case "backspace":
			if state.name != "" {
				state.name = state.name[:len(state.name)-1]
			}
		case "enter":
			if len(state.name) >= 3 {
				fmt.Println(state.boilerplate)
				utils.CreateAzleProject(state.name, &state.boilerplate)
			}
		}
	}

	return state, nil
}

func (state nameState) View() string {
	s := "What's the project name: "
	projectName := lipgloss.NewStyle().Foreground(colors.Green).Render(state.name)
	return s + projectName
}
