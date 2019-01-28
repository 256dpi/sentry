package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriter(t *testing.T) {
	var cl []string
	var pl []string

	w := newWriter(func(s string) {
		cl = append(cl, s)
	}, func(s string) {
		pl = append(pl, s)
	}, nil)

	_, _ = w.Write([]byte("foo bar"))
	_, _ = w.Write([]byte("bar baz"))

	assert.Equal(t, []string{"foo bar", "bar baz"}, cl)
	assert.Equal(t, []string{"foo bar", "bar baz"}, pl)
}
