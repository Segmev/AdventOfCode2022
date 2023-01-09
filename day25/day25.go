package day25

import (
	"fmt"
	"strings"

	"github.com/Segmev/AdventOfCode2022/tools"
)

var (
	SNAFU_KEYS = []byte{'=', '-', '0', '1', '2'}
	SNAFU      = map[byte]int{
		SNAFU_KEYS[0]: -2,
		SNAFU_KEYS[1]: -1,
		SNAFU_KEYS[2]: 0,
		SNAFU_KEYS[3]: 1,
		SNAFU_KEYS[4]: 2,
	}
)

func snafuToDec(line string) int {
	nb := 0
	for i := 0; i < len(line); i++ {
		nb = len(SNAFU_KEYS)*nb + SNAFU[line[i]]
	}
	return nb
}

func decToSnafu(nb int) string {
	if nb == 0 {
		return ""
	}
	q, r := (nb+2)/5, (nb+2)%5
	return decToSnafu(q) + string(SNAFU_KEYS[r])
}

func partOne(lines []string) {
	total := 0
	for _, line := range lines {
		total += snafuToDec(line)
	}
	fmt.Println(decToSnafu(total), "(for", total, ")")
}

func Main(path string) {
	s := tools.Readfile(path)

	lines := strings.Split(s, "\n")
	lines = lines[:len(lines)-1]
	partOne(lines)
}
