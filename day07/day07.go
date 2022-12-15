package day07

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Segmev/AdventOfCode2022/tools"
)

func constructFS(lines []string) *tools.GraphNode[int] {
	origin := tools.GraphNode[int]{Id: "/", Nodes: make(map[string]*tools.GraphNode[int])}
	currentNode := &origin

	for _, line := range lines[1:] {
		tokens := strings.Split(line, " ")
		if len(tokens) < 2 {
			continue
		}
		if tokens[0] == "$" {
			switch tokens[1] {
			case "ls":
				continue
			case "cd":
				if tokens[2] == ".." {
					currentNode = currentNode.Parent
				} else {
					folder := tokens[2] + "/"
					if !currentNode.Contains(folder) {
						currentNode.Nodes[folder] = &tools.GraphNode[int]{Id: folder, Parent: currentNode, Nodes: make(map[string]*tools.GraphNode[int])}
						currentNode = currentNode.Nodes[folder]
					}
				}
			}
		} else {
			if !(tokens[0] == "dir") {
				size, _ := strconv.Atoi(tokens[0])
				if !currentNode.Contains(tokens[1]) {
					currentNode.Nodes[tokens[1]] = &tools.GraphNode[int]{Id: tokens[1], Parent: currentNode, Nodes: nil, Value: size}
					currentNode.Value += size
					parentNode := currentNode.Parent
					for parentNode != nil {
						parentNode.Value += size
						parentNode = parentNode.Parent
					}
				}
			}
		}
	}
	return &origin
}

func getFoldersAtMost(size int, currentNode tools.GraphNode[int], selectedNodes *[]*tools.GraphNode[int]) {
	if currentNode.Id[len(currentNode.Id)-1:] != "/" {
		return
	}
	if currentNode.Value <= size {
		*selectedNodes = append(*selectedNodes, &currentNode)
	}
	for _, n := range currentNode.Nodes {
		getFoldersAtMost(size, *n, selectedNodes)
	}
}

func partOne(origin *tools.GraphNode[int]) {
	foundFolders := []*tools.GraphNode[int]{}
	getFoldersAtMost(100_000, *origin, &foundFolders)
	sum := 0
	for _, n := range foundFolders {
		sum += n.Value
	}
	fmt.Println(sum)
}

func getUpperClosest(size int, currentNode tools.GraphNode[int], selectedNode *tools.GraphNode[int]) *tools.GraphNode[int] {
	if currentNode.Value >= size && currentNode.Value < selectedNode.Value {
		selectedNode = &currentNode
	}
	for _, n := range currentNode.Nodes {
		selectedNode = getUpperClosest(size, *n, selectedNode)
	}
	return selectedNode
}

func partTwo(origin *tools.GraphNode[int]) {
	target := 30000000 - (70000000 - origin.Value)
	selected := getUpperClosest(target, *origin, origin)
	fmt.Println(selected.Value)
}

func Main(path string) {
	s := tools.Readfile(path)

	lines := strings.Split(s, "\n")
	origin := constructFS(lines)
	partOne(origin)
	partTwo(origin)
}
