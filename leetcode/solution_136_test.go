package leetcode

import "testing"

func TestSingleNumber(t *testing.T) {
	input, expected := []int{2, 2, 1}, 1
	result := singleNumber(input)
	if result != expected {
		t.Logf("For %v, expected %v, but got %v", input, expected, result)
		t.Fail()
	}
}
