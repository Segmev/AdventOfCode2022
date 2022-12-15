package tools

import (
	"os"
)

func Readfile(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func IndexInSlice[K comparable](s []K, value K) int {
	for idx, v := range s {
		if v == value {
			return idx
		}
	}
	return -1
}

func HasToken[V comparable](s []V, v V) (bool, int) {
	for i, value := range s {
		if value == v {
			return true, i
		}
	}
	return false, -1
}

func Abs(nb int) int {
	if nb < 0 {
		return -nb
	}
	return nb
}
