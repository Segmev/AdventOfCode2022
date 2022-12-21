package day18

import (
	"fmt"
	"math"
	"strings"

	"github.com/Segmev/AdventOfCode2022/tools"
)

type Cube struct {
	x, y, z     int
	LinkedCubes map[string]*Cube
	waterFace   int
	id          string
}

func (c *Cube) GetId() string {
	if c.id == "" {
		c.id = fmt.Sprintf("%d_%d_%d", c.x, c.y, c.z)
	}
	return c.id
}

func partOne(lines []string) {
	cubes := map[string]*Cube{}
	for _, line := range lines {
		coors := strings.Split(line, ",")
		cube := Cube{x: tools.Atoi(coors[0]), y: tools.Atoi(coors[1]), z: tools.Atoi(coors[2]), LinkedCubes: map[string]*Cube{}}
		for i := -1; i <= 1; i++ {
			if i == 0 {
				continue
			}
			if c, ok := cubes[fmt.Sprintf("%d_%d_%d", cube.x+i, cube.y, cube.z)]; ok {
				cube.LinkedCubes[c.GetId()] = cubes[fmt.Sprintf("%d_%d_%d", cube.x+i, cube.y, cube.z)]
				cubes[fmt.Sprintf("%d_%d_%d", cube.x+i, cube.y, cube.z)].LinkedCubes[cube.GetId()] = &cube
			}
			if c, ok := cubes[fmt.Sprintf("%d_%d_%d", cube.x, cube.y+i, cube.z)]; ok {
				cube.LinkedCubes[c.GetId()] = cubes[fmt.Sprintf("%d_%d_%d", cube.x, cube.y+i, cube.z)]
				cubes[fmt.Sprintf("%d_%d_%d", cube.x, cube.y+i, cube.z)].LinkedCubes[cube.GetId()] = &cube
			}
			if c, ok := cubes[fmt.Sprintf("%d_%d_%d", cube.x, cube.y, cube.z+i)]; ok {
				cube.LinkedCubes[c.GetId()] = cubes[fmt.Sprintf("%d_%d_%d", cube.x, cube.y, cube.z+i)]
				cubes[fmt.Sprintf("%d_%d_%d", cube.x, cube.y, cube.z+i)].LinkedCubes[cube.GetId()] = &cube
			}
		}

		cubes[cube.GetId()] = &cube
	}
	total := 0
	for _, cube := range cubes {
		total += 6 - len(cube.LinkedCubes)
	}
	fmt.Println(total)
}

func (c Cube) getNeighboors() []*Cube {
	return []*Cube{
		{x: c.x - 1, y: c.y, z: c.z},
		{x: c.x + 1, y: c.y, z: c.z},
		{x: c.x, y: c.y - 1, z: c.z},
		{x: c.x, y: c.y + 1, z: c.z},
		{x: c.x, y: c.y, z: c.z - 1},
		{x: c.x, y: c.y, z: c.z + 1},
	}
}

func inZone(minX, minY, minZ, maxX, maxY, maxZ int, c Cube) bool {
	return minX <= c.x && c.x <= maxX && minY <= c.y && c.y <= maxY && minZ <= c.z && c.z <= maxZ
}

func countOutside(lavaCubes map[string]*Cube, airCubes tools.Set[string]) int {
	count := 0
	for _, cube := range lavaCubes {
		for _, nCube := range cube.getNeighboors() {
			if airCubes.Contains(nCube.GetId()) {
				count += 1
			}
		}
	}
	return count
}

func exploreZone(minX, minY, minZ, maxX, maxY, maxZ int, lavaCubes tools.Set[string]) tools.Set[string] {
	visitedAirCubes := tools.GetSet[string]()
	toVisitAirCubes := []*Cube{
		{x: minX, y: minY, z: minZ},
	}

	visitCount := 0
	for len(toVisitAirCubes) > 0 {
		visitCount += 1
		inspectedCube := toVisitAirCubes[0]
		toVisitAirCubes = toVisitAirCubes[1:]

		if inZone(minX, minY, minZ, maxX, maxY, maxZ, *inspectedCube) &&
			!lavaCubes.Contains(inspectedCube.GetId()) &&
			!visitedAirCubes.Contains(inspectedCube.GetId()) {
			visitedAirCubes.Add(inspectedCube.GetId())
			toVisitAirCubes = append(toVisitAirCubes, inspectedCube.getNeighboors()...)
		}
	}
	return visitedAirCubes
}

func partTwo(lines []string) {
	cubes := tools.GetSet[string]()
	lavaCubes := map[string]*Cube{}
	minX, maxX := math.MaxInt, math.MinInt
	minY, maxY := math.MaxInt, math.MinInt
	minZ, maxZ := math.MaxInt, math.MinInt
	for _, line := range lines {
		coors := strings.Split(line, ",")
		cubeX, cubeY, cubeZ := tools.Atoi(coors[0]), tools.Atoi(coors[1]), tools.Atoi(coors[2])
		cube := Cube{x: cubeX, y: cubeY, z: cubeZ, LinkedCubes: map[string]*Cube{}}

		if cubeX < minX {
			minX = cubeX
		}
		if cubeX > maxX {
			maxX = cubeX
		}
		if cubeY < minY {
			minY = cubeY
		}
		if cubeY > maxY {
			maxY = cubeY
		}
		if cubeZ < minZ {
			minZ = cubeZ
		}
		if cubeZ > maxZ {
			maxZ = cubeZ
		}
		cubes.Add(cube.GetId())
		lavaCubes[cube.GetId()] = &cube
	}
	airCubes := exploreZone(minX-1, minY-1, minZ-1, maxX+1, maxY+1, maxZ+1, cubes)
	fmt.Println(countOutside(lavaCubes, airCubes))
}

func Main(path string) {
	s := tools.Readfile(path)

	lines := strings.Split(s, "\n")
	partOne(lines[:len(lines)-1])
	partTwo(lines[:len(lines)-1])
}
