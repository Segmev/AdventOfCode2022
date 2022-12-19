package day16

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Segmev/AdventOfCode2022/tools"
)

const nb_routine = 128

type GraphNode struct {
	Id          string
	Value       int
	StepToNodes map[string]int
	Nodes       map[string]*GraphNode
}

func parseLine(line string) (string, int, []string) {
	re := regexp.MustCompile("valves* ")
	lineParts := strings.Split(line, ";")
	firstPart := strings.Split(lineParts[0], " ")
	ratePart := strings.Split(firstPart[len(firstPart)-1], "=")[1]
	valvesPart := re.Split(lineParts[1], -1)[1]

	valves := strings.Split(valvesPart, ", ")
	rate, _ := strconv.Atoi(ratePart)

	return firstPart[1], rate, valves
}

func getSumOfReleased(released []int) (sum int) {
	for _, rv := range released {
		sum += rv
	}
	return sum
}

func mapStepsToNode(start, end *GraphNode, interestingNodes map[string]int, interestingNodesSteps map[string]int) {
	type NodeWithSteps struct {
		node *GraphNode
		step int
	}
	visited := tools.Set[string]{}
	toVisit := []*NodeWithSteps{{start, 0}}

	for len(toVisit) > 0 {
		studiedNode := toVisit[0]
		toVisit = toVisit[1:]
		visited.Add(studiedNode.node.Id)
		for _, node := range studiedNode.node.Nodes {
			if visited.Contains(node.Id) {
				continue
			}
			if node.Id == end.Id {
				seKey := fmt.Sprintf("%s_%s", start.Id, node.Id)
				if v, ok := interestingNodesSteps[seKey]; !ok || v > studiedNode.step+1 {
					interestingNodesSteps[seKey] = studiedNode.step + 1
					interestingNodesSteps[fmt.Sprintf("%s_%s", node.Id, start.Id)] = studiedNode.step + 1
				}
			} else {
				toVisit = append(toVisit, &NodeWithSteps{node, studiedNode.step + 1})
			}
		}
	}
}

func getRouteScore(toVisitNodes tools.Set[string], node string, interestingNodes, interestingNodesSteps map[string]int, stepsLeft int) (int, string) {
	if stepsLeft < 0 {
		return 0, ""
	}
	best := 0
	nodesPathBest := ""
	toVisitNodes.Delete(node)
	for toVisitNode := range toVisitNodes {
		score, nodesPath := getRouteScore(toVisitNodes.Clone(), toVisitNode, interestingNodes, interestingNodesSteps, stepsLeft-1-interestingNodesSteps[fmt.Sprintf("%s_%s", node, toVisitNode)])
		if score > best {
			best = score
			nodesPathBest = nodesPath
		}
	}
	return best + interestingNodes[node]*stepsLeft, node + "=>" + fmt.Sprint(stepsLeft) + "   " + nodesPathBest
}

func partOne(lines []string) {
	allNodes := map[string]*GraphNode{}
	interestingNodes := map[string]int{}
	toBeLinked := map[string][]string{}
	for _, line := range lines {
		name, rate, linkedTo := parseLine(line)
		toBeLinked[name] = linkedTo

		allNodes[name] = &GraphNode{Id: name, Value: rate, Nodes: map[string]*GraphNode{}, StepToNodes: map[string]int{}}
		if rate > 0 {
			interestingNodes[name] = rate
		}
	}
	for key := range toBeLinked {
		for _, linkedToName := range toBeLinked[key] {
			allNodes[key].Nodes[linkedToName] = allNodes[linkedToName]
		}
	}

	interestingNodesSteps := map[string]int{}
	interestingNodesSet := tools.GetSet[string]()

	for nodeName := range interestingNodes {
		mapStepsToNode(allNodes["AA"], allNodes[nodeName], interestingNodes, interestingNodesSteps)
		for otherNodeName := range interestingNodes {
			if otherNodeName != nodeName {
				mapStepsToNode(allNodes[nodeName], allNodes[otherNodeName], interestingNodes, interestingNodesSteps)
			}
		}
	}

	for k := range interestingNodes {
		interestingNodesSet.Add(k)
	}
	fmt.Println(getRouteScore(interestingNodesSet, "AA", interestingNodes, interestingNodesSteps, 30))
}

func partTwo(lines []string) {
	interestingNodesKeys := []string{}
	allNodes := map[string]*GraphNode{}
	interestingNodes := map[string]int{}
	toBeLinked := map[string][]string{}
	for _, line := range lines {
		name, rate, linkedTo := parseLine(line)
		toBeLinked[name] = linkedTo

		allNodes[name] = &GraphNode{Id: name, Value: rate, Nodes: map[string]*GraphNode{}, StepToNodes: map[string]int{}}
		if rate > 0 {
			interestingNodes[name] = rate
			interestingNodesKeys = append(interestingNodesKeys, name)
		}
	}
	for key := range toBeLinked {
		for _, linkedToName := range toBeLinked[key] {
			allNodes[key].Nodes[linkedToName] = allNodes[linkedToName]
		}
	}

	interestingNodesSteps := map[string]int{}

	for nodeName := range interestingNodes {
		mapStepsToNode(allNodes["AA"], allNodes[nodeName], interestingNodes, interestingNodesSteps)
		for otherNodeName := range interestingNodes {
			if otherNodeName != nodeName {
				mapStepsToNode(allNodes[nodeName], allNodes[otherNodeName], interestingNodes, interestingNodesSteps)
			}
		}
	}

	pairs := twoPartitions(interestingNodesKeys)
	channel := make(chan int)

	for i := 0; i < nb_routine; i++ {
		go getPairsRouteScores(channel, pairs[i*len(pairs)/nb_routine:(i+1)*len(pairs)/nb_routine], interestingNodesKeys, interestingNodes, interestingNodesSteps)
	}

	bestScore := 0
	for i := 0; i < nb_routine; i++ {
		score := <-channel
		if score > bestScore {
			bestScore = score
		}
	}

	fmt.Println(bestScore)
}

func getPairsRouteScores(channel chan int, pairs [][][]string, interestingNodesKeys []string, interestingNodes, interestingNodesSteps map[string]int) {
	bestScore := 0
	for _, pair := range pairs {
		if len(pair[0]) < len(interestingNodesKeys)/2 || len(pair[1]) < len(interestingNodesKeys)/2 {
			continue
		}
		scoreA, _ := getRouteScore(tools.GetSetFrom(pair[0]), "AA", interestingNodes, interestingNodesSteps, 26)
		scoreB, _ := getRouteScore(tools.GetSetFrom(pair[1]), "AA", interestingNodes, interestingNodesSteps, 26)

		if bestScore < scoreA+scoreB {
			bestScore = scoreA + scoreB
		}
	}
	channel <- bestScore
}

func twoPartitions(s []string) [][][]string {
	resList := [][][]string{}
	for l := 0; l <= len(s)/2; l++ {
		combis := combinations(s, l)
		for _, c := range combis {
			resList = append(resList, [][]string{(c), (difference(s, c))})
		}
	}
	return resList
}

func combinations(s []string, l int) [][]string {
	if l == 0 {
		return [][]string{{}}
	}
	if len(s) == 0 {
		return [][]string{}
	}
	c := combinations(s[1:], l-1)
	for i := range c {
		c[i] = append(c[i], s[0])
	}
	return append(c, combinations(s[1:], l)...)
}

func difference(s, c []string) []string {
	m := make(map[string]bool)
	for _, e := range c {
		m[e] = true
	}
	res := []string{}
	for _, e := range s {
		if !m[e] {
			res = append(res, e)
		}
	}
	return res
}

func Main(path string) {
	s := tools.Readfile(path)

	lines := strings.Split(s, "\n")
	partOne(lines[:len(lines)-1])
	timeStart := time.Now()
	partTwo(lines[:len(lines)-1])
	fmt.Println(time.Since(timeStart), "for part2")
}
