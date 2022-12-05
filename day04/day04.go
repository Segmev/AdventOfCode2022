package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Segmev/AdventOfCode2022/tools"
)

func areBoundsOverlapping(fBounds []int, sBounds []int) bool {
	lowestBounds, highestBounds := fBounds, sBounds
	if sBounds[0] < fBounds[0] {
		lowestBounds = sBounds
		highestBounds = fBounds
	}
	return highestBounds[0] <= lowestBounds[1]
}

func areBoundsContained(fBounds []int, sBounds []int) bool {
	smallerBounds, biggerBounds := fBounds, sBounds
	if smallerBounds[1]-smallerBounds[0] > biggerBounds[1]-biggerBounds[0] {
		smallerBounds = sBounds
		biggerBounds = fBounds
	}
	return smallerBounds[0] >= biggerBounds[0] && smallerBounds[1] <= biggerBounds[1]
}

func getIntBounds(elfBounds string) []int {
	fElfBounds := strings.Split(elfBounds, "-")
	lBound, _ := strconv.Atoi(fElfBounds[0])
	rBound, _ := strconv.Atoi(fElfBounds[1])
	return []int{lBound, rBound}
}

func partOne(lines []string) {
	count := 0
	for _, line := range lines {
		elves := strings.Split(line, ",")
		if len(elves) >= 2 {
			fBounds := getIntBounds(elves[0])
			sBounds := getIntBounds(elves[1])
			if areBoundsContained(fBounds, sBounds) {
				count += 1
			}
		}
	}
	fmt.Println(count)
}

func partTwo(lines []string) {
	count := 0
	for _, line := range lines {
		elves := strings.Split(line, ",")
		if len(elves) >= 2 {
			fBounds := getIntBounds(elves[0])
			sBounds := getIntBounds(elves[1])
			if areBoundsOverlapping(fBounds, sBounds) {
				count += 1
			}
		}
	}
	fmt.Println(count)
}

func main() {
	s := tools.Readfile("./input.txt")

	lines := strings.Split(s, "\n")
	partOne(lines)
	partTwo(lines)
}
