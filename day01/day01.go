package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Segmev/AdventOfCode2022/tools"
)

func main() {
	s := tools.Readfile("./input.txt")

	lines := strings.Split(s, "\n")

	var elfSumCalories []int
	var elfCalories []int

	for _, line := range lines {
		if len(line) > 0 {
			caloriesValue, _ := strconv.Atoi(line)
			elfCalories = append(elfCalories, caloriesValue)
		} else {
			acc := 0
			for _, v := range elfCalories {
				acc += v
			}
			elfSumCalories = append(elfSumCalories, acc)
			elfCalories = elfCalories[:0]
		}
	}

	sort.Slice(elfSumCalories, func(i, j int) bool {
		return elfSumCalories[i] > elfSumCalories[j]
	})

	fmt.Println(elfSumCalories[0])
	fmt.Println(elfSumCalories[0] + elfSumCalories[1] + elfSumCalories[2])
}
