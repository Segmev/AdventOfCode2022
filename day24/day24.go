package day24

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Segmev/AdventOfCode2022/tools"
)

type Entity struct {
	X, Y       int
	dirX, dirY int
	c          string
}

type Zone struct {
	winds      map[string][]*Entity
	allWinds   []map[string][]*Entity
	start, end Entity
	maxX, maxY int
}

type Move struct {
	e       Entity
	count   int
	targets []Entity
}

func (z Zone) Clone() (newZone Zone) {
	newZone.maxX = z.maxX
	newZone.maxY = z.maxY
	newZone.start = z.start
	newZone.end = z.end
	newZone.winds = tools.CloneMap(z.winds)
	return newZone
}

func (e Entity) GetId() string {
	return fmt.Sprintf("%d_%d", e.X, e.Y)
}

func (e Entity) Clone() (ne Entity) {
	ne.X = e.X
	ne.Y = e.Y
	ne.dirX = e.dirX
	ne.dirY = e.dirY
	ne.c = e.c

	return ne
}

func (e *Entity) Move(maxX, maxY int) {
	e.X = tools.Modulo(e.X+e.dirX, maxX)
	e.Y = tools.Modulo(e.Y+e.dirY, maxY)
}

func (z *Zone) MoveWinds() {
	winds := map[string][]*Entity{}
	for k := range z.winds {
		for ie, e := range z.winds[k] {
			z.winds[k][ie].Move(z.maxX, z.maxY)
			winds[z.winds[k][ie].GetId()] = AddToSlice(winds[e.GetId()], e)
		}
	}
	z.winds = winds
}

func (z *Zone) GetNextMovedWinds(w map[string][]*Entity) map[string][]*Entity {
	winds := map[string][]*Entity{}
	for k := range w {
		for ie, e := range w[k] {
			w[k][ie].Move(z.maxX, z.maxY)
			winds[w[k][ie].GetId()] = AddToSlice(winds[w[k][ie].GetId()], e)
		}
	}
	return winds
}

func (z *Zone) GenerateWinds(originalWinds map[string][]*Entity) {
	z.allWinds = make([]map[string][]*Entity, z.maxX*z.maxY)
	for i := 0; i < z.maxX*z.maxY; i++ {
		winds := map[string][]*Entity{}
		for k := range originalWinds {
			for _, e := range originalWinds[k] {
				ne := e.Clone()
				winds[k] = AddToSlice(winds[k], &ne)
			}
		}
		originalWinds = z.GetNextMovedWinds(originalWinds)
		z.allWinds[i] = winds
	}
}

func constructMap(lines []string) (zone Zone) {
	zone.start, zone.end = Entity{Y: -1}, Entity{Y: len(lines) - 2}
	zone.maxX = len(lines[0]) - 2
	zone.maxY = len(lines) - 2

	for i, c := range lines[0] {
		if c == '.' {
			zone.start.X = i - 1
		}
	}
	for i, c := range lines[len(lines)-1] {
		if c == '.' {
			zone.end.X = i - 1
		}
	}

	winds := map[string][]*Entity{}
	for y, line := range lines[1 : len(lines)-1] {
		for x, c := range line[1 : len(line)-1] {
			switch c {
			case '>':
				e := Entity{X: x, Y: y, dirX: 1, c: ">"}
				winds[e.GetId()] = AddToSlice(winds[e.GetId()], &e)
			case '<':
				e := Entity{X: x, Y: y, dirX: -1, c: "<"}
				winds[e.GetId()] = AddToSlice(winds[e.GetId()], &e)
			case '^':
				e := Entity{X: x, Y: y, dirY: -1, c: "^"}
				winds[e.GetId()] = AddToSlice(winds[e.GetId()], &e)
			case 'v':
				e := Entity{X: x, Y: y, dirY: 1, c: "v"}
				winds[e.GetId()] = AddToSlice(winds[e.GetId()], &e)
			}
		}
	}
	zone.GenerateWinds(winds)
	return zone
}

func AddToSlice[T any](s []T, e T) []T {
	if s == nil {
		return []T{e}
	}
	return append(s, e)
}

func ManhDist(a, b Entity) int {
	return tools.Abs(a.X-b.X) + tools.Abs(a.Y-b.Y)
}

func (z Zone) getNextMoves(e Entity, count int) (es []Entity) {
	for i := -1; i <= 1; i += 1 {
		tempX, tempY := e.X+i, e.Y
		if !(tempX < 0 || tempX > z.maxX || tempY < 0 || tempY >= z.maxY) ||
			(tempX == z.start.X && tempY == z.start.Y) ||
			(tempX == z.end.X && tempY == z.end.Y) {
			if _, ok := z.allWinds[count%(z.maxX*z.maxY)][fmt.Sprintf("%d_%d", tempX, tempY)]; !ok {
				es = append(es, Entity{X: tempX, Y: tempY})
			}
		}
		tempX, tempY = e.X, e.Y+i
		if !(tempX < 0 || tempX > z.maxX || tempY < 0 || tempY >= z.maxY) ||
			(tempX == z.start.X && tempY == z.start.Y) ||
			(tempX == z.end.X && tempY == z.end.Y) {
			if _, ok := z.allWinds[count%(z.maxX*z.maxY)][fmt.Sprintf("%d_%d", tempX, tempY)]; !ok {
				es = append(es, Entity{X: tempX, Y: tempY})
			}
		}
	}
	if _, ok := z.allWinds[count%(z.maxX*z.maxY)][fmt.Sprintf("%d_%d", e.X, e.Y)]; !ok {
		es = append(es, Entity{X: e.X, Y: e.Y})
	}
	return es
}

func (m Move) displayMap(z Zone) {
	fmt.Println("Move at", m.count)
	for y := 0; y <= z.maxY; y++ {
		fmt.Print("#")
		for x := 0; x < z.maxX; x++ {
			if x == m.e.X && y == m.e.Y {
				fmt.Print("E")
			} else {
				ce, ok := z.allWinds[m.count][fmt.Sprintf("%d_%d", x, y)]
				if !ok {
					fmt.Print(".")
				} else {
					if len(ce) > 1 {
						fmt.Print(len(ce))
					} else {
						fmt.Print(ce[0].c)
					}
				}
			}
		}
		fmt.Println("#")
	}
	fmt.Println()
}

func partOne(zone Zone) {
	movesInQueue := tools.GetSet[string]()
	movesInQueue.Add(zone.start.GetId())
	queue := []Move{{count: 0, e: Entity{X: zone.start.X, Y: zone.start.Y}}}
	bestCount := 500
	for len(queue) > 0 {
		sort.SliceStable(queue, func(i, j int) bool {
			return (ManhDist(queue[i].e, zone.end) + queue[i].count) < (ManhDist(queue[j].e, zone.end) + queue[j].count)
		})

		move := queue[0]
		queue = queue[1:]
		if move.e.X == zone.end.X && move.e.Y == zone.end.Y {
			if bestCount > move.count {
				bestCount = move.count
			}
		}
		if move.count > bestCount {
			continue
		}
		for _, e := range zone.getNextMoves(move.e, move.count) {
			if !movesInQueue.Contains(fmt.Sprintf("%s_%d", e.GetId(), ((move.count + 1) % (zone.maxX * zone.maxY)))) {
				queue = append(queue, Move{e: e, count: move.count + 1})
				movesInQueue.Add(fmt.Sprintf("%s_%d", e.GetId(), (move.count+1)%(zone.maxX*zone.maxY)))
			}
		}
	}
	fmt.Println(bestCount - 1)
}

func partTwo(zone Zone) {
	movesInQueue := tools.GetSet[string]()
	movesInQueue.Add(zone.start.GetId())
	queue := []Move{{count: 0, e: Entity{X: zone.start.X, Y: zone.start.Y}, targets: []Entity{
		{X: zone.end.X, Y: zone.end.Y},
		{X: zone.start.X, Y: zone.start.Y},
		{X: zone.end.X, Y: zone.end.Y},
	}}}
	bestCount := 1000

	for len(queue) > 0 {
		sort.SliceStable(queue, func(i, j int) bool {
			return (ManhDist(queue[i].e, queue[i].targets[0]) + queue[i].count) < (ManhDist(queue[j].e, queue[j].targets[0]) + queue[j].count)
		})
		move := queue[0]
		queue = queue[1:]

		if move.e.X == move.targets[0].X && move.e.Y == move.targets[0].Y {
			move.targets = move.targets[1:]
			if len(move.targets) == 0 {
				if bestCount > move.count {
					bestCount = move.count
				}
				continue
			}
		}
		if move.count > bestCount {
			continue
		}
		for _, e := range zone.getNextMoves(move.e, move.count) {
			if !movesInQueue.Contains(fmt.Sprintf("%s_%d_%d", e.GetId(), ((move.count + 1) % (zone.maxX * zone.maxY)), len(move.targets))) {
				queue = append(queue, Move{e: e, count: move.count + 1, targets: tools.CloneSlice(move.targets)})
				movesInQueue.Add(fmt.Sprintf("%s_%d_%d", e.GetId(), (move.count+1)%(zone.maxX*zone.maxY), len(move.targets)))
			}
		}
	}
	fmt.Println(bestCount - 1)
}

func Main(path string) {
	s := tools.Readfile(path)

	lines := strings.Split(s, "\n")
	lines = lines[:len(lines)-1]
	z := constructMap(lines)

	partOne(z)
	partTwo(z)
}
