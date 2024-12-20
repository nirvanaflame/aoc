package day3

import (
	"aoc/util"
	"strings"
)

func Solve(lines []string) int {
	total := 0
	for _, line := range lines {
		total += solveLine(line)
	}

	return total
}

func solveLine(line string) int {
	prefix := "mul("
	mem := []rune{}

	total := 0
	for i := 0; i < len(line); i++ {
		ch := rune(line[i])

		mem = append(mem, ch)
		if prefix == string(mem) {
			i, first := getNumber(line, i+1, ',')
			i, second := getNumber(line, i+1, ')')

			total += first * second
			mem = []rune{}
		} else if !strings.HasPrefix(prefix, string(mem)) {
			mem = []rune{}
		}

	}

	return total
}

func getNumber(line string, start int, end rune) (index, number int) {
	var str []rune
	for ch := rune(line[start]); ch != end; ch = rune(line[start]) {
		if start == len(line)-1 {
			break
		}
		str = append(str, ch)
		start++
	}
	return start, util.ParseInt(string(str))
}
