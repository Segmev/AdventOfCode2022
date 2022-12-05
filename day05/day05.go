package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Segmev/AdventOfCode2022/tools"
)

func partOne(lines []string) {
	piles, instrs := getStackAndInstruction(lines)

	for _, instr := range instrs {
		instParts := strings.Split(instr, " ")

		max, _ := strconv.Atoi(instParts[1])
		for i := 0; i < max; i++ {
			pileLen := len(piles[instParts[3]])
			piles[instParts[5]] = append(piles[instParts[5]], piles[instParts[3]][pileLen-1])
			piles[instParts[3]] = piles[instParts[3]][:pileLen-1]
		}
	}
	for i := 0; i < len(piles); i++ {
		pi := strconv.Itoa(i + 1)
		fmt.Print(piles[pi][len(piles[pi])-1])
	}
	fmt.Println()
}

func partTwo(lines []string) {
	piles, instrs := getStackAndInstruction(lines)

	for _, instr := range instrs {
		instParts := strings.Split(instr, " ")

		max, _ := strconv.Atoi(instParts[1])
		pileLen := len(piles[instParts[3]])
		piles[instParts[5]] = append(piles[instParts[5]], piles[instParts[3]][pileLen-max:pileLen]...)
		piles[instParts[3]] = piles[instParts[3]][:pileLen-max]

	}

	for i := 0; i < len(piles); i++ {
		pi := strconv.Itoa(i + 1)
		fmt.Print(piles[pi][len(piles[pi])-1])
	}
	fmt.Println()
}

func getStackAndInstruction(lines []string) (map[string][]string, []string) {
	stackInstEnd := 0
	for i := range lines {
		if len(lines[i]) == 0 {
			stackInstEnd = i - 2
			break
		}
	}

	piles := map[string][]string{}
	for _, stackId := range lines[stackInstEnd+1] {
		if stackId != rune(' ') {
			piles[string(stackId)] = []string{}
		}
	}

	for _, line := range lines[:stackInstEnd+1] {
		for i := 0; i <= len(line)-3; i += 4 {
			if line[i+1:i+2] != " " {
				pileIdx := strconv.Itoa((i / 4) + 1)
				piles[pileIdx] = append([]string{line[i+1 : i+2]}, piles[pileIdx]...)
			}
		}
	}

	instrs := []string{}
	for _, line := range lines[stackInstEnd+2:] {
		if len(line) > 0 {
			instrs = append(instrs, line)
		}
	}
	return piles, instrs
}

func main() {
	s := tools.Readfile("./input.txt")

	lines := strings.Split(s, "\n")
	partOne(lines)
	partTwo(lines)
}
