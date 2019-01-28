package main

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

var lsErr = "ls: foo-bar-baz: No such file or directory\n"
var exitErr = "exit status 1"

func init() {
	if runtime.GOOS == "linux" {
		lsErr = "ls: cannot access foo-bar-baz: No such file or directory"
		exitErr = "exit status 2"
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
	assert.Equal(t, []string{lsErr, exitErr}, cl)
	assert.Equal(t, []string{lsErr, exitErr}, pl)
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
	assert.Equal(t, []string{exitErr}, cl)
	assert.Equal(t, []string{lsErr, exitErr}, pl)
}
