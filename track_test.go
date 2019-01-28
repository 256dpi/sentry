package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrack(t *testing.T) {
	var cl []string
	var pl []string

	w := newWriter(func(s string) {
		cl = append(cl, s)
	}, func(s string) {
		pl = append(pl, s)
	}, nil)

	track(w, "ls", "-al")
	assert.Equal(t, []string(nil), cl)
	assert.Equal(t, []string(nil), pl)
}

func TestTrackError(t *testing.T) {
	var cl []string
	var pl []string

	w := newWriter(func(s string) {
		cl = append(cl, s)
	}, func(s string) {
		pl = append(pl, s)
	}, nil)

	track(w, "ls", "foo-bar-baz")
	assert.Equal(t, []string{"ls: foo-bar-baz: No such file or directory\n", "exit status 1"}, cl)
	assert.Equal(t, []string{"ls: foo-bar-baz: No such file or directory\n", "exit status 1"}, pl)
}

func TestTrackErrorFilter(t *testing.T) {
	var cl []string
	var pl []string

	w := newWriter(func(s string) {
		cl = append(cl, s)
	}, func(s string) {
		pl = append(pl, s)
	}, newFilter("No such file"))

	track(w, "ls", "foo-bar-baz")
	assert.Equal(t, []string{"exit status 1"}, cl)
	assert.Equal(t, []string{"ls: foo-bar-baz: No such file or directory\n", "exit status 1"}, pl)
}
