package cmd

import (
	"fmt"
	"os"

	"github.com/Jonath-z/azle-start/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "azle-start",
	Short: "Welcome to azle-start.",
	Long:  "Welcome to azle-start, your tool for quickly and easily starting with azle.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			p := tea.NewProgram(ui.InitialModel())
			_, err := p.Run()
			if err != nil {
				fmt.Printf("Alas, there's been an error: %v", err)
				os.Exit(1)
			}

		}
	},
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.Flags().BoolP("help", "h", false, "Help")
}
