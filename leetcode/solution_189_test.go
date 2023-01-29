package leetcode

import (
	"testing"
)

func TestRotate(t *testing.T) {
	input, k, output := []int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}
	Rotate(input, k)
	if !Equal(input, output) {
		t.Logf("For expected %v, but got %v", output, input)
		t.Fail()
	}

	input, k, output = []int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}
	Rotate(input, k)
	if !Equal(input, output) {
		t.Logf("For expected %v, but got %v", output, input)
		t.Fail()
	}

	input, k, output = []int{1, 2}, 2, []int{1, 2}
	Rotate(input, k)
	if !Equal(input, output) {
		t.Logf("For expected %v, but got %v", output, input)
		t.Fail()
	}

	input, k, output = []int{1, 2}, 3, []int{2, 1}
	Rotate(input, k)
	if !Equal(input, output) {
		t.Logf("For expected %v, but got %v", output, input)
		t.Fail()
	}

}
