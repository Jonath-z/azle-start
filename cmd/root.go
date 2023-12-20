package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "azle-start",
	Short: "Welcome to azle-start.",
	Long:  "Welcome to azle-start, your tool for quickly and easily starting with azle.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			folderName := args[0]
			createDefaulAzleProject(folderName)
		} else {
			createDefaulAzleProject(".")
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
	var initialzedProjectPath = ""
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if folderName == "." {
		pathComponents := strings.Split(path, "/")
		currentDir := pathComponents[len(pathComponents)-1]
		initialzedProjectPath = path + "/" + currentDir
	} else {
		initialzedProjectPath = path + "/" + folderName
	}

	defaultAzleProjectPath := "./starter-kits/default"
	cmd := exec.Command("cp", "-r", defaultAzleProjectPath, initialzedProjectPath)
	initializedProjectErr := cmd.Run()
	if initializedProjectErr != nil {
		log.Fatal(initializedProjectErr)
	}
	fmt.Println("-------------------Created a azle project------------------------------")
	fmt.Println("-------------------Installing Dependencies------------------------------")
	installDependenciesCmd := exec.Command("npm", "install")
	fmt.Println(installDependenciesCmd.ProcessState.String())

	installDependenciesCmdErr := installDependenciesCmd.Run()
	if installDependenciesCmdErr != nil {
		log.Fatal(installDependenciesCmdErr)
	}
	fmt.Println("-------------------Installed dependencies------------------------------")
}
