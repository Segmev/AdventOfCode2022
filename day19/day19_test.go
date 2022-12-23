package day19

import "testing"

func TestLetRobotWorkFor(t *testing.T) {
	resources := map[string]int{
		ORE:      1,
		CLAY:     2,
		OBSIDIAN: 3,
	}
	robots := map[string]int{
		ORE:      2,
		CLAY:     3,
		OBSIDIAN: 4,
	}

	letRobotWorkFor(2, resources, robots)
	if resources[ORE] != 5 {
		t.Fatal("error")
	}
}
