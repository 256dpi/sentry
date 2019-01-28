package main

// TODO: Use scanner to read line by line?

type writer struct {
	capturer func(string)
	printer  func(string)
	filter   func(string) bool
}

func newWriter(capturer, printer func(string), filter func(string) bool) *writer {
	return &writer{
		capturer: capturer,
		printer:  printer,
		filter:   filter,
	}
}

func (w *writer) Write(data []byte) (int, error) {
	// get string
	str := string(data)

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

	return len(data), nil
}

func (w *writer) close() {

}
