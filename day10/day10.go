package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Segmev/AdventOfCode2022/tools"
)

type CPU struct {
	cycle       int
	reg         int
	strenghtSum int
	pixel       int
}

func (cpu *CPU) addCycle() {
	cpu.cycle += 1
	diff := tools.Abs(cpu.cycle%40 - (cpu.reg + 1))
	if cpu.cycle%40 == 0 {
		fmt.Println()
	} else {
		if diff <= 1 {
			fmt.Print("# ")
		} else {
			fmt.Print(". ")
		}
	}

	if (cpu.cycle+20)%40 == 0 {
		cpu.strenghtSum += (cpu.cycle * cpu.reg)
	}
}

func (cpu *CPU) Noop() {
	cpu.addCycle()
}

func (cpu *CPU) Addx(x int) {
	cpu.addCycle()
	cpu.addCycle()
	cpu.reg += x
}

func partOne(lines []string) {
	cpu := CPU{reg: 1}
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		if tokens[0] == "noop" {
			cpu.Noop()
		} else if tokens[0] == "addx" {
			x, _ := strconv.Atoi(tokens[1])
			cpu.Addx(x)
		}
	}
	fmt.Println("partOne", cpu.strenghtSum)
}

func partTwo(lines []string) {
}

func main() {
	s := tools.Readfile("./input.txt")

	lines := strings.Split(s, "\n")
	partOne(lines)
	partTwo(lines)
}
