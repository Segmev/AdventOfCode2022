package day06

import (
	"fmt"
	"strings"

	"github.com/Segmev/AdventOfCode2022/tools"
)

func isMarker(marker string) bool {
	markerSet := tools.GetSet[rune]()
	for _, r := range marker {
		if markerSet.Contains(r) {
			return false
		}
		markerSet.Add(r)
	}
	return true
}

func partOne(lines []string) {
	line := lines[0]

	for i := 4; i < len(line); i++ {
		if isMarker(line[i-4 : i]) {
			fmt.Println(i)
			break
		}
	}
}

func partTwo(lines []string) {
	line := lines[0]

	for i := 14; i < len(line); i++ {
		if isMarker(line[i-14 : i]) {
			fmt.Println(i)
			break
		}
	}
}

func Main(path string) {
	s := tools.Readfile(path)

	lines := strings.Split(s, "\n")
	partOne(lines)
	partTwo(lines)
}
