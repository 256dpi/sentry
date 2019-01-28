package main

import (
	"os"

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

	// track command
	ok := track(writer, os.Args[1], os.Args[2:]...)
	if !ok {
		// exit with error
		os.Exit(1)
	}

	// exit without error
	os.Exit(0)
}
