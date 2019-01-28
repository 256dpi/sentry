package main

import (
	"io"
	"os"
	"os/exec"
)

func track(w *writer, name string, args ...string) bool {
	// create command
	cmd := exec.Command(name, args...)

	// set standard out to current process
	cmd.Stdout = os.Stdout

	// set standard in to current process
	cmd.Stdin = os.Stdin

	// error tracker
	cmd.Stderr = w

	// run command
	err := cmd.Run()
	if err == nil {
		return true
	}

	// write exec error
	_, _ = io.WriteString(cmd.Stderr, err.Error())

	return false
}
