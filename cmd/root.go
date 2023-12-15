/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "azle-start",
	Short: "Welcome to azle-start.",
	Long:  "Welcome to azle-start, your tool for quickly and easily starting with azle.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			folderName := args[0]
			createDefaulAzleProject(folderName)
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createDefaulAzleProject(folderName string) {
	fmt.Println("Create a azle project")
}
