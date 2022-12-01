package tools

import (
	"errors"
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

func IndexInMap[K comparable, S comparable](m map[S]K, key S) (S, error) {
	for idx := range m {
		if idx == key {
			return idx, nil
		}
	}
	var notFound S
	return notFound, errors.New("Not Found")
}

func HasToken[V comparable](s []V, v V) (bool, int) {
	for i, value := range s {
		if value == v {
			return true, i
		}
	}
	return false, -1
}
