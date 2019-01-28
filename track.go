package main

import (
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
		// close writer
		w.close()

		return true
	}

	// write exec error
	_, _ = cmd.Stderr.Write([]byte(err.Error()))

	// close writer
	w.close()

	return false
}
