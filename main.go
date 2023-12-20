package main

import (
	"fmt"
	"os"

	"github.com/Jonath-z/azle-start/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// cmd.Execute()
	p := tea.NewProgram(ui.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}
