package main

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/getsentry/raven-go"
)

func main() {
	// check dsn
	if os.Getenv("SENTRY_DSN") == "" {
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
	if err == nil {
		// exit immediately if everything did go well
		os.Exit(0)
	} else if _, ok := err.(*exec.ExitError); !ok {
		// write any other error to buffer
		buf.WriteString(err.Error())
	}

	// send report to sentry
	raven.CaptureMessageAndWait(buf.String(), nil)

	// print message again
	print(buf.String())

	// exit with error
	os.Exit(1)
}
