package main

import (
	"fmt"
	"strings"

	"github.com/Segmev/AdventOfCode2022/tools"
)

func partOne(lines []string) {
	//      X: rock       Y: paper     Z: scissors
	roundsValues := map[string]map[string][]int{
		"A": {"X": {1, 3}, "Y": {2, 6}, "Z": {3, 0}}, // rock
		"B": {"X": {1, 0}, "Y": {2, 3}, "Z": {3, 6}}, // paper
		"C": {"X": {1, 6}, "Y": {2, 0}, "Z": {3, 3}}, // scissors
	}

	myScore := 0
	for _, line := range lines {
		tokens := strings.Split(line, " ")

		if len(tokens) > 1 {
			myScore += roundsValues[tokens[0]][tokens[1]][0] + roundsValues[tokens[0]][tokens[1]][1]
		}
	}
	fmt.Println(myScore)
}

func partTwo(lines []string) {
	// X: loose, Y: draw, Z: win
	moves := []string{"A", "B", "C"}
	endStyle := map[string][]int{"X": {-1, 0}, "Y": {0, 3}, "Z": {1, 6}}
	moveVals := map[string]int{moves[0]: 1, moves[1]: 2, moves[2]: 3}

	myScore := 0
	for _, line := range lines {
		tokens := strings.Split(line, " ")

		if len(tokens) > 1 {
			key := tools.IndexInSlice(moves, tokens[0])
			if key >= 0 {
				myStrat := endStyle[tokens[1]]
				moveKey := (key + myStrat[0]) % 3
				if moveKey < 0 {
					moveKey = 2
				}
				myMove := moves[moveKey]
				myScore += moveVals[myMove] + myStrat[1]
			}
		}
	}
	fmt.Println(myScore)
}

func main() {
	s := tools.Readfile("./input.txt")

	lines := strings.Split(s, "\n")
	partOne(lines)
	partTwo(lines)
}
