package day14

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Segmev/AdventOfCode2022/tools"
)

type Coor struct {
	id   string
	x, y int
}

func (c *Coor) setX(x int) {
	c.x = x
	c.id = fmt.Sprintf("%d,%d", c.x, c.y)
}

func (c *Coor) setY(y int) {
	c.y = y
	c.id = fmt.Sprintf("%d,%d", c.x, c.y)
}

func GetCoorFromInts(x, y int) Coor {
	return Coor{fmt.Sprintf("%d,%d", x, y), x, y}
}

func GetCoorFromStrings(coordinates string) Coor {
	coords := strings.Split(coordinates, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])

	return Coor{coordinates, x, y}
}

func GetRockMap(lines []string) (map[string]*Coor, int) {
	taken := map[string]*Coor{}
	var prevCoords *Coor
	deepestY := -1

	for _, line := range lines[:len(lines)-1] {
		for i, coordinates := range strings.Split(line, " -> ") {
			coords := GetCoorFromStrings(coordinates)
			if deepestY < coords.y {
				deepestY = coords.y
			}
			taken[coords.id] = &coords
			if i > 0 {
				if coords.x == prevCoords.x {
					lowestBound, higherBound := 0, 0
					if coords.y < prevCoords.y {
						lowestBound, higherBound = coords.y, prevCoords.y
					} else {
						lowestBound, higherBound = prevCoords.y, coords.y
					}
					for i := lowestBound; i <= higherBound; i++ {
						intercoords := GetCoorFromInts(coords.x, i)
						taken[intercoords.id] = &intercoords
					}
				} else {
					lowestBound, higherBound := 0, 0
					if coords.x < prevCoords.x {
						lowestBound, higherBound = coords.x, prevCoords.x
					} else {
						lowestBound, higherBound = prevCoords.x, coords.x
					}
					for i := lowestBound; i <= higherBound; i++ {
						intercoords := GetCoorFromInts(i, coords.y)
						taken[intercoords.id] = &intercoords
					}
				}
			}
			prevCoords = &coords
		}
	}
	return taken, deepestY
}

func PoorSand(rockMap map[string]*Coor, lowestY int) (count int) {
	for {
		sandUnit := GetCoorFromInts(500, 0)
		for {
			if sandUnit.y > lowestY {
				return count
			}
			_, ok := rockMap[fmt.Sprintf("%d,%d", sandUnit.x, sandUnit.y+1)]
			if !ok {
				sandUnit.setY(sandUnit.y + 1)
				continue
			} else {
				_, ok = rockMap[fmt.Sprintf("%d,%d", sandUnit.x-1, sandUnit.y+1)]
				if !ok {
					sandUnit.setX(sandUnit.x - 1)
					sandUnit.setY(sandUnit.y + 1)
					continue
				} else {
					_, ok = rockMap[fmt.Sprintf("%d,%d", sandUnit.x+1, sandUnit.y+1)]
					if !ok {
						sandUnit.setX(sandUnit.x + 1)
						sandUnit.setY(sandUnit.y + 1)
						continue
					} else {
						rockMap[sandUnit.id] = &sandUnit
						break
					}
				}
			}
		}

		count += 1
	}
}

func partOne(lines []string) {
	taken, deepestY := GetRockMap(lines)
	fmt.Println(PoorSand(taken, deepestY))
}

func PoorSandInf(rockMap map[string]*Coor, lowestY int) (count int) {
	groundY := lowestY + 2
	for {
		if _, ok := rockMap["500,0"]; ok {
			return count
		}
		sandUnit := GetCoorFromInts(500, 0)

		for {
			_, ok := rockMap[fmt.Sprintf("%d,%d", sandUnit.x, sandUnit.y+1)]
			if !ok && sandUnit.y+1 < groundY {
				sandUnit.setY(sandUnit.y + 1)
				continue
			} else {
				_, ok = rockMap[fmt.Sprintf("%d,%d", sandUnit.x-1, sandUnit.y+1)]
				if !ok && sandUnit.y+1 < groundY {
					sandUnit.setX(sandUnit.x - 1)
					sandUnit.setY(sandUnit.y + 1)
					continue
				} else {
					_, ok = rockMap[fmt.Sprintf("%d,%d", sandUnit.x+1, sandUnit.y+1)]
					if !ok && sandUnit.y+1 < groundY {
						sandUnit.setX(sandUnit.x + 1)
						sandUnit.setY(sandUnit.y + 1)
						continue
					} else {
						rockMap[sandUnit.id] = &sandUnit
						break
					}
				}
			}
		}

		count += 1
	}
}

func partTwo(lines []string) {
	taken, deepestY := GetRockMap(lines)
	fmt.Println(PoorSandInf(taken, deepestY))
}

func Main(path string) {
	s := tools.Readfile(path)

	lines := strings.Split(s, "\n")
	partOne(lines)
	partTwo(lines)
}
