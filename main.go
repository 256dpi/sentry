package main

import (
	"os"
	"os/exec"

	"github.com/getsentry/raven-go"
)

var skipper *filter

func main() {
	// check dsn
	if os.Getenv("SENTRY_DSN") == "" {
		panic("missing SENTRY_DSN")
	}

	// check arguments
	if len(os.Args) < 2 {
		panic("missing arguments")
	}

	// create skipper
	skipper = newFilter(os.Getenv("SKIP"))

	// create command
	cmd := exec.Command(os.Args[1], os.Args[2:]...)

	// set standard out to current process
	cmd.Stdout = os.Stdout

	// set standard in to current process
	cmd.Stdin = os.Stdin

	// error tracker
	cmd.Stderr = &errorWriter{}

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

type errorWriter struct{}

func (w *errorWriter) Write(data []byte) (int, error) {
	// get string
	str := string(data)

	// check filter
	filtered := skipper.match(str)

	// capture error if not filtered
	if !filtered {
		raven.CaptureMessageAndWait(str, nil)
	}

	// log error
	println(str)

	return len(data), nil
}
