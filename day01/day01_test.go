package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDo(t *testing.T) {
	var (
		err error
	)

	err = Do()

	assert.NoError(t, err)
}
