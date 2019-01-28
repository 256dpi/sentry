package main

import (
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

	// create filter
	filter := newFilter(os.Getenv("FILTER"))

	// create capturer
	capturer := func(s string) {
		raven.CaptureMessageAndWait(s, nil)
	}

	// create printer
	printer := func(s string) {
		println(s)
	}

	// create writer
	writer := newWriter(capturer, printer, filter)

	// create command
	cmd := exec.Command(os.Args[1], os.Args[2:]...)

	// set standard out to current process
	cmd.Stdout = os.Stdout

	// set standard in to current process
	cmd.Stdin = os.Stdin

	// error tracker
	cmd.Stderr = writer

	// run command
	err := cmd.Run()
	if err == nil {
		// exit immediately if everything did go well
		os.Exit(0)
	}

	// write exec error
	_, _ = cmd.Stderr.Write([]byte(err.Error()))

	// exit with error
	os.Exit(1)
}
