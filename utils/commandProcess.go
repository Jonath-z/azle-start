package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func ProcessCommand(cmd *exec.Cmd) error {
	stdin, err := cmd.StdinPipe()

	if err != nil {
		return err
	}
	defer stdin.Close()
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr

	if err = cmd.Start(); err != nil {
		fmt.Println("An error occured: ", err)
		return err
	}

	return cmd.Wait()
}
