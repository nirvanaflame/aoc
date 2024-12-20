package util

import (
	"log"
	"strconv"
)

func AsIntArray(strs []string) []int {
	result := make([]int, len(strs))
	for i, l := range strs {
		n, err := strconv.Atoi(l)
		if err != nil {
			log.Fatal(err)
		}
		result[i] = n
	}
	return result
}

func ParseInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
