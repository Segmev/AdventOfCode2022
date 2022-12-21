package tools

import (
	"os"
	"strconv"
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
