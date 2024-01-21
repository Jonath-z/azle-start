package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/Jonath-z/azle-start/helpers"
)

var availableBoilerPlates = []string{"default", "chat-completion-bot", "assistant-deBot"}

func CreateAzleProject(folderName string, boilerplate *string) {
	if boilerplate != nil && !helpers.Contains(availableBoilerPlates, *boilerplate) {
		log.Fatal("The boilerplate specified does not exist")
		os.Exit(1)
	}

	var initialzedProjectPath = ""
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	if folderName == "." {
		pathComponents := strings.Split(path, "/")
		currentDir := pathComponents[len(pathComponents)-1]
		initialzedProjectPath = path + "/" + currentDir
	} else {
		initialzedProjectPath = path + "/" + folderName
	}

	defaultAzleProjectPath := path + "/" + "starter-kits/" + *boilerplate
	cmd := exec.Command("cp", "-r", defaultAzleProjectPath, initialzedProjectPath)
	initializedProjectErr := cmd.Run()
	if initializedProjectErr != nil {
		fmt.Println(initializedProjectErr.Error())
		os.Exit(1)
	}
	fmt.Println("-------------------Created an azle project------------------------------")
	checkoutToProjectDirErr := os.Chdir(initialzedProjectPath)
	if checkoutToProjectDirErr != nil {
		log.Fatal("Faile to checkout to the project directory", checkoutToProjectDirErr.Error())
		os.Exit(1)
	}
	installDependenciesCmd := exec.Command("npm", "install")
	depsErr := ProcessCommand(installDependenciesCmd)
	if depsErr != nil {
		log.Fatal("Failed to install dependencies")
		os.Exit(1)
	}
	fmt.Println("-------------------Installed dependencies------------------------------")
	os.Exit(0)
}
