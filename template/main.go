package main

import (
	"strings"

	"github.com/Segmev/AdventOfCode2022/tools"
)

func partOne(lines []string) {
}

func partTwo(lines []string) {
}

func Main(path string) {
	s := tools.Readfile(path)

	lines := strings.Split(s, "\n")
	partOne(lines)
	partTwo(lines)
}
