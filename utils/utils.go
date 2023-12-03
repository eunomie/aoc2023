package utils

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func OrIndex(i, j int) int {
	if i < 0 {
		return j
	}
	return i
}

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
