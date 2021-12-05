package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	if len(cmd) == 0 {
		return 1
	}

	for k, v := range env {
		err := os.Unsetenv(k)
		if err != nil {
			log.Println(err)
			return 1
		}

		if !v.NeedRemove {
			err = os.Setenv(k, v.Value)
			if err != nil {
				log.Println(err)
				return 1
			}
		}
	}

	var args []string
	if len(cmd) > 1 {
		args = cmd[1:]
	}

	commandName := cmd[0]

	command := exec.Command(commandName, args...)

	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		var e *exec.ExitError
		if errors.As(err, &e) {
			return e.ExitCode()
		}
	}

	return command.ProcessState.ExitCode()
}
