package tools

import (
	"os"
	"strconv"
)

func CloneSlice[T any](s []T) []T {
	res := make([]T, len(s))
	copy(res, s)
	return res
}

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

func HasKey[K comparable, V any](key K, m map[K]V) bool {
	_, ok := m[key]
	return ok
}

func Abs(nb int) int {
	if nb < 0 {
		return -nb
	}
	return nb
}

func Modulo(a, b int) int {
	return ((a % b) + b) % b
}

func Atoi(s string) int {
	res, _ := strconv.Atoi(s)
	return res
}

func CloneMap[k comparable, v any](m map[k]v) map[k]v {
	n := map[k]v{}
	for k, v := range m {
		n[k] = v
	}
	return n
}
