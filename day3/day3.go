package day3

import "fmt"

func Solve(lines []string) int {
	total := 0
	for _, line := range lines {
		solveLine(line)
	}

	return total
}

func solveLine(line string) int {
	prefix := "mul("
	mem := []rune{}
	for _, rune := range line {
		mem = append(mem, rune)
		fmt.Println(mem)
		if prefix == string(mem) {

		}
	}

	return 0
}
