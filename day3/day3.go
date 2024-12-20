package day3

import (
	"fmt"
	"strconv"
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

	total := 0

	for strings.Contains(line, prefix) {
		startAt := strings.Index(line, prefix) + len(prefix)
		endAt := endAt(line, startAt)
		if endAt == -1 {
			fmt.Println("endAt == -1")
			break
		}
		nums := strings.Split(line[startAt:endAt], ",")

		if len(nums) < 2 {
			fmt.Println("len(nums) < 2")
			break
		}

		x, err := strconv.Atoi(nums[0])
		if err != nil {
			line = line[startAt:]
			fmt.Println("x produces:", err)
			continue
		}

		y, err := strconv.Atoi(nums[1])
		if err != nil {
			line = line[startAt:]
			fmt.Println("y produces:", err)
			continue
		}

		total += x * y
		line = line[endAt+1:]
	}

	return total
}

func endAt(line string, startAt int) int {
	for i := startAt; i < len(line); i++ {
		if line[i] == ')' {
			return i
		}
	}
	return -1
}
