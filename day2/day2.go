package day2

import (
	"aoc/util"
	"strings"
)

func Solve(input []string) int {
	count := 0
	for _, l := range input {
		split := strings.Split(l, " ")
		nums := util.AsIntArray(split)
		if validTolerateRow(nums) {
			count++
		}
	}
	return count
}

func validTolerateRow(nums []int) bool {
	inc := isIncrementOrder(nums)

	count := inOrder(nums, inc)
	count2 := inOrder(nums, !inc)

	return count < 2 || count2 < 2
}

func inOrder(nums []int, inc bool) int {
	count := 0
	for i, safeIndex := 1, 0; i < len(nums); i++ {
		if isSafe(nums[safeIndex], nums[i], inc) {
			safeIndex = i
		} else {
			count++
		}
	}
	return count
}

func isSafe(current, next int, inc bool) bool {
	if inc && current < next && next-current < 4 {
		return true
	}
	if !inc && current > next && current-next < 4 {
		return true
	}
	return false
}

func isIncrementOrder(nums []int) bool {
	inc := false
	if nums[0] == nums[1] {
		inc = nums[0] < nums[2]
	} else {
		inc = nums[0] < nums[1]
	}
	return inc
}

func validRow(nums []int) int {
	if len(nums) < 2 || nums[0] == nums[1] {
		return 0
	}
	inc := false
	if nums[0] < nums[1] {
		inc = true
	}

	for i := 0; i < len(nums)-1; i++ {
		if !isSafe(nums[i], nums[i+1], inc) {
			return 0
		}
	}

	return 1
}
