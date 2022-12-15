package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/Segmev/AdventOfCode2022/tools"
)

const (
	minBound = 0
	maxBound = 4_000_000
	Yrow     = 10
)

type Coor struct {
	x, y  int
	bDist int
}

func manhDist(a, b Coor) int {
	return tools.Abs(a.x-b.x) + tools.Abs(a.y-b.y)
}

func parseCoors(info string) Coor {
	oInfo := strings.Split(strings.Split(info, "at ")[1], ", ")
	x, _ := strconv.Atoi(strings.Split(oInfo[0], "=")[1])
	y, _ := strconv.Atoi(strings.Split(oInfo[1], "=")[1])

	return Coor{x: x, y: y}
}

func canBeaconExistsAt(x, y int, sensors []*Coor) bool {
	for _, sensor := range sensors {
		if manhDist(Coor{x: x, y: y}, *sensor) <= sensor.bDist {
			return false
		}
	}
	return true
}

func parseInput(lines []string) ([]*Coor, tools.Set[string], int, int) {
	lowestX, highestX := math.MaxInt, math.MinInt
	sensors := []*Coor{}
	beacons := tools.Set[string]{}
	for _, line := range lines {
		lineParts := strings.Split(line, ": ")
		B := parseCoors(lineParts[1])
		beacons.Add(fmt.Sprintf("%d,%d", B.x, B.y))
		if B.x < lowestX {
			lowestX = B.x
		}
		if B.x > highestX {
			highestX = B.x
		}
		S := parseCoors(lineParts[0])
		S.bDist = manhDist(S, B)
		if S.x-S.bDist < lowestX {
			lowestX = S.x - S.bDist
		}
		if S.x+S.bDist > highestX {
			highestX = S.x + S.bDist
		}
		sensors = append(sensors, &S)
	}
	sort.Slice(sensors, func(i, j int) bool {
		return sensors[i].bDist > sensors[j].bDist
	})
	return sensors, beacons, lowestX, highestX
}

func partOne(sensors []*Coor, beacons tools.Set[string], lowestX, highestX int) {
	count := 0
	for i := lowestX; i <= highestX; i++ {
		if res := canBeaconExistsAt(i, Yrow, sensors); !res && !beacons.Contains(fmt.Sprintf("%d,%d", i, Yrow)) {
			count++
		}
	}
	fmt.Println(count)
}

func partTwo(sensors []*Coor) Coor {
	for _, sensor := range sensors {
		yStep := -1
		for x := sensor.x - sensor.bDist - 1; x <= sensor.x; x++ {
			yStep += 1
			if x < minBound || x > maxBound {
				continue
			}
			if sensor.y-yStep >= minBound && canBeaconExistsAt(x, sensor.y-yStep, sensors) {
				return Coor{x: x, y: sensor.y - yStep}
			}
			if sensor.y+yStep <= maxBound && canBeaconExistsAt(x, sensor.y+yStep, sensors) {
				return Coor{x: x, y: sensor.y + yStep}
			}
		}

		yStep = sensor.y + 2
		for x := sensor.x + sensor.bDist + 1; x >= sensor.x; x-- {
			yStep -= 1
			if x < minBound || x > maxBound {
				continue
			}
			if sensor.y-yStep >= minBound && canBeaconExistsAt(x, sensor.y-yStep, sensors) {
				return Coor{x: x, y: sensor.y - yStep}
			}
			if sensor.y+yStep <= maxBound && canBeaconExistsAt(x, sensor.y+yStep, sensors) {
				return Coor{x: x, y: sensor.y + yStep}
			}
		}
	}
	return Coor{}
}

func main() {
	s := tools.Readfile("./input.txt")

	lines := strings.Split(s, "\n")
	sensors, beacons, lowestX, highestX := parseInput(lines[:len(lines)-1])
	partOne(sensors, beacons, lowestX, highestX)
	res := partTwo(sensors)
	fmt.Println(maxBound*res.x+res.y, "at", res.x, res.y)
}
