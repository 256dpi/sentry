package main

import (
	"bufio"
	"io"
)

type writer struct {
	capturer func(string)
	printer  func(string)
	filter   func(string) bool

	in  *io.PipeWriter
	out *io.PipeReader
}

func newWriter(capturer, printer func(string), filter func(string) bool) *writer {
	// create pipe
	out, in := io.Pipe()

	// create writer
	w := &writer{
		capturer: capturer,
		printer:  printer,
		filter:   filter,

		in:  in,
		out: out,
	}

	// run processor
	go w.processor()

	return w
}

func (w *writer) Write(data []byte) (int, error) {
	// write to pipe
	return w.in.Write(data)
}

func (w *writer) processor() {
	// create scanner
	scanner := bufio.NewScanner(w.out)

	// process all lines
	for scanner.Scan() {
		// get string
		str := scanner.Text()

		// check filter
		filtered := false
		if w.filter != nil {
			filtered = w.filter(str)
		}

		// capture if not filtered
		if !filtered {
			w.capturer(str)
		}

		// print if provided
		if w.printer != nil {
			w.printer(str)
		}
	}
}

func (w *writer) close() {
	_ = w.in.Close()
}
