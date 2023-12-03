package utils

import (
	"slices"
	"strings"
)

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func ReserveString(str string) string {
	bytes := []byte(str)
	slices.Reverse(bytes)
	return string(bytes)
}

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func Lines(str string) []string {
	return strings.Split(str, "\n")
}
