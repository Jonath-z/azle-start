package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func ProcessCommand(cmd *exec.Cmd) {
	stdin, err := cmd.StdinPipe()

	if err != nil {
		fmt.Println(err)
	}
	defer stdin.Close()
	buf := new(bytes.Buffer) // THIS STORES THE NODEJS OUTPUT
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr

	if err = cmd.Start(); err != nil {
		fmt.Println("An error occured: ", err)
	}

	cmd.Wait()
}
