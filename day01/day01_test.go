package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDo(t *testing.T) {
	var (
		err error
		res int
	)

	res, err = Do(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`)

	assert.NoError(t, err)
	assert.Equal(t, 281, res)
}
