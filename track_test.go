package main

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

var lsErr = []string{"ls: foo-bar-baz: No such file or directory\n", "exit status 1"}
var exitErr = []string{"exit status 1"}

func init() {
	if runtime.GOOS == "linux" {
		lsErr = []string{"ls: cannot access foo-bar-baz", ": No such file or directory\n", "exit status 2"}
		exitErr = []string{"exit status 2"}
	}
}

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
	assert.Equal(t, lsErr, cl)
	assert.Equal(t, lsErr, pl)
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
	assert.Equal(t, exitErr, cl)
	assert.Equal(t, lsErr, pl)
}
