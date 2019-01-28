package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	f := newFilter("foo bar")

	assert.True(t, f.match("baz foo bar"))
	assert.True(t, f.match("foo bar baz"))
	assert.False(t, f.match("baz bar"))
}

func TestFilterEmpty(t *testing.T) {
	f := newFilter("")

	assert.False(t, f.match("foo"))
	assert.False(t, f.match("foo bar"))
	assert.False(t, f.match(""))
}

func TestFilterMultiple(t *testing.T) {
	f := newFilter("foo bar;bar baz")

	assert.True(t, f.match("foo bar"))
	assert.True(t, f.match("bar baz"))
	assert.False(t, f.match("bar foo"))
	assert.False(t, f.match("baz bar"))
}

func TestFilterAdvanced(t *testing.T) {
	f := newFilter("[a-c];[x-z]")

	assert.True(t, f.match("ahl"))
	assert.False(t, f.match("nfl"))
	assert.True(t, f.match("xkl"))
}
