package main

import (
	"regexp"
	"strings"
)

type filter struct {
	patterns []*regexp.Regexp
}

func newFilter(config string) *filter {
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
	filter := &filter{
		patterns: patterns,
	}

	return filter
}

func (f *filter) match(str string) bool {
	for _, p := range f.patterns {
		if p.MatchString(str) {
			return true
		}
	}

	return false
}
