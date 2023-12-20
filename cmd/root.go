package cmd

import (
	"fmt"
	"os"

	"github.com/Jonath-z/azle-start/ui"
	"github.com/Jonath-z/azle-start/utils"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "azle-start",
	Short: "Welcome to azle-start.",
	Long:  "Welcome to azle-start, your tool for quickly and easily starting with azle.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			p := tea.NewProgram(ui.InitialModel())
			model, err := p.Run()
			if err != nil {
				fmt.Printf("Alas, there's been an error: %v", err)
				os.Exit(1)
			}

			fmt.Println("the model", model)
			// createDefaulAzleProject()
		} else {
			utils.CreateDefaulAzleProject(".")
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("help", "h", false, "Help")
}
