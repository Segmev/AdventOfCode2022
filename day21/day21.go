package day21

import (
	"fmt"
	"math"
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
	panic("Operator not found")
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

func hasHumnInstr(instrName string, numbers map[string]int, instrs map[string][]string, records *[]string) bool {
	if tools.HasKey(instrName, instrs) {
		if instrs[instrName][0] == "humn" || instrs[instrName][2] == "humn" || hasHumnInstr(instrs[instrName][0], numbers, instrs, records) || hasHumnInstr(instrs[instrName][2], numbers, instrs, records) {
			*records = append(*records, instrName)
			return true
		}
	}
	return "humn" == instrName
}

func swapInstrs(instrNameA, instrNameB string, instrs map[string][]string) {
	tmp := instrs[instrNameA]
	instrs[instrNameA] = instrs[instrNameB]
	instrs[instrNameB] = tmp
}

const RATE = 0.01

func partTwo(lines []string) {
	numbers := map[string]int{}
	instrs := map[string][]string{}
	for _, line := range lines {
		instrParts := strings.Split(line, ": ")
		tokens := strings.Split(instrParts[1], " ")
		if instrParts[0] == "root" {
			tokens[1] = "-"
		}

		if len(tokens) == 1 {
			numbers[instrParts[0]] = (tools.Atoi(tokens[0]))
		} else {
			instrs[instrParts[0]] = tokens
			if tokens[0] == "humn" {
				fmt.Println(line)
			}
		}
	}

	numbers["humn"] = 0
	previousGuessedNb := numbers["humn"]
	calcTokens(instrs["root"][2], numbers, instrs)

	guessedNb := 0
	errorDiff := math.Abs(float64(numbers[instrs["root"][2]]))
	previousErrorDiff := 1.0
	for errorDiff > 0.1 {
		copiedNumbers := tools.CloneMap(numbers)
		copiedNumbers["humn"] = guessedNb
		calcTokens("root", copiedNumbers, instrs)
		divider := int(errorDiff - previousErrorDiff)
		var gradiant int
		if divider == 0 {
			if errorDiff < previousErrorDiff {
				gradiant = 1
			} else {
				gradiant = -1
			}
		} else {
			gradiant = (guessedNb - previousGuessedNb) / int(errorDiff-previousErrorDiff)
		}

		previousGuessedNb = guessedNb
		previousErrorDiff = errorDiff
		guessedNb -= int((RATE * float64(gradiant) * errorDiff))
		copiedNumbers = tools.CloneMap(numbers)
		copiedNumbers["humn"] = guessedNb
		calcTokens("root", copiedNumbers, instrs)
		errorDiff = math.Abs(float64(copiedNumbers["root"]))
	}
	fmt.Println(guessedNb)
}

func Main(path string) {
	s := tools.Readfile(path)

	lines := strings.Split(s, "\n")
	lines = lines[:len(lines)-1]
	partOne(lines)
	partTwo(lines)
}
