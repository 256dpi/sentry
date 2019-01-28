package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriter(t *testing.T) {
	cq := make(chan string, 2)
	pq := make(chan string, 2)

	w := newWriter(func(s string) {
		cq <- s
	}, func(s string) {
		pq <- s
	}, nil)

	_, _ = w.Write([]byte("foo bar\n"))
	_, _ = w.Write([]byte("bar baz\n"))

	c1 := <-cq
	c2 := <-cq
	assert.Equal(t, "foo bar", c1)
	assert.Equal(t, "bar baz", c2)
	assert.Len(t, cq, 0)

	p1 := <-pq
	p2 := <-pq
	assert.Equal(t, "foo bar", p1)
	assert.Equal(t, "bar baz", p2)
	assert.Len(t, pq, 0)
}
