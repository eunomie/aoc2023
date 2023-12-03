package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDo(t *testing.T) {
	var (
		res      int
		in       = ``
		expected = 0
	)

	res = Do(in)

	assert.Equal(t, expected, res)
}
