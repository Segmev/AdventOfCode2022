package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Segmev/AdventOfCode2022/tools"
)

type Expression struct {
	Operation string
	Operand   string
}

type monkey struct {
	NextTrueMonkey, NextFalseMonkey int
	TestNumber                      int
	Holding                         []int
	Expr                            Expression
	InspectCount                    int
}

func (m *monkey) InspectsAll(monkeys map[int]*monkey, stayWorried bool, modulo int) {
	for _, hObj := range m.Holding {
		m.InspectCount += 1
		operand := hObj % modulo
		res := 0
		if m.Expr.Operand != "old" {
			operand, _ = strconv.Atoi(m.Expr.Operand)
		}
		switch m.Expr.Operation {
		case "+":
			res = hObj + operand
		case "*":
			res = (hObj * operand) % modulo
		}
		if !stayWorried {
			res /= 3
		}
		if res%m.TestNumber == 0 {
			monkeys[m.NextTrueMonkey].Holding = append(monkeys[m.NextTrueMonkey].Holding, res)
		} else {
			monkeys[m.NextFalseMonkey].Holding = append(monkeys[m.NextFalseMonkey].Holding, res)
		}
	}
	m.Holding = []int{}
}

func parseHoldingItems(line string) (res []int) {
	line = strings.Split(line, ": ")[1]
	for _, nbS := range strings.Split(line, ", ") {
		nb, _ := strconv.Atoi(nbS)
		res = append(res, nb)
	}

	return res
}

func parseOperation(line string) Expression {
	operation := strings.Split(line, "= old ")[1]
	tokens := strings.Split(operation, " ")
	return Expression{Operation: tokens[0], Operand: tokens[1]}
}

func parseMonkeyTarget(line string) int {
	nb, _ := strconv.Atoi(strings.Split(line, " monkey ")[1])
	return nb
}

func parseTestNb(line string) int {
	nb, _ := strconv.Atoi(strings.Split(line, " by ")[1])
	return nb
}

func partOne(monkeysString []string) {
	modulo := 1
	monkeysMap := map[int]*monkey{}
	for i, monkeyS := range monkeysString {
		monkeysMap[i] = &monkey{}
		lines := strings.Split(monkeyS, "\n")[1:6]
		monkeysMap[i].Holding = parseHoldingItems(lines[0])
		monkeysMap[i].Expr = parseOperation(lines[1])
		monkeysMap[i].TestNumber = parseTestNb(lines[2])
		modulo *= monkeysMap[i].TestNumber
		monkeysMap[i].NextTrueMonkey = parseMonkeyTarget(lines[3])
		monkeysMap[i].NextFalseMonkey = parseMonkeyTarget(lines[4])
	}
	for round := 0; round < 20; round++ {
		for i := 0; i < len(monkeysMap); i++ {
			monkeysMap[i].InspectsAll(monkeysMap, false, modulo)
		}
	}
	counts := []int{}
	for _, monkey := range monkeysMap {
		counts = append(counts, monkey.InspectCount)
	}
	sort.Ints(counts)
	fmt.Println(counts[len(counts)-1] * counts[len(counts)-2])
}

func partTwo(monkeysString []string) {
	modulo := 1
	monkeysMap := map[int]*monkey{}
	for i, monkeyS := range monkeysString {
		monkeysMap[i] = &monkey{}
		lines := strings.Split(monkeyS, "\n")[1:6]
		monkeysMap[i].Holding = parseHoldingItems(lines[0])
		monkeysMap[i].Expr = parseOperation(lines[1])
		monkeysMap[i].TestNumber = parseTestNb(lines[2])
		modulo *= monkeysMap[i].TestNumber
		monkeysMap[i].NextTrueMonkey = parseMonkeyTarget(lines[3])
		monkeysMap[i].NextFalseMonkey = parseMonkeyTarget(lines[4])
	}
	for round := 0; round < 10_000; round++ {
		for i := 0; i < len(monkeysMap); i++ {
			monkeysMap[i].InspectsAll(monkeysMap, true, modulo)
		}
	}
	counts := []int{}
	for _, monkey := range monkeysMap {
		counts = append(counts, monkey.InspectCount)
	}
	sort.Ints(counts)
	fmt.Println(counts[len(counts)-1] * counts[len(counts)-2])
}

func main() {
	s := tools.Readfile("./input.txt")

	monkeys := strings.Split(s, "\n\n")
	partOne(monkeys)
	partTwo(monkeys)
}
