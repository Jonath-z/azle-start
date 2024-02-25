package script

import (
	"log"
	"os"
)

func GetExamplesList() []string {
	azleExamplesPath := "./azle/examples"
	var examples []string

	azleExamples, err := os.ReadDir(azleExamplesPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, path := range azleExamples {
		examples = append(examples, path.Name())
	}

	return examples
}
