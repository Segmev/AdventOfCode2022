package main

import (
	"fmt"
	"os"

	"github.com/Segmev/AdventOfCode2022/day01"
	"github.com/Segmev/AdventOfCode2022/day02"
	"github.com/Segmev/AdventOfCode2022/day03"
	"github.com/Segmev/AdventOfCode2022/day04"
	"github.com/Segmev/AdventOfCode2022/day05"
	"github.com/Segmev/AdventOfCode2022/day06"
	"github.com/Segmev/AdventOfCode2022/day07"
	"github.com/Segmev/AdventOfCode2022/day08"
	"github.com/Segmev/AdventOfCode2022/day09"
	"github.com/Segmev/AdventOfCode2022/day10"
	"github.com/Segmev/AdventOfCode2022/day11"
	"github.com/Segmev/AdventOfCode2022/day12"
	"github.com/Segmev/AdventOfCode2022/day13"
	"github.com/Segmev/AdventOfCode2022/day14"
	"github.com/Segmev/AdventOfCode2022/day15"
	"github.com/Segmev/AdventOfCode2022/day16"
	"github.com/Segmev/AdventOfCode2022/day17"
	"github.com/Segmev/AdventOfCode2022/day18"
	"github.com/Segmev/AdventOfCode2022/day19"
	"github.com/Segmev/AdventOfCode2022/day20"
	"github.com/Segmev/AdventOfCode2022/day21"
)

func runningFn(dailyFuncs map[string]func(string), arg string) {
	fmt.Println("Running", arg)
	fmt.Println()
	dailyFuncs[arg](fmt.Sprintf("./%s/input.txt", arg))
	fmt.Println("=========================================================")
}

func main() {
	dailyFuncs := map[string]func(string){
		"day01": day01.Main,
		"day02": day02.Main,
		"day03": day03.Main,
		"day04": day04.Main,
		"day05": day05.Main,
		"day06": day06.Main,
		"day07": day07.Main,
		"day08": day08.Main,
		"day09": day09.Main,
		"day10": day10.Main,
		"day11": day11.Main,
		"day12": day12.Main,
		"day13": day13.Main,
		"day14": day14.Main,
		"day15": day15.Main,
		"day16": day16.Main,
		"day17": day17.Main,
		"day18": day18.Main,
		"day19": day19.Main,
		"day20": day20.Main,
		"day21": day21.Main,
	}
	if len(os.Args) < 2 {
		for i := 1; i <= len(dailyFuncs); i++ {
			runningFn(dailyFuncs, fmt.Sprintf("day%.02d", i))
		}
	} else {
		for _, arg := range os.Args[1:] {
			if _, ok := dailyFuncs[arg]; ok {
				runningFn(dailyFuncs, arg)
			} else {
				fmt.Println()
				fmt.Println("Invalid", arg)
				fmt.Println()
			}
		}
	}
}
