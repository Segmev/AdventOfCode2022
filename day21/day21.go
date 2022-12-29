package day21

import (
	"fmt"
	"strings"

	"github.com/Segmev/AdventOfCode2022/tools"
)

func inverseOpeLeftX(ope string) string {
	switch ope {
	case "+":
		return "-"
	case "-":
		return "+"
	case "*":
		return "/"
	case "/":
		return "*"
	}
	panic("Operator not found")
}

func inverseOpeRightX(ope string) string {
	switch ope {
	case "+":
		return "-"
	case "-":
		return "-"
	case "*":
		return "/"
	case "/":
		return "/"
	}
	panic("Operator not found")
}

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

func reverseHumnEquation(instrName string, instrs map[string][]string) {
	fmt.Println(instrs["humn"])
	fmt.Println(instrs[instrName])
	if instrs[instrName][0] == "humn" {
		instrs[instrName][1] = inverseOpeLeftX(instrs[instrName][1])
		swapInstrs("humn", instrName, instrs)
		instrs["humn"][0] = instrName
	} else if instrs[instrName][2] == "humn" {
		instrs[instrName][1] = inverseOpeRightX(instrs[instrName][1])
		swapInstrs("humn", instrName, instrs)
		instrs["humn"][2] = instrName

	} else {
		panic(fmt.Sprint("humn not found in", instrName, instrs[instrName]))
	}
	fmt.Println(instrs["humn"])
	fmt.Println(instrs[instrName])
}

func reverseEquation(records []string, instrs map[string][]string) {
	fmt.Println(instrs)

	for i := 1; i < len(records); i++ {
		fmt.Println("reversing", records[i-1], records[i])
		reverseHumnEquation(records[i-1], instrs)
		if instrs[records[i]][0] == records[i-1] {
			instrs[records[i]][0] = "humn"
		} else {
			instrs[records[i]][2] = "humn"
		}
	}
}

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
			numbers[instrParts[0]] = tools.Atoi(tokens[0])
		} else {
			instrs[instrParts[0]] = tokens
			if tokens[0] == "humn" {
				fmt.Println(line)
			}
		}
	}

	fmt.Println(instrs)
	delete(numbers, "humn")
	calcTokens(instrs["root"][2], numbers, instrs)

	records := &[]string{}
	fmt.Println(hasHumnInstr("root", numbers, instrs, records))
	fmt.Println(records)
	reverseEquation(*records, instrs)
	fmt.Println(instrs)

	numbers[(*records)[0]] = numbers[instrs["root"][2]]

	fmt.Println(numbers)
	calcTokens("root", numbers, instrs)
	fmt.Println(numbers)

	fmt.Println(numbers["humn"])
	fmt.Println(numbers[instrs["root"][0]])
	fmt.Println(numbers[instrs["root"][2]])
}

func Main(path string) {
	s := tools.Readfile(path)

	lines := strings.Split(s, "\n")
	lines = lines[:len(lines)-1]
	partOne(lines)
	partTwo(lines)
}
