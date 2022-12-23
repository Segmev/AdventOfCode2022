package day19

import (
	"fmt"
	"strings"

	"github.com/Segmev/AdventOfCode2022/tools"
)

const (
	ORE      = "ore"
	CLAY     = "clay"
	OBSIDIAN = "obsidian"
	GEODE    = "geode"
)

type Robot struct {
	cost         map[string]int
	resourceName string
}

func buildRobot(robotBlueprint Robot, resources map[string]int, robots map[string]int) {
	for key, vCost := range robotBlueprint.cost {
		resources[key] -= vCost
	}
	robots[robotBlueprint.resourceName] += 1
}

func minutesBeforeRobot(robotBlueprint Robot, resources, robots map[string]int) int {
	min := 1
	for key, cost := range robotBlueprint.cost {
		if robots[key] == 0 {
			return 1000
		}
		necessaryResources := cost - resources[key]
		if necessaryResources < 0 {
			necessaryResources = 0
		}
		minimumNecessaryTime := (necessaryResources) / (robots[key])
		for necessaryResources-minimumNecessaryTime*robots[key] > 0 || minimumNecessaryTime == 0 {
			minimumNecessaryTime += 1
		}
		if resGenerationTime := minimumNecessaryTime; min < resGenerationTime {
			min = resGenerationTime
		}
	}
	return min
}

func letRobotsWorkFor(minutes int, resources, robots map[string]int) {
	for resKey, robotCount := range robots {
		resources[resKey] += robotCount * minutes
	}
}

func getBestCaseGeodeCount(resources, robots map[string]int, minutes int) int {
	return resources[GEODE] + (robots[GEODE] * (robots[GEODE] + 1) / 2) + minutes*5
}

func getWorstCostFor(resourceName string, blueprints []*Robot) int {
	cost := 0
	for _, blueprint := range blueprints {
		if cost < blueprint.cost[resourceName] {
			cost = blueprint.cost[resourceName]
		}
	}
	return cost
}

func getGeodeNumbers(blueprints []*Robot, resources map[string]int, robots map[string]int, minutes int, bestGeodeCount *int, robotBuilt []string) {
	if currentBestCaseCount := getBestCaseGeodeCount(resources, robots, minutes); currentBestCaseCount < *bestGeodeCount {
		// trim branch if cannot be better
		return
	}
	if robots[GEODE] > 0 {
		geodeCount := resources[GEODE] + (minutes)*robots[GEODE]
		if geodeCount > *bestGeodeCount {
			*bestGeodeCount = geodeCount
		}
	}
	// fmt.Println(minutes, "minutes", robots, resources)

	for i := len(blueprints) - 1; i >= 0; i-- {
		robot := blueprints[i]
		if robot.resourceName == GEODE || (robots[robot.resourceName] < getWorstCostFor(robot.resourceName, blueprints)) {
			if neededTime := minutesBeforeRobot(*robot, resources, robots); neededTime < minutes {
				// fmt.Println("building", robot, "for", neededTime)
				futurResources := tools.CloneMap(resources)
				futurRobots := tools.CloneMap(robots)
				letRobotsWorkFor(neededTime, futurResources, futurRobots)
				futurResources[robot.resourceName] -= 1
				buildRobot(*robot, futurResources, futurRobots)
				getGeodeNumbers(blueprints, futurResources, futurRobots, minutes-neededTime,
					bestGeodeCount, append(robotBuilt, fmt.Sprintf("[%d-%s]", minutes-neededTime-1, robot.resourceName)))
			}
		}
	}
}

func partOne(lines []string) {
	total := 0
	for i, situation := range parseBluePrints(lines) {
		geodeCount := 0
		getGeodeNumbers(situation, map[string]int{ORE: 0, CLAY: 0, OBSIDIAN: 0, GEODE: 0}, map[string]int{ORE: 1, CLAY: 0, OBSIDIAN: 0, GEODE: 0}, 24, &geodeCount, []string{ORE})
		fmt.Println(i+1, "=>", geodeCount)
		total += (i + 1) * (geodeCount)
	}
	fmt.Println(total)
}

func partTwo(lines []string) {
	total := 1
	for i, situation := range parseBluePrints(lines[2:3]) {
		geodeCount := 0
		getGeodeNumbers(situation, map[string]int{ORE: 0, CLAY: 0, OBSIDIAN: 0, GEODE: 0}, map[string]int{ORE: 1, CLAY: 0, OBSIDIAN: 0, GEODE: 0}, 32, &geodeCount, []string{ORE})
		fmt.Println(i+1, "=>", geodeCount)
		total *= geodeCount
	}
	fmt.Println(total)
}

func parseBluePrints(lines []string) (situations [][]*Robot) {
	for _, line := range lines {
		line = strings.Split(line, ": ")[1]
		instr := strings.Split(line, "Each ")[1:]
		oreBot := Robot{resourceName: ORE, cost: map[string]int{ORE: tools.Atoi(strings.Split(instr[0], " ")[3])}}
		clayBot := Robot{resourceName: CLAY, cost: map[string]int{ORE: tools.Atoi(strings.Split(instr[1], " ")[3])}}
		obsiBot := Robot{resourceName: OBSIDIAN, cost: map[string]int{ORE: tools.Atoi(strings.Split(instr[2], " ")[3]), CLAY: tools.Atoi(strings.Split(instr[2], " ")[6])}}
		geoBot := Robot{resourceName: GEODE, cost: map[string]int{ORE: tools.Atoi(strings.Split(instr[3], " ")[3]), OBSIDIAN: tools.Atoi(strings.Split(instr[3], " ")[6])}}

		situations = append(situations, []*Robot{&oreBot, &clayBot, &obsiBot, &geoBot})
	}
	return situations
}

func Main(path string) {
	s := tools.Readfile(path)

	lines := strings.Split(s, "\n")
	partOne(lines[:len(lines)-1])
	partTwo(lines[:len(lines)-1])
}
