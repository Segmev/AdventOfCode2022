package day09

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Segmev/AdventOfCode2022/tools"
)

type Vector struct {
	x, y   int
	px, py int

	next *Vector
}

func (node *Vector) AddCoors(x, y int) {
	node.px = node.x
	node.py = node.y

	node.x += x
	node.y += y
}

func initRope(nodeAmount int) *Vector {
	H := Vector{next: nil}
	B := &H
	for i := 0; i < nodeAmount; i++ {
		B.next = &Vector{next: nil}
		B = B.next
	}
	return &H
}

func updateNextNode(H, N *Vector) {
	xDiff := H.x - N.x
	yDiff := H.y - N.y
	if tools.Abs(xDiff) <= 1 && tools.Abs(yDiff) <= 1 {
		return
	}
	N.px = N.x
	N.py = N.y

	if xDiff == 0 {
		N.y = H.py
	} else if yDiff == 0 {
		N.x = H.px
	} else {
		if xDiff > 0 {
			N.x += 1
		} else {
			N.x -= 1
		}

		if yDiff > 0 {
			N.y += 1
		} else {
			N.y -= 1
		}
	}
}

func updateToTail(H *Vector, visited *tools.Set[string]) {
	B := H

	for B.next != nil {
		updateNextNode(B, B.next)
		B = B.next
	}
	visited.Add(fmt.Sprintf("%d_%d", B.x, B.y))
}

func printMap(H Vector, boundX, boundY int) {
	coors := tools.Set[string]{}
	for B := &H; B != nil; B = B.next {
		coors.Add(fmt.Sprintf("%d_%d", B.x, B.y))
	}
	for i := boundX; i >= -boundX; i-- {
		for j := -boundY; j < boundY; j++ {
			if coors.Contains(fmt.Sprintf("%d_%d", j, i)) {
				fmt.Print("X")
			} else if i == 0 && j == 0 {
				fmt.Print("S")
			} else {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func partOne(lines []string) {
	actionSteps := map[string]Vector{
		"L": {x: -1, y: 0},
		"R": {x: 1, y: 0},
		"D": {x: 0, y: -1},
		"U": {x: 0, y: 1},
	}
	visited := tools.Set[string]{}

	H := initRope(1)

	for _, line := range lines[:len(lines)-1] {
		instr := strings.Split(line, " ")
		stepsNumber, _ := strconv.Atoi(instr[1])

		for i := 0; i < stepsNumber; i++ {
			H.AddCoors(actionSteps[instr[0]].x, actionSteps[instr[0]].y)
			updateToTail(H, &visited)
		}
	}

	fmt.Println(len(visited))
}

func partTwo(lines []string) {
	actionSteps := map[string]Vector{
		"L": {x: -1, y: 0},
		"R": {x: 1, y: 0},
		"D": {x: 0, y: -1},
		"U": {x: 0, y: 1},
	}
	visited := tools.Set[string]{}
	H := initRope(9)

	for _, line := range lines[:len(lines)-1] {
		instr := strings.Split(line, " ")
		stepsNumber, _ := strconv.Atoi(instr[1])

		for i := 0; i < stepsNumber; i++ {
			H.AddCoors(actionSteps[instr[0]].x, actionSteps[instr[0]].y)
			updateToTail(H, &visited)
		}
	}
	fmt.Println(len(visited))
}

func Main(path string) {
	s := tools.Readfile(path)

	lines := strings.Split(s, "\n")
	partOne(lines)
	partTwo(lines)
}
