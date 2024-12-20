package day1

import (
	"aoc/util"
	"math"
	"slices"
	"strings"
)

type entry struct {
	count      int
	multiplier int
}

func Solve(input []string) int {
	left := make([]int, len(input))
	right := make([]int, len(input))
	for i, row := range input {
		temp := strings.Split(row, "   ")
		left[i] = util.ParseInt(temp[0])
		right[i] = util.ParseInt(temp[1])
	}

	slices.Sort(left)
	slices.Sort(right)

	//total := findDifference(left, right)
	total := findSimilarity(left, right)

	return total
}

func findDifference(left []int, right []int) int {
	var total int
	for i := 0; i < len(left); i++ {
		total += int(math.Abs(float64(left[i] - right[i])))
	}
	return total
}

func findSimilarity(left []int, right []int) int {
	mem := map[int]entry{}
	for _, key := range left {
		v, ok := mem[key]
		if !ok {
			mem[key] = entry{0, 1}
		} else {
			mem[key] = entry{0, v.multiplier + 1}
		}
	}

	for _, key := range right {
		v, ok := mem[key]
		if ok {
			mem[key] = entry{v.count + 1, v.multiplier}
		}
	}

	var total int
	for k, entry := range mem {
		if entry.count > 0 {
			total += k * entry.count * entry.multiplier
		}
	}

	return total
}
