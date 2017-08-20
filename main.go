package main

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/getsentry/raven-go"
)

func main() {
	// get dsn
	dsn := os.Getenv("SENTRY_DSN")
	if dsn == "" {
		panic("missing SENTRY_DSN")
	}

	// check arguments
	if len(os.Args) < 2 {
		panic("missing arguments")
	}

	// create command
	cmd := exec.Command(os.Args[1], os.Args[2:]...)

	// set standard out to current process
	cmd.Stdout = os.Stdout

	// set standard in to current process
	cmd.Stdin = os.Stdin

	// error tracker
	buf := bytes.NewBuffer(nil)
	cmd.Stderr = buf

	// run command
	err := cmd.Run()
	if _, ok := err.(*exec.ExitError); ok {
		// send report to sentry
		raven.CaptureMessageAndWait(buf.String(), nil)

		// print message again
		print(buf.String())

		os.Exit(1)
	} else if err != nil {
		panic(err)
	}
}
