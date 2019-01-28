package main

import (
	"regexp"
	"strings"
)

func newFilter(config string) func(string) bool {
	// prepare segments
	var segments []string

	// parse segments if available
	if config != "" {
		segments = strings.Split(config, ";")
	}

	// prepare patterns
	patterns := make([]*regexp.Regexp, 0, len(segments))

	// compile all patterns
	for _, seg := range segments {
		patterns = append(patterns, regexp.MustCompile(seg))
	}

	// create filter
	filter := func(str string) bool {
		for _, p := range patterns {
			if p.MatchString(str) {
				return true
			}
		}

		return false
	}

	return filter
}
