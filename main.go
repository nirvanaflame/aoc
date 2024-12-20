package main

import (
	//"aoc/day1"
	//"aoc/day2"
	"aoc/day3"
	"aoc/util"
	"fmt"
)

func main() {
	lines := util.ReadFile("day3/input.txt")
	answer := day3.Solve(lines)
	//answer := day2.Solve(lines)
	fmt.Println(answer)
}
