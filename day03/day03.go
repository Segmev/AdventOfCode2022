package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/Segmev/AdventOfCode2022/tools"
)

func partOne(lines []string) {
	itemsLetters := "abcdefghijklmnopqrstuvwxyz"

	sum := 0
	for _, line := range lines {
		firstRucksack := line[:len(line)/2]
		secondRuckSack := line[len(line)/2:]

		for i := range firstRucksack {
			if strings.Contains(secondRuckSack, firstRucksack[i:i+1]) {

				itemVal := tools.IndexInSlice([]rune(itemsLetters), unicode.ToLower(rune(firstRucksack[i])))
				if unicode.IsUpper(rune(firstRucksack[i])) {
					sum += 26
				}
				sum += 1 + itemVal
				break
			}
		}
	}
	fmt.Println(sum)
}

func partTwo(lines []string) {
	itemsLetters := "abcdefghijklmnopqrstuvwxyz"
	sum := 0
	for i := 0; i < len(lines)-1; i += 3 {
		for letterIdx := range lines[i] {
			if strings.Contains(lines[i+1], lines[i][letterIdx:letterIdx+1]) && strings.Contains(lines[i+2], lines[i][letterIdx:letterIdx+1]) {
				itemVal := tools.IndexInSlice([]rune(itemsLetters), unicode.ToLower(rune(lines[i][letterIdx])))
				if unicode.IsUpper(rune(rune(lines[i][letterIdx]))) {
					sum += 26
				}
				sum += 1 + itemVal
				break
			}
		}
	}
	fmt.Println(sum)
}

func main() {
	s := tools.Readfile("./input.txt")

	lines := strings.Split(s, "\n")
	partOne(lines)
	partTwo(lines)
}
