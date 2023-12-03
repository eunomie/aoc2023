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

	res, err = Do(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`)

	assert.NoError(t, err)
	assert.Equal(t, 142, res)
}
