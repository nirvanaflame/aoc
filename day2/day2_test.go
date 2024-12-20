package day2

import "testing"

type test struct {
	input    []int
	expected bool
}

var tests = []test{
	//{[]int{7, 6, 4, 2, 1}, true},
	//{[]int{1, 2, 7, 8, 9}, false},
	//{[]int{9, 7, 6, 2, 1}, false},
	//{[]int{1, 3, 2, 4, 5}, true},
	//{[]int{8, 6, 4, 4, 1}, true},
	//{[]int{1, 3, 6, 7, 9}, true},
	//{[]int{5, 5, 5, 5, 5}, false},                // All elements are the same
	//{[]int{1, 2, 3, 4, 5}, true},                 // Strictly increasing
	//{[]int{5, 4, 3, 2, 1}, true},                 // Strictly decreasing
	//{[]int{1, 3, 5, 7, 9}, true},                 // Increasing with gaps
	//{[]int{9, 7, 5, 3, 1}, true},                 // Decreasing with gaps
	//{[]int{1, 2, 3, 2, 1}, false},                // Increasing then decreasing
	//{[]int{5, 4, 3, 4, 5}, false},                // Decreasing then increasing
	//{[]int{1, 1, 2, 2, 3, 3}, false},             // Repeated pairs, increasing
	//{[]int{3, 3, 2, 2, 1, 1}, false},             // Repeated pairs, decreasing
	//{[]int{1, 2, 1, 2, 1}, false},                // Alternating increase and decrease
	//{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, true}, // Long strictly decreasing
	//{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, true}, // Long strictly increasing
	//{[]int{1, 3, 2, 4, 3, 5}, false},             // Mixed increasing and decreasing
	//{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, false}, // Long sequence of same elements
	//{[]int{7, 4, 2, 1, 1}, true},
	{[]int{3, 4, 2, 1, 0}, true},
}

func TestIsSafeCombination(t *testing.T) {
	for _, test := range tests {
		if output := validTolerateRow(test.input); output != test.expected {
			t.Errorf("Input %v give Output %v not equal to expected %v", test.input, output, test.expected)
		}
	}
}
