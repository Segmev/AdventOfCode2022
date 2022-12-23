package day21

import (
	"fmt"
	"strings"

	"github.com/Segmev/AdventOfCode2022/tools"
)

func computeNbs(fnb, snb int, ope string) int {
	switch ope {
	case "+":
		return fnb + snb
	case "-":
		return fnb - snb
	case "*":
		return fnb * snb
	case "/":
		return fnb / snb
	}
	return 0
}

func calcTokens(instrName string, numbers map[string]int, instrs map[string][]string) {
	if tools.HasKey(instrName, numbers) {
		return
	}
	var fnb, snb int
	if !tools.HasKey(instrs[instrName][0], numbers) {
		calcTokens(instrs[instrName][0], numbers, instrs)
	}
	if !tools.HasKey(instrs[instrName][2], numbers) {
		calcTokens(instrs[instrName][2], numbers, instrs)
	}
	fnb = numbers[instrs[instrName][0]]
	snb = numbers[instrs[instrName][2]]
	numbers[instrName] = computeNbs(fnb, snb, instrs[instrName][1])
}

func partOne(lines []string) {
	numbers := map[string]int{}
	instrs := map[string][]string{}
	for _, line := range lines {
		instrParts := strings.Split(line, ": ")

		tokens := strings.Split(instrParts[1], " ")
		if len(tokens) == 1 {
			numbers[instrParts[0]] = tools.Atoi(tokens[0])
		} else {
			instrs[instrParts[0]] = tokens
		}
	}

	calcTokens("root", numbers, instrs)
	fmt.Println(numbers["root"])
}

func partTwo(lines []string) {
}

func Main(path string) {
	s := tools.Readfile(path)

	lines := strings.Split(s, "\n")
	lines = lines[:len(lines)-1]
	partOne(lines)
	partTwo(lines)
}
