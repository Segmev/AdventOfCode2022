package day22

import (
	"fmt"
	"math"
	"strings"

	"github.com/Segmev/AdventOfCode2022/tools"
)

type MonkeyMap struct {
	walls                                        tools.Set[string]
	rowShifts, rowLens                           []int
	posX, posY, dir, prevPosX, prevPosY, prevDir int
	directions                                   [][2]int
	warps                                        []Warp
	cubesize                                     int
}

type Warp struct {
	id                     string
	rotation               int
	symX, symY             bool
	minX, maxX, minY, maxY int
	destX, destY           int
}

func (w Warp) IsWarping(x, y int) bool {
	return w.minX <= x && x <= w.maxX && w.minY <= y && y <= w.maxY
}

func (w Warp) Teleport(mm *MonkeyMap) bool {
	size := mm.cubesize
	x, y := mm.posX%size, mm.posY%size
	if w.symX {
		x = size - x - 1
	}
	if w.symY {
		y = size - y - 1
	}
	if tools.Abs(w.rotation) == 1 {
		x, y = y, x
	}
	x, y = x+w.destX, y+w.destY
	if mm.walls.Contains(fmt.Sprintf("%d_%d", x, y)) {
		return false
	}
	mm.posX, mm.posY = x, y
	if w.rotation == 1 {
		mm.TurnRight()
	} else if w.rotation == -1 {
		mm.TurnLeft()
	} else if w.rotation == 2 {
		mm.TurnRight()
		mm.TurnRight()
	}
	return true
}

func (mm *MonkeyMap) TurnRight() {
	mm.prevDir = mm.dir
	mm.dir = (1 + mm.dir) % 4
}

func (mm *MonkeyMap) TurnLeft() {
	mm.prevDir = mm.dir
	mm.dir -= 1
	if mm.dir < 0 {
		mm.dir = 3
	}
}

func (mm MonkeyMap) IsWall(x, y int) bool {
	return mm.walls.Contains(fmt.Sprintf("%d_%d", x, y))
}

func (mm *MonkeyMap) Walk(steps int) {
	mm.prevPosX = mm.posX
	mm.prevPosY = mm.posY
	for s := 1; s <= steps; s++ {
		if mm.dir%2 == 0 {
			futurX := tools.Modulo(mm.posX-mm.rowShifts[mm.posY]+mm.directions[mm.dir][0], mm.rowLens[mm.posY]) + mm.rowShifts[mm.posY]
			if mm.IsWall(futurX, mm.posY) {
				break
			}
			mm.posX = futurX

		} else {
			var futurRowY int
			for r := 1; ; r++ {
				futurRowY = tools.Modulo(mm.posY+(r*mm.directions[mm.dir][1]), len(mm.rowShifts))
				if mm.rowShifts[futurRowY] <= mm.posX && mm.posX < mm.rowShifts[futurRowY]+mm.rowLens[futurRowY] {
					break
				}
			}
			if mm.IsWall(mm.posX, futurRowY) {
				break
			}
			mm.posY = futurRowY
		}
	}
}

func (mm *MonkeyMap) GetWarp(x, y int) *Warp {
	for i := range mm.warps {
		if mm.warps[i].IsWarping(x, y) {
			return &mm.warps[i]
		}
	}
	return nil
}

func (mm *MonkeyMap) WalkCube(steps int) {
	mm.prevPosX = mm.posX
	mm.prevPosY = mm.posY
	for s := 1; s <= steps; s++ {
		if warp := mm.GetWarp(mm.posX+mm.directions[mm.dir][0], mm.posY+mm.directions[mm.dir][1]); warp != nil {
			if !(warp.Teleport(mm)) {
				break
			}
		} else {
			if mm.dir%2 == 0 {
				futurX := tools.Modulo(mm.posX-mm.rowShifts[mm.posY]+mm.directions[mm.dir][0], mm.rowLens[mm.posY]) + mm.rowShifts[mm.posY]
				if mm.IsWall(futurX, mm.posY) {
					break
				}
				mm.posX = futurX
			} else {
				var futurRowY int
				for r := 1; ; r++ {
					futurRowY = tools.Modulo(mm.posY+(r*mm.directions[mm.dir][1]), len(mm.rowShifts))
					if mm.rowShifts[futurRowY] <= mm.posX && mm.posX < mm.rowShifts[futurRowY]+mm.rowLens[futurRowY] {
						break
					}
				}
				if mm.IsWall(mm.posX, futurRowY) {
					break
				}
				mm.posY = futurRowY
			}
		}
	}
}

func (mm MonkeyMap) printMap() {
	getCursor := func(dir int) rune {
		switch dir {
		case 0:
			return '>'
		case 1:
			return 'V'
		case 2:
			return '<'
		case 3:
			return 'A'
		}
		return '+'
	}

	for y := 0; y < len(mm.rowShifts); y++ {
		for x := 0; x < mm.rowShifts[y]; x++ {
			fmt.Print(" ")
		}
		for ix := 0; ix < mm.rowLens[y]; ix++ {
			x := ix + mm.rowShifts[y]
			if mm.walls.Contains(fmt.Sprintf("%d_%d", x, y)) {
				fmt.Print("#")
			} else if x == mm.posX && y == mm.posY {
				fmt.Printf("\033[1;32m%s\033[0m", string(getCursor(mm.dir)))
			} else if x == mm.prevPosX && y == mm.prevPosY {
				fmt.Printf("\033[1;31m%s\033[0m", string(getCursor(mm.prevDir)))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (mm *MonkeyMap) initMap(lines []string) {
	cubesize := math.MaxInt
	for y, line := range lines {
		x := 0
		for x < len(line) {
			if line[x] != ' ' {
				break
			}
			x++
		}
		rowShift := x
		if y == 0 {
			mm.posX = rowShift
		}
		mm.rowShifts = append(mm.rowShifts, rowShift)
		for x < len(line) {
			if line[x] == '#' {
				mm.walls.Add(fmt.Sprintf("%d_%d", x, y))
			}
			x++
		}
		if (x - rowShift) < cubesize {
			cubesize = x - rowShift
		}
		mm.rowLens = append(mm.rowLens, x-rowShift)
	}
	mm.cubesize = cubesize
}

func partOne(lines []string) {
	mm := MonkeyMap{walls: tools.GetSet[string](), rowShifts: []int{}, rowLens: []int{}, dir: 0, directions: [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}}
	mm.initMap(lines[:len(lines)-2])

	instrs := lines[len(lines)-1]
	leftI, rightI := 0, 0
	for leftI < len(instrs) {
		for rightI < len(instrs) && '0' <= instrs[rightI] && instrs[rightI] <= '9' {
			rightI++
		}
		mm.Walk(tools.Atoi(instrs[leftI:rightI]))
		if rightI >= len(instrs) {
			break
		}
		if instrs[rightI] == 'R' {
			mm.TurnRight()
		} else if instrs[rightI] == 'L' {
			mm.TurnLeft()
		} else {
			panic(fmt.Sprintln("Unexpected", instrs[rightI]))
		}
		rightI += 1
		leftI = rightI
	}
	fmt.Println((mm.posX+1)*4 + (mm.posY+1)*1000 + mm.dir)
}

func (mm *MonkeyMap) ScanWarps() {
	s := tools.Readfile("./day22/warps.txt")

	lines := strings.Split(s, "\n")
	lines = lines[:len(lines)-1]
	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		lsplitted := strings.Split(line, ":")
		tokens := strings.Split(lsplitted[1], ",")
		mm.warps = append(mm.warps, Warp{
			id:       lsplitted[0],
			minX:     tools.Atoi(tokens[0]),
			maxX:     tools.Atoi(tokens[1]),
			minY:     tools.Atoi(tokens[2]),
			maxY:     tools.Atoi(tokens[3]),
			destX:    tools.Atoi(tokens[4]),
			destY:    tools.Atoi(tokens[5]),
			rotation: tools.Atoi(tokens[6]),
			symX:     tokens[7] == "T",
			symY:     tokens[8] == "T",
		})
	}
}

func partTwo(lines []string) {
	mm := MonkeyMap{walls: tools.GetSet[string](), rowShifts: []int{}, rowLens: []int{}, dir: 0, directions: [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}}
	mm.initMap(lines[:len(lines)-2])
	mm.ScanWarps()

	instrs := lines[len(lines)-1]
	leftI, rightI := 0, 0
	for leftI < len(instrs) {
		for rightI < len(instrs) && '0' <= instrs[rightI] && instrs[rightI] <= '9' {
			rightI++
		}

		mm.WalkCube(tools.Atoi(instrs[leftI:rightI]))
		if rightI >= len(instrs) {
			break
		}
		if instrs[rightI] == 'R' {
			mm.TurnRight()
		} else if instrs[rightI] == 'L' {
			mm.TurnLeft()
		} else {
			panic(fmt.Sprintln("Unexpected", instrs[rightI]))
		}

		rightI += 1
		leftI = rightI
	}
	fmt.Println((mm.posX+1)*4 + (mm.posY+1)*1000 + mm.dir)
}

func Main(path string) {
	s := tools.Readfile(path)

	lines := strings.Split(s, "\n")
	lines = lines[:len(lines)-1]
	partOne(lines)
	partTwo(lines)
}
