package day20

import (
	"fmt"
	"strings"

	"github.com/Segmev/AdventOfCode2022/tools"
)

func IndexOf(nb *int, numbers []*int) int {
	for i := range numbers {
		if numbers[i] == nb {
			return i
		}
	}
	return -1
}

func calculateSum(numbers []*int, cycleNb int) {
	refs := []*int{}
	for i := range numbers {
		refs = append(refs, numbers[i])
	}

	for cycle := 0; cycle < cycleNb; cycle++ {
		for i := 0; i < len(refs); i++ {
			index := IndexOf(refs[i], numbers)
			popNb := numbers[index]
			numbers = append(numbers[:index], numbers[index+1:]...)
			posDiff := *popNb + index
			insertAt := ((posDiff % len(numbers)) + len(numbers)) % len(numbers)
			if insertAt == 0 {
				insertAt = len(numbers)
			}
			numbers = append(numbers[:insertAt], append([]*int{popNb}, numbers[insertAt:]...)...)
		}
	}
	i := 0
	for i < len(numbers) {
		if *numbers[i] == 0 {
			break
		}
		i++
	}
	fmt.Println(*numbers[(i+1000)%len(numbers)] + *numbers[(i+2000)%len(numbers)] + *numbers[(i+3000)%len(numbers)])
}

func partOne(lines []string) {
	numbers := []*int{}
	for _, line := range lines {
		nb := tools.Atoi(line)
		numbers = append(numbers, &nb)
	}

	calculateSum(numbers, 1)
}

func partTwo(lines []string) {
	numbers := []*int{}
	for _, line := range lines {
		nb := tools.Atoi(line) * 811589153
		numbers = append(numbers, &nb)
	}

	calculateSum(numbers, 10)
}

func Main(path string) {
	s := tools.Readfile(path)

	lines := strings.Split(s, "\n")
	lines = lines[:len(lines)-1]
	partOne(lines)
	partTwo(lines)
}
