package day17

import (
	"fmt"
	"strings"

	"github.com/Segmev/AdventOfCode2022/tools"
)

type Coor struct {
	x, y int

	id string
}

func (c *Coor) AddX(x int) {
	c.SetX(c.x + x)
}

func (c *Coor) AddY(y int) {
	c.SetY(c.y + y)
}

func (c *Coor) SetY(y int) {
	c.y = y
	c.id = ""
}

func (c *Coor) SetX(x int) {
	c.x = x
	c.id = ""
}

func (c *Coor) Id() string {
	if c.id == "" {
		c.id = fmt.Sprintf("%d_%d", c.x, c.y)
	}
	return c.id
}

const (
	boundLeftX  = -1
	boundRightX = 7
)

func canApplyForceOn(m tools.Set[string], rocks []Coor, appliedForce int) bool {
	for _, r := range rocks {
		movedRock := Coor{x: r.x + appliedForce, y: r.y}
		if movedRock.x <= boundLeftX || movedRock.x >= boundRightX || m.Contains(movedRock.Id()) {
			return false
		}
	}
	return true
}

func canMoveDown(m tools.Set[string], rocks []Coor) bool {
	for _, r := range rocks {
		movedRock := Coor{x: r.x, y: r.y - 1}
		if m.Contains(movedRock.Id()) {
			return false
		}
	}
	return true
}

func getRocks(id int, shiftX int, topY int) []Coor {
	rockForms := map[int][]Coor{
		0: {{x: 0 + shiftX, y: topY + 0}, {x: 1 + shiftX, y: topY}, {x: 2 + shiftX, y: topY + 0}, {x: 3 + shiftX, y: topY}},                                   // -
		1: {{x: 1 + shiftX, y: topY + 2}, {x: 1 + shiftX, y: topY}, {x: 0 + shiftX, y: topY + 1}, {x: 1 + shiftX, y: topY + 1}, {x: 2 + shiftX, y: topY + 1}}, // +
		2: {{x: 2 + shiftX, y: topY + 2}, {x: 2 + shiftX, y: topY + 1}, {x: 2 + shiftX, y: topY}, {x: 1 + shiftX, y: topY}, {x: 0 + shiftX, y: topY}},         // ⅃
		3: {{x: 0 + shiftX, y: topY + 3}, {x: 0 + shiftX, y: topY}, {x: 0 + shiftX, y: topY + 1}, {x: 0 + shiftX, y: topY + 2}},                               // |
		4: {{x: 1 + shiftX, y: topY + 1}, {x: 0 + shiftX, y: topY}, {x: 1 + shiftX, y: topY + 0}, {x: 0 + shiftX, y: topY + 1}},                               // #
	}
	return rockForms[id%5]
}

func isNewFloor(m tools.Set[string], y int) bool {
	for x := 1 + boundLeftX; x < boundRightX; x++ {
		if !m.Contains(fmt.Sprintf("%d_%d", x, y)) {
			return false
		}
	}
	return true
}

func partOne(line string) {
	m := tools.GetSet[string]()
	for i := 0; i <= 7; i++ {
		m.Add(fmt.Sprintf("%d_0", i))
	}
	topY := 0
	rocks := getRocks(0, 2, topY+4)

	forceCount := 0
	for rocksCount := 0; rocksCount < 2022; {
		var force int
		if line[forceCount%(len(line))] == '>' {
			force = 1
		} else {
			force = -1
		}

		if canApplyForceOn(m, rocks, force) {
			for rockI := range rocks {
				rocks[rockI].AddX(force)
			}
		}

		if canMoveDown(m, rocks) {
			for rockI := range rocks {
				rocks[rockI].AddY(-1)
			}
		} else {
			for _, rock := range rocks {
				m.Add(rock.Id())
				if topY < rock.y {
					topY = rock.y
				}
			}

			if isNewFloor(m, topY) {
				fmt.Println("new floor at", topY, "for", rocksCount, rocksCount%5, forceCount%len(line))
			}

			rocksCount++
			rocks = getRocks(rocksCount, 2, topY+4)
		}
		forceCount++
	}
	fmt.Println(topY)
}

const part2Limit = 1000000000000

func ceilingShot(m tools.Set[string], topY int) [7]int {
	ceiling := [7]int{}
	for x := 0; x < 7; x++ {
		y := topY
		for !m.Contains(fmt.Sprintf("%d_%d", x, y)) {
			y--
		}
		ceiling[x] = topY - y
	}
	return ceiling
}

type CeilingShot struct {
	rockCount int
	ceiling   [7]int
	Y         int
}

func partTwo(line string) {
	topYdiff := map[string]CeilingShot{}
	m := tools.GetSet[string]()
	for i := 0; i <= 7; i++ {
		m.Add(fmt.Sprintf("%d_0", i))
	}
	topY := 0
	rocks := getRocks(0, 2, topY+4)

	firstCycleRockKey := ""
	hasJumped := false

	forceCount := 0
	for rocksCount := 0; rocksCount < part2Limit; {

		var force int
		if line[forceCount%(len(line))] == '>' {
			force = 1
		} else {
			force = -1
		}

		if canApplyForceOn(m, rocks, force) {
			for rockI := range rocks {
				rocks[rockI].AddX(force)
			}
		}

		if canMoveDown(m, rocks) {
			for rockI := range rocks {
				rocks[rockI].AddY(-1)
			}
		} else {
			for _, rock := range rocks {
				m.Add(rock.Id())

				if topY < rock.y {
					topY = rock.y
				}
			}

			if _, ok := topYdiff[fmt.Sprint(ceilingShot(m, topY), rocksCount%5, forceCount%len(line))]; ok && !hasJumped {
				if firstCycleRockKey == "" {
					firstCycleRockKey = fmt.Sprint(ceilingShot(m, topY), rocksCount%5, forceCount%len(line))
				} else if firstCycleRockKey == fmt.Sprint(ceilingShot(m, topY), rocksCount%5, forceCount%len(line)) {
					firstOccurence := topYdiff[firstCycleRockKey]
					fmt.Println("first at", firstOccurence.rockCount, "second at", rocksCount)
					rocksCountRepeat := rocksCount - firstOccurence.rockCount
					repeatYDiff := topY - firstOccurence.Y
					repeatTo := (part2Limit - rocksCount) / rocksCountRepeat
					topY += (repeatTo) * repeatYDiff

					rocksCount = part2Limit - ((part2Limit - rocksCount) - (rocksCountRepeat * (repeatTo)))

					fmt.Println(rocksCountRepeat, repeatYDiff, repeatTo, rocksCount, topY)

					m.Clear()
					for x := 0; x < 7; x++ {
						m.Add(fmt.Sprintf("%d_%d", x, topY-topYdiff[firstCycleRockKey].ceiling[x]))
					}
					hasJumped = true
					forceCount++
				}
			}
			topYdiff[fmt.Sprint(ceilingShot(m, topY), rocksCount%5, forceCount%len(line))] = CeilingShot{rocksCount, ceilingShot(m, topY), topY}

			rocksCount++
			rocks = getRocks(rocksCount, 2, topY+4)

		}
		forceCount++
	}
	fmt.Println(topY)
}

func Main(path string) {
	s := tools.Readfile(path)

	lines := strings.Split(s, "\n")
	partOne(lines[0])
	partTwo(lines[0])
}
