package utils

import (
	"fmt"
	"strconv"
)

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func MustLen[T any](l []T, want int) {
	if len(l) != want {
		panic(fmt.Sprintf("unexpected length of list, want %d, got %d\nitems:\n%v", want, len(l), l))
	}
}

func MustInt(s string) int {
	if i, err := strconv.Atoi(s); err != nil {
		panic(err)
	} else {
		return i
	}
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
