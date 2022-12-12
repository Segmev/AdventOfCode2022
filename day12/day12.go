package main

import (
	"fmt"
	"strings"

	"github.com/Segmev/AdventOfCode2022/tools"
)

type Coor struct {
	x, y  int
	steps int
}

func (c Coor) GetId() string {
	return fmt.Sprintf("%d_%d", c.x, c.y)
}

func findAllStart(lines []string) (As []Coor) {
	for i, line := range lines {
		for j, letter := range line {
			if letter == 'a' {
				As = append(As, Coor{x: j, y: i})
			}
		}
	}
	return As
}

func findSandE(lines []string) (S, E Coor) {
	for i, line := range lines {
		for j, letter := range line {
			if letter == 'S' || letter == '`' {
				S.x = j
				S.y = i
				lines[i] = strings.Replace(lines[i], "S", "`", 1)
			}
			if letter == 'E' || letter == '{' {
				E.x = j
				E.y = i
				lines[i] = strings.Replace(lines[i], "E", "{", 1)
			}
		}
	}
	return S, E
}

func getAccessibleEdges(lines []string, node Coor) (edges []Coor) {
	value := lines[node.y][node.x]
	if node.x-1 >= 0 && (int(value)-int(lines[node.y][node.x-1]) >= -1) {
		edges = append(edges, Coor{x: node.x - 1, y: node.y, steps: node.steps + 1})
	}
	if node.y-1 >= 0 && (int(value)-int(lines[node.y-1][node.x]) >= -1) {
		edges = append(edges, Coor{x: node.x, y: node.y - 1, steps: node.steps + 1})
	}
	if node.x+1 < len(lines[0]) && (int(value)-int(lines[node.y][node.x+1]) >= -1) {
		edges = append(edges, Coor{x: node.x + 1, y: node.y, steps: node.steps + 1})
	}
	if node.y+1 < len(lines) && (int(value)-int(lines[node.y+1][node.x]) >= -1) {
		edges = append(edges, Coor{x: node.x, y: node.y + 1, steps: node.steps + 1})
	}
	return edges
}

func getStepsFromStoE(lines []string, S, E Coor) int {
	visited := tools.GetSet[string]()
	qu := []Coor{S}

	for len(qu) >= 1 {
		node := qu[0]
		qu = qu[1:]
		visited.Add(node.GetId())
		if node.GetId() == E.GetId() {
			return node.steps - 2
		} else {
			edges := getAccessibleEdges(lines, node)
			for _, edge := range edges {
				if !visited.Contains(edge.GetId()) {
					qu = append(qu, edge)
					visited.Add(edge.GetId())
				}
			}
		}
	}
	return 50000
}

func partOne(lines []string) {
	S, E := findSandE(lines)
	fmt.Println(getStepsFromStoE(lines, S, E))
}

func partTwo(lines []string) {
	S, E := findSandE(lines)
	min := 50000
	starts := append(findAllStart(lines), S)
	for _, s := range starts {
		if steps := getStepsFromStoE(lines, s, E); min > steps {
			min = steps
		}
	}
	fmt.Println(min)
}

func main() {
	s := tools.Readfile("./input.txt")

	lines := strings.Split(s, "\n")
	partOne(lines[:len(lines)-1])
	partTwo(lines[:len(lines)-1])
}
