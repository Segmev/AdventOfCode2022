package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Segmev/AdventOfCode2022/tools"
)

type Tree struct {
	height int
	amount int
}

func updateRowVisibilities(row map[int][]*Tree) {
	for rowIndex := 0; rowIndex < len(row); rowIndex++ {
		trees := row[rowIndex]
		upBound := trees[0].height
		for idx, tree := range trees {
			if tree.height <= upBound && idx != 0 {
				tree.amount -= 1
			} else if tree.height > upBound {
				upBound = tree.height
			}
			for rIdx := idx + 1; rIdx < len(trees); rIdx++ {
				if tree.height <= trees[rIdx].height {
					tree.amount -= 1
					break
				}
			}
		}
	}
}

func updateRowScore(row map[int][]*Tree) {
	for rowIndex := 0; rowIndex < len(row); rowIndex++ {
		trees := row[rowIndex]

		for idx, tree := range trees {
			if idx == 0 || idx == len(trees)-1 {
				tree.amount = tree.amount * 0
			} else {
				for lIdx := idx - 1; lIdx >= 0; lIdx-- {
					if tree.height <= trees[lIdx].height {
						tree.amount = tree.amount * (idx - lIdx)
						break
					} else if lIdx == 0 {
						tree.amount = tree.amount * idx
					}
				}

				for rIdx := idx + 1; rIdx < len(trees); rIdx++ {
					if tree.height <= trees[rIdx].height {
						tree.amount = tree.amount * (rIdx - idx)
						break
					} else if rIdx == len(trees)-1 {
						tree.amount = tree.amount * (rIdx - idx)
					}
				}
			}
		}
	}
}

func partOne(lines []string) {
	gridLines := map[int][]*Tree{}
	gridCols := map[int][]*Tree{}
	for j, line := range lines {
		for i := range line {
			height, _ := strconv.Atoi(line[i : i+1])
			t := Tree{height: height, amount: 4}

			gridLines[j] = append(gridLines[j], &t)
			gridCols[i] = append(gridCols[i], &t)
		}
	}

	updateRowVisibilities(gridCols)
	updateRowVisibilities(gridLines)

	total := 0
	count := 0
	for li := 0; li < len(gridLines); li++ {
		for _, tree := range gridLines[li] {
			total += 1
			if tree.amount > 0 {
				count += 1
			} else {
			}
		}
	}
	fmt.Println(count)
}

func partTwo(lines []string) {
	gridLines := map[int][]*Tree{}
	gridCols := map[int][]*Tree{}
	for j, line := range lines {
		for i := range line {
			height, _ := strconv.Atoi(line[i : i+1])
			t := Tree{height: height, amount: 1}

			gridLines[j] = append(gridLines[j], &t)
			gridCols[i] = append(gridCols[i], &t)
		}
	}

	updateRowScore(gridCols)
	updateRowScore(gridLines)

	max := 0
	for li := 0; li < len(gridLines); li++ {
		for _, tree := range gridLines[li] {
			if max < tree.amount {
				max = tree.amount
			}
		}
	}
	fmt.Println(max)
}

func main() {
	s := tools.Readfile("./input.txt")

	lines := strings.Split(s, "\n")
	lines = lines[:len(lines)-1]
	partOne(lines)
	partTwo(lines)
}
