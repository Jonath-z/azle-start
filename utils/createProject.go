package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/Jonath-z/azle-start/helpers"
)

var availableBoilerPlates = helpers.GetExamplesList()

func CreateAzleProject(folderName string, boilerplate *string) {
	if boilerplate != nil && !helpers.Contains(availableBoilerPlates, *boilerplate) {
		log.Fatal("The boilerplate specified does not exist")
		os.Exit(1)
	}

	var initializedProjectPath = ""
	path, err := os.Getwd()
	fmt.Println(path)
	if err != nil {
		fmt.Println("Reached here")
		log.Fatal(err.Error())
		os.Exit(1)
	}

	if folderName == "." {
		pathComponents := strings.Split(path, "/")
		currentDir := pathComponents[len(pathComponents)-1]
		initializedProjectPath = path + "/" + currentDir
	} else {
		initializedProjectPath = path + "/" + folderName
	}

	boilerplatePath := path + "/" + "azle/examples/" + *boilerplate
	cmd := exec.Command("cp", "-r", boilerplatePath, initializedProjectPath)
	initializedProjectErr := cmd.Run()
	if initializedProjectErr != nil {
		fmt.Println(initializedProjectErr.Error())
		os.Exit(1)
	}

	checkoutToProjectDirErr := os.Chdir(initializedProjectPath)
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

	fmt.Println("-------------------Dependencies installed------------------------------")
	os.Exit(0)
}
