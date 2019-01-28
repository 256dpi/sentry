package main

import (
	"io"
	"os"
	"os/exec"

	"github.com/armon/circbuf"
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

	// circular buffer capped at 2K
	buf, _ := circbuf.NewBuffer(2000)

	// create a multi writer
	writer := io.MultiWriter(os.Stderr, buf)

	// create command
	cmd := exec.Command(os.Args[1], os.Args[2:]...)

	// set standard out to current process
	cmd.Stdout = os.Stdout

	// set standard in to current process
	cmd.Stdin = os.Stdin

	// route standard error to multi writer
	cmd.Stderr = writer

	// run command
	err := cmd.Run()
	if err == nil {
		// exit immediately if everything did go well
		os.Exit(0)
	} else if _, ok := err.(*exec.ExitError); !ok {
		// append error to Stderr
		_, _ = io.WriteString(cmd.Stderr, err.Error())
	}

	// send report to sentry
	raven.CaptureMessageAndWait(buf.String(), nil)

	// print buffer
	print(buf.String())

	// exit with error
	os.Exit(1)
}
