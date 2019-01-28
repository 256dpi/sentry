package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	f := newFilter("foo bar")

	assert.True(t, f("baz foo bar"))
	assert.True(t, f("foo bar baz"))
	assert.False(t, f("baz bar"))
}

func TestFilterEmpty(t *testing.T) {
	f := newFilter("")

	assert.False(t, f("foo"))
	assert.False(t, f("foo bar"))
	assert.False(t, f(""))
}

func TestFilterMultiple(t *testing.T) {
	f := newFilter("foo bar;bar baz")

	assert.True(t, f("foo bar"))
	assert.True(t, f("bar baz"))
	assert.False(t, f("bar foo"))
	assert.False(t, f("baz bar"))
}

func TestFilterAdvanced(t *testing.T) {
	f := newFilter("[a-c];[x-z]")

	assert.True(t, f("ahl"))
	assert.False(t, f("nfl"))
	assert.True(t, f("xkl"))
}
