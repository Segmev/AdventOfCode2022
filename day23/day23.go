package day23

import (
	"fmt"
	"math"
	"strings"

	"github.com/Segmev/AdventOfCode2022/tools"
)

type Coordinates struct {
	X, Y         int
	NextX, NextY int
}

func (c *Coordinates) GetId() string {
	return fmt.Sprintf("%d_%d", c.X, c.Y)
}

func HasElfInLine(x, y int, elves map[string]*Coordinates) bool {
	for i := -1; i <= 1; i++ {
		if _, ok := elves[fmt.Sprintf("%d_%d", x+i, y)]; ok {
			return true
		}
	}
	return false
}

func HasElfInColumn(x, y int, elves map[string]*Coordinates) bool {
	for i := -1; i <= 1; i++ {
		if _, ok := elves[fmt.Sprintf("%d_%d", x, y+i)]; ok {
			return true
		}
	}
	return false
}

func (elf *Coordinates) GetNextMove(elves map[string]*Coordinates, directions []string) (int, int, bool) {
	for _, d := range directions {
		if d == "N" && !HasElfInLine(elf.X, elf.Y-1, elves) {
			// fmt.Println("N")
			return elf.X, elf.Y - 1, true
		}

		if d == "S" && !HasElfInLine(elf.X, elf.Y+1, elves) {
			// fmt.Println("S")
			return elf.X, elf.Y + 1, true
		}

		if d == "W" && !HasElfInColumn(elf.X-1, elf.Y, elves) {
			// fmt.Println("W")
			return elf.X - 1, elf.Y, true
		}

		if d == "E" && !HasElfInColumn(elf.X+1, elf.Y, elves) {
			// fmt.Println("E")
			return elf.X + 1, elf.Y, true
		}
	}

	return elf.X, elf.Y, false
}

func GetNextState(elves map[string]*Coordinates, directions []string) map[string]*Coordinates {
	nextState := map[string]*Coordinates{}
	tempNextState := map[string][]*Coordinates{}
	for _, elf := range elves {
		if len(elf.GetNeighboor(elves)) == 0 {
			nextState[elf.GetId()] = elf
			continue
		}
		if nextX, nextY, canMove := elf.GetNextMove(elves, directions); canMove {
			elf.NextX = nextX
			elf.NextY = nextY
			tempNextState[fmt.Sprintf("%d_%d", nextX, nextY)] = append(tempNextState[fmt.Sprintf("%d_%d", nextX, nextY)], elf)
		} else {
			nextState[elf.GetId()] = elf
		}
	}
	// fmt.Println(tempNextState)
	for _, potentialMove := range tempNextState {
		if len(potentialMove) == 1 {
			move := Coordinates{X: potentialMove[0].NextX, Y: potentialMove[0].NextY}
			nextState[move.GetId()] = &move
		} else {
			for _, elf := range potentialMove {
				nextState[elf.GetId()] = elf
			}
		}
	}

	return nextState
}

func (e Coordinates) GetNeighboor(elves map[string]*Coordinates) []*Coordinates {
	neighboors := []*Coordinates{}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i == 0 && j == 0) {
				elf, ok := elves[fmt.Sprintf("%d_%d", e.X+i, e.Y+j)]
				if ok {
					neighboors = append(neighboors, elf)
				}
			}
		}
	}
	return neighboors
}

func debugMap(elves map[string]*Coordinates) {
	fmt.Println(elves)
	fmt.Print("   ")
	for x := -5; x < 15; x++ {
		fmt.Printf("%.2d ", int(math.Abs(float64(x))))
	}
	fmt.Println()
	for y := -5; y < 15; y++ {
		fmt.Printf("%.2d ", int(math.Abs(float64(y))))
		for x := -5; x < 15; x++ {
			if _, ok := elves[fmt.Sprintf("%d_%d", x, y)]; ok {
				fmt.Print("#  ")
			} else {
				fmt.Print(".  ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func displayMap(elves map[string]*Coordinates) {
	for y := -5; y < 15; y++ {
		for x := -5; x < 15; x++ {
			if _, ok := elves[fmt.Sprintf("%d_%d", x, y)]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func EmptySpaceCount(elves map[string]*Coordinates) int {
	minX, maxX := math.MaxInt, math.MinInt
	minY, maxY := math.MaxInt, math.MinInt
	for _, elf := range elves {
		if minX > elf.X {
			minX = elf.X
		}
		if maxX < elf.X {
			maxX = elf.X
		}
		if minY > elf.Y {
			minY = elf.Y
		}
		if maxY < elf.Y {
			maxY = elf.Y
		}
	}
	// fmt.Println(minX, minY, maxX, maxY)
	return (1+maxX-minX)*(1+maxY-minY) - len(elves)
}

func partOne(lines []string) {
	elves := map[string]*Coordinates{}

	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				elf := Coordinates{X: x, Y: y}
				elves[elf.GetId()] = &elf
			}
		}
	}
	// displayMap(elves)
	directions := []string{"N", "S", "W", "E"}

	for i := 0; i < 10; i++ {
		// fmt.Println(i + 1)
		elves = GetNextState(elves, directions)
		// displayMap(elves)
		directions = append(directions[1:], directions[0])
	}
	fmt.Println(EmptySpaceCount(elves))
}

func AreStatesEquals(elves, futurElves map[string]*Coordinates) bool {
	for k := range elves {
		if _, ok := futurElves[k]; !ok {
			return false
		}
	}
	return true
}

func partTwo(lines []string) {
	elves := map[string]*Coordinates{}

	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				elf := Coordinates{X: x, Y: y}
				elves[elf.GetId()] = &elf
			}
		}
	}
	// displayMap(elves)
	directions := []string{"N", "S", "W", "E"}
	i := 1
	futurElves := map[string]*Coordinates{}
	for {
		futurElves = GetNextState(elves, directions)
		if AreStatesEquals(elves, futurElves) {
			fmt.Println(i)
			return
		}
		directions = append(directions[1:], directions[0])
		elves = futurElves
		i++
	}
}

func Main(path string) {
	s := tools.Readfile(path)

	lines := strings.Split(s, "\n")
	lines = lines[:len(lines)-1]
	partOne(lines)
	partTwo(lines)
}
