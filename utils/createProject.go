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

func CreateDefaulAzleProject(folderName string, boilerplate *string) {
	if boilerplate != nil && !helpers.Contains(availableBoilerPlates, boilerplate) {
		log.Fatal("The boilerplate specified does not exist")
	}

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
	fmt.Println("-------------------Created an azle project------------------------------")
	installDependenciesCmd := exec.Command("npm", "install")
	fmt.Println(installDependenciesCmd.ProcessState.String())

	installDependenciesCmdErr := installDependenciesCmd.Run()
	if installDependenciesCmdErr != nil {
		log.Fatal(installDependenciesCmdErr)
	}
	fmt.Println("-------------------Installed dependencies------------------------------")
}
