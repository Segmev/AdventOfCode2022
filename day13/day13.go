package day13

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/Segmev/AdventOfCode2022/tools"
)

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func compareSlices(firstSlice []any, secondSlice []any) float64 {
	for i := 0; i < len(firstSlice) && i < len(secondSlice); i++ {
		fValue := reflect.ValueOf(firstSlice[i])
		sValue := reflect.ValueOf(secondSlice[i])
		if fValue.Kind() == reflect.Float64 && sValue.Kind() == reflect.Float64 {
			if fValue.Float() != sValue.Float() {
				return sValue.Float() - fValue.Float()
			}
		} else if fValue.Kind() == reflect.Slice && sValue.Kind() == reflect.Slice {
			fSlice, _ := fValue.Interface().([]any)
			sSlice, _ := sValue.Interface().([]any)
			if res := compareSlices(fSlice, sSlice); res != 0 {
				return res
			}
		} else if sValue.Kind() == reflect.Slice {
			fSlice := []any{fValue.Interface().(float64)}
			sSlice, _ := sValue.Interface().([]any)
			if res := compareSlices(fSlice, sSlice); res != 0 {
				return res
			}
		} else if fValue.Kind() == reflect.Slice {
			fSlice, _ := fValue.Interface().([]any)
			sSlice := []any{sValue.Interface().(float64)}
			if res := compareSlices(fSlice, sSlice); res != 0 {
				return res
			}
		}
	}
	return float64(len(secondSlice) - len(firstSlice))
}

func rightOrderStrings(lines []string) int {
	var firstLine []any
	var secondLine []any
	json.Unmarshal([]byte(lines[0]), &firstLine)
	json.Unmarshal([]byte(lines[1]), &secondLine)
	return int(compareSlices(firstLine, secondLine))
}

func partOne(input string) {
	sum := 0
	for i, packet := range strings.Split(input, "\n\n") {
		if rightOrderStrings(strings.Split(packet, "\n")) < 0 {
		} else {
			sum += i + 1
		}
	}
	fmt.Println(sum)
}

func partTwo(input string) {
	rawPackets := strings.Split(input, "\n")
	packets := []string{"[[2]]", "[[6]]"}
	for _, packet := range rawPackets {
		if packet != "" {
			packets = append(packets, packet)
		}
	}
	sort.Slice(packets, func(i, j int) bool {
		var firstLine []any
		var secondLine []any
		json.Unmarshal([]byte(packets[i]), &firstLine)
		json.Unmarshal([]byte(packets[j]), &secondLine)
		return compareSlices(firstLine, secondLine) > 0
	})
	a, b := 0, 0
	for i := range packets {
		if a != 0 && b != 0 {
			break
		}
		if packets[i] == "[[2]]" {
			a = i + 1
		} else if packets[i] == "[[6]]" {
			b = i + 1
		}
	}
	fmt.Println(a * b)
}

func Main(path string) {
	s := tools.Readfile(path)

	partOne(s)
	partTwo(s)
}
